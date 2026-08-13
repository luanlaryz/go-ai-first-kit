#!/usr/bin/env python3
"""Prepare plan review artifacts without bypassing the review gate.

This script never approves a plan on its own authority: it computes the plan body
hash, writes the verdict artifact and, only when an independent reviewer is
recorded, writes the `review:` frontmatter block that
`scripts/verify-plan-review.sh` accepts. The plan body is never modified.
"""

from __future__ import annotations

import argparse
import hashlib
import os
import subprocess
import sys
from datetime import datetime, timezone
from pathlib import Path

# Risk classification is shared with scripts/verify-plan-review.sh so the writer
# and the verifier cannot disagree about what counts as high risk.
# Bytecode is disabled so importing a repo-local module never litters the working
# tree with __pycache__, which the packaging gate would then flag.
sys.dont_write_bytecode = True
sys.path.insert(0, str(Path(__file__).resolve().parent / "lib"))
from plan_risk import detect_high_risk  # noqa: E402

VALID_MODES = {"DUAL_AUTOMATED", "DUAL_TURN_BASED", "OPERATOR_FALLBACK"}


class PlanReviewError(Exception):
    pass


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Prepare plan review artifacts without bypassing the dual-model review gate."
    )
    parser.add_argument("--plan", required=True, help="Path to .cursor/plans/*.plan.md")
    parser.add_argument(
        "--mode",
        choices=sorted(VALID_MODES),
        default=os.getenv("PLAN_REVIEW_MODE"),
        help="Review mode. Defaults to env PLAN_REVIEW_MODE or safe auto-selection.",
    )
    parser.add_argument(
        "--operator-reviewed",
        action="store_true",
        help=(
            "Confirm that an independent operator/model already reviewed the "
            "final plan body. Required before DUAL_TURN_BASED writes APPROVED."
        ),
    )
    parser.add_argument(
        "--planner-model",
        default=os.getenv("PLAN_REVIEW_PLANNER_MODEL", ""),
        help="Planner model identifier. Defaults to PLAN_REVIEW_PLANNER_MODEL.",
    )
    parser.add_argument(
        "--reviewer-model",
        default=os.getenv("PLAN_REVIEW_REVIEWER_MODEL", ""),
        help="Reviewer model identifier. Defaults to PLAN_REVIEW_REVIEWER_MODEL.",
    )
    parser.add_argument(
        "--provider",
        default=os.getenv("PLAN_REVIEW_PROVIDER", ""),
        help="Optional review provider label. Defaults to PLAN_REVIEW_PROVIDER.",
    )
    parser.add_argument(
        "--ttl-days",
        default=os.getenv("PLAN_REVIEW_TTL_DAYS", "7"),
        help="Review TTL in days, from 1 to 30. Defaults to 7.",
    )
    parser.add_argument(
        "--operator-fallback-reason",
        default="",
        help="Required when --mode OPERATOR_FALLBACK is used.",
    )
    return parser.parse_args()


def find_repo_root(start: Path) -> Path:
    try:
        completed = subprocess.run(
            ["git", "rev-parse", "--show-toplevel"],
            cwd=start,
            check=True,
            capture_output=True,
            text=True,
        )
        root = completed.stdout.strip()
        if root:
            return Path(root)
    except (subprocess.CalledProcessError, FileNotFoundError):
        pass

    current = start.resolve()
    for candidate in [current, *current.parents]:
        if (candidate / "AGENTS.md").exists() and (candidate / "Makefile").exists():
            return candidate
    raise PlanReviewError("repo root not found; run from inside the repository")


def resolve_plan_path(repo_root: Path, value: str) -> Path:
    path = Path(value)
    if not path.is_absolute():
        path = repo_root / path
    path = path.resolve()
    try:
        path.relative_to(repo_root.resolve())
    except ValueError as exc:
        raise PlanReviewError("plan must be inside the repository") from exc
    if not path.exists():
        raise PlanReviewError(f"plan file not found: {path}")
    if not path.is_file():
        raise PlanReviewError(f"plan path is not a file: {path}")
    return path


def split_frontmatter(raw: str) -> tuple[list[str], str, bool]:
    lines = raw.splitlines(keepends=True)
    if not lines or lines[0].strip() != "---":
        return [], raw, False

    end_index = None
    for index in range(1, len(lines)):
        if lines[index].strip() == "---":
            end_index = index
            break

    if end_index is None:
        raise PlanReviewError("frontmatter starts with --- but has no closing ---")

    return lines[1:end_index], "".join(lines[end_index + 1 :]), True


def compute_body_hash(body: str) -> str:
    return hashlib.sha256(body.encode("utf-8")).hexdigest()


