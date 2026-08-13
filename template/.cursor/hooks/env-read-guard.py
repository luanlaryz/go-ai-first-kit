#!/usr/bin/env python3
"""Env read guard - preToolUse classifier for Read.

Requires human approval before the agent reads a `.env` file. Implemented on
`preToolUse` (matcher `Read`) rather than `beforeReadFile` because `preToolUse`
documents `permission` support, while `beforeReadFile` matches by tool type and
its permission support is not part of the documented output contract.

Path filtering lives here, in the script, not in the matcher.

This guard does not affect test suites or dev servers: those load `.env` inside
their own processes (source/dotenv), which never goes through the agent's Read
tool. Use `make env-keys` to list which keys exist without exposing values.
"""

from __future__ import annotations

import json
import re
import sys

ENV_FILE = re.compile(r"(^|/)\.env(\.[A-Za-z0-9_.-]+)?$")
ALLOWED_SUFFIX = re.compile(r"\.env\.example$|\.env\.sample$|\.env\.template$")


def emit(permission: str, *, user_message: str = "", agent_message: str = "") -> None:
    payload: dict[str, str] = {"permission": permission}
    if user_message:
        payload["user_message"] = user_message
    if agent_message:
        payload["agent_message"] = agent_message
    sys.stdout.write(json.dumps(payload))
    sys.stdout.flush()


def collect_paths(node: object, found: list[str]) -> None:
    """Walk the event payload collecting values of any *path-ish* key, so the
    guard survives small differences in the tool input schema."""
    if isinstance(node, dict):
        for key, value in node.items():
            if isinstance(value, str) and (
                key.endswith("path") or key.endswith("Path") or key in {"file", "filename", "target"}
            ):
                found.append(value)
            else:
                collect_paths(value, found)
    elif isinstance(node, list):
        for item in node:
            collect_paths(item, found)


def main() -> int:
    raw = sys.stdin.read()
    try:
        event = json.loads(raw) if raw.strip() else {}
    except json.JSONDecodeError:
        emit("ask", user_message="Guard de .env nao conseguiu ler o evento; decisao segura = ask.")
        return 0

    paths: list[str] = []
    collect_paths(event, paths)

    for path in paths:
        if ALLOWED_SUFFIX.search(path):
            continue
        if ENV_FILE.search(path):
            emit(
                "ask",
                user_message=(
                    f"O agent quer ler {path}, que pode conter credenciais. Aprove somente se for "
                    "necessario."
                ),
                agent_message=(
                    "Leitura de .env requer aprovacao humana. Para descobrir quais variaveis existem "
                    "sem expor valores, use `make env-keys`."
                ),
            )
            return 0

    emit("allow")
    return 0


if __name__ == "__main__":
    try:
        raise SystemExit(main())
    except SystemExit:
        raise
    except BaseException:  # fail-closed
        emit("ask", user_message="Guard de .env falhou internamente; decisao segura = ask.")
        raise SystemExit(0)
