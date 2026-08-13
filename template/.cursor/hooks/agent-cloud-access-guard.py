#!/usr/bin/env python3
"""Agent cloud access guard - beforeShellExecution classifier for aws/kubectl.

Reads the hook event JSON on stdin and emits a permission decision. The
classification is by TARGET (account, profile, context, endpoint), never by verb
alone: a read command pointed at production can leak as much as a write, so a
verb denylist with a read allowlist would leave the main exfiltration path open.

Targets are resolved against the marker lists in `cloud-access-config.json`.
Empty lists are safe, not permissive: nothing resolves to dev, every target
becomes `unknown`, mutation is denied and reading requires approval.

Fail-closed by contract: every internal error path emits `ask`, never `allow`.
See `docs/runbooks/agent-cloud-access.md` for the decision matrix and limits.
"""

from __future__ import annotations

import json
import os
import re
import shlex
import sys
import time
from pathlib import Path

CONFIG_PATH = ".cursor/hooks/cloud-access-config.json"
AUDIT_PATH = ".cursor/hooks/.agent-cloud-access-audit.log"
MAX_AUDIT_BYTES = 512_000

DEFAULT_LOCAL_MARKERS = ("localhost", "127.0.0.1", "host.docker.internal", "localstack", ":4566")

# Reading a secret is the highest-value exfiltration path, so it is classified
# apart from ordinary reads even though it is technically read-only.
SECRET_READ = (
    ("secretsmanager", "get-secret-value"),
    ("secretsmanager", "batch-get-secret-value"),
    ("ssm", "get-parameter"),
    ("ssm", "get-parameters"),
    ("ssm", "get-parameters-by-path"),
)

AWS_MUTATING_VERB = re.compile(
    r"^(create|delete|put|update|modify|remove|set|attach|detach|add|associate|disassociate"
    r"|purge|terminate|stop|start|reboot|restore|import|upload|copy|move|sync|rm|tag|untag"
    r"|enable|disable|register|deregister|invoke|send|publish|reset|rotate|cancel|abort|apply)"
    r"(-|$)"
)
AWS_READ_VERB = re.compile(r"^(describe|get|list|search|scan|query|head|lookup|filter|batch-get|select)(-|$)")

KUBECTL_MUTATING = {
    "apply", "create", "patch", "delete", "replace", "edit", "exec", "scale", "annotate",
    "label", "cordon", "uncordon", "drain", "taint", "rollout", "set", "attach",
    "port-forward", "proxy", "cp", "run", "expose", "autoscale", "evict",
}
KUBECTL_READ = {"get", "describe", "logs", "top", "explain", "api-resources", "api-versions", "version", "auth"}
KUBECTL_SECRET_RESOURCE = re.compile(r"^secrets?(\.|/|$)|^secret/")


def project_root() -> Path:
    return Path(os.environ.get("CURSOR_PROJECT_ROOT") or os.getcwd())


def load_markers() -> tuple[tuple[str, ...], tuple[str, ...], tuple[str, ...]]:
    """Return (prod, dev, local) markers, lowercased.

    A missing or malformed config yields empty prod/dev lists, which makes every
    target `unknown` - the restrictive end of the matrix.
    """
    prod: list[str] = []
    dev: list[str] = []
    local: list[str] = list(DEFAULT_LOCAL_MARKERS)
    try:
        raw = (project_root() / CONFIG_PATH).read_text(encoding="utf-8")
        config = json.loads(raw)
        if isinstance(config, dict):
            prod = [str(item) for item in config.get("prod_markers") or [] if str(item).strip()]
            dev = [str(item) for item in config.get("dev_markers") or [] if str(item).strip()]
            configured_local = [str(item) for item in config.get("local_markers") or [] if str(item).strip()]
            if configured_local:
                local = configured_local
    except (OSError, json.JSONDecodeError, TypeError):
        pass
    lower = lambda items: tuple(item.lower() for item in items)  # noqa: E731
    return lower(prod), lower(dev), lower(local)


PROD_MARKERS, DEV_MARKERS, LOCAL_MARKERS = load_markers()


def emit(permission: str, *, user_message: str = "", agent_message: str = "") -> None:
    payload: dict[str, str] = {"permission": permission}
    if user_message:
        payload["user_message"] = user_message
    if agent_message:
        payload["agent_message"] = agent_message
    sys.stdout.write(json.dumps(payload))
    sys.stdout.flush()