def plan_id_for(path: Path) -> str:
    name = path.name
    suffix = ".plan.md"
    if name.endswith(suffix):
        return name[: -len(suffix)]
    return path.stem


def parse_ttl_days(value: str) -> int:
    try:
        ttl_days = int(value)
    except ValueError as exc:
        raise PlanReviewError(f"PLAN_REVIEW_TTL_DAYS/--ttl-days must be an integer, got {value}") from exc
    if ttl_days <= 0 or ttl_days > 30:
        raise PlanReviewError("ttl days must be between 1 and 30")
    return ttl_days


def choose_mode(args: argparse.Namespace) -> str:
    if args.mode:
        return args.mode

    api_key = os.getenv("PLAN_REVIEW_API_KEY", "")
    planner = args.planner_model.strip()
    reviewer = args.reviewer_model.strip()
    if api_key and planner and reviewer and planner != reviewer:
        return "DUAL_AUTOMATED"
    return "DUAL_TURN_BASED"


def quote_yaml(value: str) -> str:
    escaped = value.replace("\\", "\\\\").replace('"', '\\"')
    return f'"{escaped}"'


def remove_existing_review(frontmatter_lines: list[str]) -> list[str]:
    kept: list[str] = []
    skipping_review = False
    for line in frontmatter_lines:
        if not skipping_review and line.strip() == "review:":
            skipping_review = True
            continue
        if skipping_review:
            if line.startswith("  ") or not line.strip():
                continue
            skipping_review = False
        kept.append(line)
    return kept


def format_review_block(
    *,
    status: str,
    mode: str,
    planner_model: str,
    reviewer_model: str,
    reviewed_at: str,
    plan_sha256: str,
    artifact_path: str,
    ttl_days: int,
    operator_fallback_reason: str = "",
) -> list[str]:
    lines = [
        "review:\n",
        f"  status: {quote_yaml(status)}\n",
        f"  path: {quote_yaml(mode)}\n",
        f"  planner_model: {quote_yaml(planner_model)}\n",
        f"  reviewer_model: {quote_yaml(reviewer_model)}\n",
        f"  reviewed_at: {quote_yaml(reviewed_at)}\n",
        f"  plan_sha256: {quote_yaml(plan_sha256)}\n",
        f"  verdict_artifact: {quote_yaml(artifact_path)}\n",
        f"  ttl_days: {ttl_days}\n",
    ]
    if mode == "OPERATOR_FALLBACK":
        lines.append(f"  operator_fallback_reason: {quote_yaml(operator_fallback_reason)}\n")
    return lines


def replace_review_frontmatter(
    *,
    frontmatter_lines: list[str],
    body: str,
    review_lines: list[str],
) -> str:
    updated_frontmatter = remove_existing_review(frontmatter_lines)
    if updated_frontmatter and updated_frontmatter[-1].strip():
        updated_frontmatter.append("\n")
    updated_frontmatter.extend(review_lines)
    return "---\n" + "".join(updated_frontmatter) + "---\n" + body


def next_artifact_path(repo_root: Path, plan_id: str) -> Path:
    directory = repo_root / "automation" / "PLAN_REVIEWS"
    directory.mkdir(parents=True, exist_ok=True)
    first = directory / f"{plan_id}.md"
    if not first.exists():
        return first

    version = 2
    while True:
        candidate = directory / f"{plan_id}-v{version}.md"
        if not candidate.exists():
            return candidate
        version += 1


def relative_to_repo(repo_root: Path, path: Path) -> str:
    return path.relative_to(repo_root).as_posix()


def require_different_models(planner_model: str, reviewer_model: str) -> None:
    if not planner_model.strip():
        raise PlanReviewError(
            "O review nao foi criado porque falta planner model. "
            "Defina PLAN_REVIEW_PLANNER_MODEL ou passe --planner-model."
        )
    if not reviewer_model.strip():
        raise PlanReviewError(
            "O review nao foi criado porque falta reviewer model independente. "
            "Defina PLAN_REVIEW_REVIEWER_MODEL ou passe --reviewer-model com um modelo diferente."
        )
    if planner_model.strip() == reviewer_model.strip():
        raise PlanReviewError(
            "O review nao foi criado porque planner e reviewer model sao iguais. "
            "Autoaprovacao silenciosa nao e permitida."
        )


def artifact_text(
    *,
    plan_rel: str,
    plan_id: str,
    status: str,
    mode: str,
    planner_model: str,
    reviewer_model: str,
    reviewed_at: str,
    ttl_days: int,
    plan_sha256: str,
    risk_reasons: list[str],
    reviewer_verdict: str,
    limitations: list[str],
    commands: list[str],
    provider: str,
) -> str:
    risk = "HIGH" if risk_reasons else "LOW"
    risk_details = ", ".join(risk_reasons) if risk_reasons else "No high-risk marker detected."
    limitation_lines = "\n".join(f"- {item}" for item in limitations)
    command_lines = "\n".join(f"- `{item}`" for item in commands)
    checklist_status = "PASS" if status in {"APPROVED", "APPROVED_WITH_CHANGES"} else "PENDING"

    return f"""# Plan Review Verdict

## Identity

- Plan: {plan_rel}
- Plan ID: {plan_id}
- Status: {status}
- Review path: {mode}
- Planner model: {planner_model or "not-recorded"}
- Reviewer model: {reviewer_model or "not-recorded"}
- Review provider: {provider or "not-configured"}
- Reviewed at: {reviewed_at}
- TTL days: {ttl_days}
- Plan SHA-256: {plan_sha256}

## Risk Classification

- Level: {risk}
- Reasons: {risk_details}

## Checklist

- Plan body hash computed using `scripts/verify-plan-review.sh` contract: PASS
- Plan body left unchanged by review preparation: PASS
- High-risk fallback guard checked: PASS
- Independent reviewer evidence present: {checklist_status}
- Approved frontmatter may be written: {checklist_status}

## Reviewer Verdict

{reviewer_verdict}

## Limitations

{limitation_lines}

## Commands To Run Next

{command_lines}
"""


def write_artifact(path: Path, text: str) -> None:
    path.write_text(text, encoding="utf-8")


def next_turn_based_command(plan_path: str) -> str:
    return (
        "python3 scripts/review-plan.py "
        f"--plan {plan_path} "
        "--mode DUAL_TURN_BASED "
        "--operator-reviewed "
        "--planner-model '<planner-model>' "
        "--reviewer-model '<reviewer-model-diferente>'"
    )


def fail(message: str, next_command: str | None = None) -> int:
    print(f"review-plan: BLOCKED: {message}", file=sys.stderr)
    print("review-plan: Este plano ainda nao e executavel.", file=sys.stderr)
    if next_command:
        print(
            f"review-plan: Para seguir com seguranca, rode este comando: {next_command}",
            file=sys.stderr,
        )
    print(
        "review-plan: alternativa segura: prepare um artifact DUAL_TURN_BASED, "
        "peca revisao a outro modelo/operador e so entao registre --operator-reviewed.",
        file=sys.stderr,
    )
    return 2