def audit(decision: str, tool: str, verb: str, target: str) -> None:
    """Append a redacted decision record. Never records the raw command, which
    may embed a credential; only the classification triple."""
    try:
        path = project_root() / AUDIT_PATH
        if path.exists() and path.stat().st_size > MAX_AUDIT_BYTES:
            path.write_text("", encoding="utf-8")
        line = "{} decision={} tool={} verb={} target={}\n".format(
            time.strftime("%Y-%m-%dT%H:%M:%S%z"), decision, tool, verb or "-", target
        )
        with path.open("a", encoding="utf-8") as handle:
            handle.write(line)
    except OSError:
        pass


def split_command(command: str) -> list[str]:
    try:
        return shlex.split(command)
    except ValueError:
        return command.split()


def env_target_hint() -> str:
    """Resolve target from the environment visible to the hook process.

    Required because the hook input carries the command string but not the agent
    shell environment: a secret read with no explicit flag can still resolve to
    production purely via an exported AWS_PROFILE.
    """
    blob = " ".join(
        os.environ.get(name, "")
        for name in ("AWS_PROFILE", "AWS_DEFAULT_PROFILE", "AWS_ENDPOINT_URL", "AWS_REGION")
    ).lower()
    if any(marker in blob for marker in PROD_MARKERS):
        return "prod"
    if any(marker in blob for marker in LOCAL_MARKERS):
        return "local"
    if any(marker in blob for marker in DEV_MARKERS):
        return "dev"
    return "unknown"


def classify_target(command: str) -> str:
    """Return prod | dev | local | unknown, resolving string first, env second."""
    haystack = command.lower()
    if any(marker in haystack for marker in PROD_MARKERS):
        return "prod"
    if any(marker in haystack for marker in LOCAL_MARKERS):
        return "local"
    if any(marker in haystack for marker in DEV_MARKERS):
        return "dev"
    return env_target_hint()


def strip_env_prefix(tokens: list[str]) -> list[str]:
    """Drop `env` and leading VAR=value assignments so the real binary surfaces."""
    out = list(tokens)
    while out:
        head = out[0]
        if head == "env":
            out = out[1:]
            continue
        if re.match(r"^[A-Za-z_][A-Za-z0-9_]*=", head):
            out = out[1:]
            continue
        break
    return out


def find_invocation(tokens: list[str], binary: str) -> list[str] | None:
    """Return the argv slice starting at `binary`, or None when the token only
    appears as data (e.g. `for value in python3 kubectl git`)."""
    for index, token in enumerate(tokens):
        base = token.rsplit("/", 1)[-1]
        if base != binary:
            continue
        if index == 0:
            return tokens[index:]
        previous = tokens[index - 1]
        if previous in {"env", "sudo", "command", "exec", "xargs", "time", "&&", "||", "|", ";"}:
            return tokens[index:]
        if re.match(r"^[A-Za-z_][A-Za-z0-9_]*=", previous):
            return tokens[index:]
        return None
    return None


def aws_service_and_verb(argv: list[str]) -> tuple[str, str]:
    positional = [token for token in argv[1:] if not token.startswith("-")]
    service = positional[0] if positional else ""
    verb = positional[1] if len(positional) > 1 else ""
    return service, verb


def decide_aws(argv: list[str], command: str) -> tuple[str, str, str, str]:
    service, verb = aws_service_and_verb(argv)
    target = classify_target(command)
    label = f"aws {service} {verb}".strip()

    is_secret_read = any(service == svc and verb == vrb for svc, vrb in SECRET_READ)
    if service == "sts" and verb == "get-caller-identity":
        return "allow", label, target, "identidade nao expoe segredo"

    if AWS_MUTATING_VERB.match(verb) and not AWS_READ_VERB.match(verb):
        if target in {"dev", "local"}:
            return "ask", label, target, "mutacao em ambiente de desenvolvimento"
        # Unknown target must not be treated as dev: the default credential
        # chain may resolve to production.
        return "deny", label, target, "mutacao em producao ou alvo nao comprovado"

    if is_secret_read:
        if target in {"dev", "local"}:
            return "allow", label, target, "leitura de segredo em dev/local"
        return "ask", label, target, "leitura de segredo em producao ou alvo nao comprovado"

    if target in {"dev", "local"}:
        return "allow", label, target, "leitura em dev/local"
    if target == "prod":
        return "ask", label, target, "leitura em producao"
    return "ask", label, target, "alvo nao comprovado"


def kubectl_subcommand(argv: list[str]) -> tuple[str, list[str]]:
    positional: list[str] = []
    skip_next = False
    for token in argv[1:]:
        if skip_next:
            skip_next = False
            continue
        if token.startswith("--"):
            if "=" not in token:
                skip_next = True
            continue
        if token.startswith("-") and len(token) > 1:
            skip_next = True
            continue
        positional.append(token)
    return (positional[0] if positional else ""), positional[1:]


def has_explicit_context(argv: list[str]) -> bool:
    for token in argv:
        if token == "--context" or token.startswith("--context="):
            return True
    return False


def decide_kubectl(argv: list[str], command: str) -> tuple[str, str, str, str]:
    verb, rest = kubectl_subcommand(argv)
    target = classify_target(command)
    label = f"kubectl {verb}".strip()

    if verb in {"config", "version", "api-resources", "api-versions", "explain"}:
        return "allow", label, target, "metadados de cliente"

    if not has_explicit_context(argv):
        return "deny", label, target, "kubectl sem --context explicito"

    if verb in KUBECTL_MUTATING:
        if target in {"dev", "local"}:
            return "ask", label, target, "mutacao em cluster de desenvolvimento"
        return "deny", label, target, "mutacao em producao ou alvo nao comprovado"

    if verb in KUBECTL_READ and any(KUBECTL_SECRET_RESOURCE.match(item) for item in rest):
        if target in {"dev", "local"}:
            return "allow", label, target, "leitura de secret em dev"
        return "ask", label, target, "leitura de secret em producao ou alvo nao comprovado"

    if target in {"dev", "local"}:
        return "allow", label, target, "leitura em dev/local"
    return "ask", label, target, "leitura em producao ou alvo nao comprovado"


def main() -> int:
    raw = sys.stdin.read()
    try:
        event = json.loads(raw) if raw.strip() else {}
    except json.JSONDecodeError:
        emit("ask", user_message="Guard de credencial nao conseguiu ler o evento do hook.")
        return 0

    command = str(event.get("command") or "")
    if not command.strip():
        emit("allow")
        return 0

    tokens = strip_env_prefix(split_command(command))
    if not tokens:
        emit("allow")
        return 0

    aws_argv = find_invocation(tokens, "aws")
    kube_argv = find_invocation(tokens, "kubectl")

    if aws_argv is not None:
        decision, label, target, reason = decide_aws(aws_argv, command)
    elif kube_argv is not None:
        decision, label, target, reason = decide_kubectl(kube_argv, command)
    else:
        # Not a cloud CLI invocation: psql, make, test entrypoints, plain text.
        emit("allow")
        return 0

    audit(decision, label.split(" ")[0], label, target)

    if decision == "allow":
        emit("allow")
        return 0

    if decision == "deny" and reason == "kubectl sem --context explicito":
        emit(
            "deny",
            user_message=(
                "Comando kubectl sem --context explicito foi bloqueado: o current-context pode "
                "apontar para producao."
            ),
            agent_message=(
                "Bloqueado: kubectl sem --context explicito. Reexecute informando o cluster, por "
                "exemplo `kubectl --context <contexto-dev> ...`. Nao confie no current-context."
            ),
        )
        return 0

    if decision == "deny":
        emit(
            "deny",
            user_message=f"Bloqueado ({label}): {reason}. Alvo classificado: {target}.",
            agent_message=(
                f"Bloqueado pelo guard de credencial: {reason} (alvo={target}). Mutacao em producao "
                "exige execucao humana pelo runbook. Se o alvo for dev, torne-o explicito com "
                "--profile/--context e confirme que o marcador esta em "
                ".cursor/hooks/cloud-access-config.json."
            ),
        )
        return 0

    emit(
        "ask",
        user_message=f"Requer aprovacao ({label}): {reason}. Alvo classificado: {target}.",
        agent_message=(
            f"Aguardando aprovacao humana: {reason} (alvo={target}). Para leitura em dev, torne o "
            "alvo explicito com --profile/--context e mantenha o marcador em "
            ".cursor/hooks/cloud-access-config.json para evitar a aprovacao."
        ),
    )
    return 0


if __name__ == "__main__":
    try:
        raise SystemExit(main())
    except SystemExit:
        raise
    except BaseException:  # fail-closed: never allow on internal error
        emit("ask", user_message="Guard de credencial falhou internamente; decisao segura = ask.")
        raise SystemExit(0)