def main() -> int:
    args = parse_args()
    try:
        repo_root = find_repo_root(Path.cwd())
        plan_path = resolve_plan_path(repo_root, args.plan)
        raw = plan_path.read_text(encoding="utf-8")
        frontmatter_lines, body, _ = split_frontmatter(raw)
        plan_sha256 = compute_body_hash(body)
        ttl_days = parse_ttl_days(args.ttl_days)
        mode = choose_mode(args)
        if mode not in VALID_MODES:
            raise PlanReviewError(f"invalid review mode: {mode}")

        plan_id = plan_id_for(plan_path)
        plan_rel = relative_to_repo(repo_root, plan_path)
        artifact_path = next_artifact_path(repo_root, plan_id)
        artifact_rel = relative_to_repo(repo_root, artifact_path)
        reviewed_at = datetime.now(timezone.utc).isoformat()
        risk_reasons = detect_high_risk(raw)
        planner_model = args.planner_model.strip()
        reviewer_model = args.reviewer_model.strip()
        provider = args.provider.strip()

        if mode == "DUAL_AUTOMATED":
            api_key = os.getenv("PLAN_REVIEW_API_KEY", "")
            if not api_key:
                return fail(
                    "O review nao foi criado porque falta PLAN_REVIEW_API_KEY para "
                    "DUAL_AUTOMATED. Sem API key, o script nao pode fingir aprovacao automatica.",
                    next_command=(
                        f"python3 scripts/review-plan.py --plan {plan_rel} --mode DUAL_TURN_BASED"
                    ),
                )
            require_different_models(planner_model, reviewer_model)
            return fail(
                "DUAL_AUTOMATED exige integracao real de provider, que este repositorio nao "
                "configura por padrao. Nenhuma aprovacao automatica foi registrada.",
                next_command=(
                    f"python3 scripts/review-plan.py --plan {plan_rel} --mode DUAL_TURN_BASED"
                ),
            )

        if mode == "OPERATOR_FALLBACK" and risk_reasons:
            return fail(
                "Nao use OPERATOR_FALLBACK neste caso porque o plano toca area de alto risco: "
                + ", ".join(risk_reasons),
                next_command=(
                    f"python3 scripts/review-plan.py --plan {plan_rel} --mode DUAL_TURN_BASED"
                ),
            )

        if mode == "DUAL_TURN_BASED" and not args.operator_reviewed:
            text = artifact_text(
                plan_rel=plan_rel,
                plan_id=plan_id,
                status="PENDING_REVIEW",
                mode=mode,
                planner_model=planner_model,
                reviewer_model=reviewer_model,
                reviewed_at=reviewed_at,
                ttl_days=ttl_days,
                plan_sha256=plan_sha256,
                risk_reasons=risk_reasons,
                reviewer_verdict=(
                    "Este plano ainda nao e executavel. Peca para outro modelo/operador "
                    "revisar o corpo final do plano, depois rode novamente com "
                    "--operator-reviewed, --planner-model e --reviewer-model."
                ),
                limitations=[
                    "Este artifact e instrucional e nao e um veredito executavel.",
                    "O frontmatter do plano nao foi alterado.",
                    "O plano continua bloqueado ate review.status ser APPROVED ou APPROVED_WITH_CHANGES.",
                ],
                commands=[
                    next_turn_based_command(plan_rel),
                    f"scripts/verify-plan-review.sh --plan {plan_rel}",
                    "make verify-plan-reviews",
                ],
                provider=provider,
            )
            write_artifact(artifact_path, text)
            print(f"review-plan: prepared pending artifact: {artifact_rel}")
            print("review-plan: Este plano ainda nao e executavel.")
            print("review-plan: nenhum review aprovado foi escrito no frontmatter.")
            print(
                "review-plan: Para seguir com seguranca, rode este comando apos a revisao independente: "
                + next_turn_based_command(plan_rel)
            )
            return 0

        if mode == "DUAL_TURN_BASED":
            require_different_models(planner_model, reviewer_model)
            status = "APPROVED"
            operator_fallback_reason = ""
            reviewer_verdict = (
                "APPROVED. Operator asserted that a different reviewer model reviewed "
                "the final plan body in a separate turn."
            )
            limitations = [
                "The script records the operator-supplied review evidence; it does not call a reviewer API.",
                "Any later plan body change invalidates this verdict by hash.",
            ]
        else:
            if not args.operator_reviewed:
                raise PlanReviewError(
                    "OPERATOR_FALLBACK exige --operator-reviewed e checklist humano explicito."
                )
            if not args.operator_fallback_reason.strip():
                raise PlanReviewError(
                    "OPERATOR_FALLBACK exige --operator-fallback-reason com uma justificativa clara."
                )
            if not planner_model:
                raise PlanReviewError(
                    "O review nao foi criado porque falta planner model. "
                    "Defina PLAN_REVIEW_PLANNER_MODEL ou passe --planner-model."
                )
            status = "APPROVED"
            reviewer_model = reviewer_model or "human-operator"
            operator_fallback_reason = args.operator_fallback_reason.strip()
            reviewer_verdict = (
                "APPROVED via OPERATOR_FALLBACK. Human operator completed the checklist "
                "and provided an explicit fallback reason."
            )
            limitations = [
                "OPERATOR_FALLBACK is valid only for low-risk plans.",
                "Any later plan body change invalidates this verdict by hash.",
            ]

        text = artifact_text(
            plan_rel=plan_rel,
            plan_id=plan_id,
            status=status,
            mode=mode,
            planner_model=planner_model,
            reviewer_model=reviewer_model,
            reviewed_at=reviewed_at,
            ttl_days=ttl_days,
            plan_sha256=plan_sha256,
            risk_reasons=risk_reasons,
            reviewer_verdict=reviewer_verdict,
            limitations=limitations,
            commands=[
                f"scripts/verify-plan-review.sh --plan {plan_rel}",
                "make verify-plan-reviews",
            ],
            provider=provider,
        )
        write_artifact(artifact_path, text)

        review_lines = format_review_block(
            status=status,
            mode=mode,
            planner_model=planner_model,
            reviewer_model=reviewer_model,
            reviewed_at=reviewed_at,
            plan_sha256=plan_sha256,
            artifact_path=artifact_rel,
            ttl_days=ttl_days,
            operator_fallback_reason=operator_fallback_reason,
        )
        updated = replace_review_frontmatter(
            frontmatter_lines=frontmatter_lines,
            body=body,
            review_lines=review_lines,
        )
        plan_path.write_text(updated, encoding="utf-8")

        print(f"review-plan: wrote approved artifact: {artifact_rel}")
        print(f"review-plan: updated only review frontmatter in: {plan_rel}")
        print("review-plan: Para seguir com seguranca, rode este comando: make verify-plan-reviews")
        return 0
    except PlanReviewError as exc:
        return fail(str(exc))


if __name__ == "__main__":
    raise SystemExit(main())
