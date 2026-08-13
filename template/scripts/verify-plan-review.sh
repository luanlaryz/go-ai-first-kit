#!/usr/bin/env bash
# Plan review gate: a plan is only executable when it carries a valid, verifiable
# review artifact. Validates ARTIFACTS (frontmatter, hash, TTL, verdict file) and
# never calls a model, so it runs in CI with no API key.
set -euo pipefail

repo_root="$(git rev-parse --show-toplevel 2>/dev/null || pwd)"
cd "$repo_root"

# --changed: validate ONLY plans changed vs the base ref (diff-scoped), mirroring
# scripts/check-sdd-compliance.sh. Per-plan checks (status/path/hash/TTL/artifact/
# reviewer) stay STRICT for each changed plan; this only narrows the SCOPE so a PR
# is not blocked by unrelated plans whose review TTL lapsed on the base branch.
# Base ref: $BASE_REF, else origin/main|origin/master|main|master. In CI on a
# pull_request a missing base FAILS; locally without a base it is an explicit skip.
scope_changed=0
for arg in "$@"; do
  if [ "$arg" = "--changed" ]; then scope_changed=1; fi
done

changed_plans=""
if [ "$scope_changed" = "1" ]; then
  base="${BASE_REF:-}"
  if [ -z "$base" ]; then
    for cand in origin/main origin/master main master; do
      if git rev-parse --verify --quiet "$cand" >/dev/null 2>&1; then base="$cand"; break; fi
    done
  fi
  if [ -z "$base" ]; then
    if [ -n "${CI:-}" ] && [ "${GITHUB_EVENT_NAME:-}" = "pull_request" ]; then
      echo "verify-plan-review: FAIL: no base ref resolvable in CI; set BASE_REF (e.g. origin/main)" >&2
      exit 1
    fi
    echo "verify-plan-review: skipped (--changed: no base ref; set BASE_REF to enable locally)"
    exit 0
  fi
  # Added/modified plans only (exclude deletions: a removed plan needs no review).
  diff_plans="$(git diff --name-only --diff-filter=d "${base}...HEAD" -- '.cursor/plans/*.plan.md' 2>/dev/null || true)"
  # The diff is against HEAD, so a plan renamed in the working tree still shows up
  # under its old path while the file is already gone from disk. Drop paths that
  # no longer exist: the same reason deletions are excluded above. The new path is
  # still validated, either from this diff once committed or via --plan.
  changed_plans=""
  while IFS= read -r plan_path; do
    [ -n "$plan_path" ] || continue
    [ -f "$plan_path" ] || continue
    changed_plans="${changed_plans:+${changed_plans}
}${plan_path}"
  done <<EOF
${diff_plans}
EOF
fi

# PYTHONDONTWRITEBYTECODE: importing the shared risk module must not litter the
# working tree with __pycache__, which the packaging gate would flag.
VPR_SCOPE_CHANGED="$scope_changed" VPR_CHANGED_PLANS="$changed_plans" \
  PYTHONDONTWRITEBYTECODE=1 python3 - "$@" <<'PY'
from __future__ import annotations

import argparse
import hashlib
import os
import sys
from datetime import datetime, timedelta, timezone
from pathlib import Path

# The shell wrapper cd'd to the repo root. Risk classification is shared with
# scripts/review-plan.py: the verifier must not accept an OPERATOR_FALLBACK that
# the writer would have refused. A missing module is a hard failure, never a
# skipped check.
sys.path.insert(0, "scripts/lib")
from plan_risk import detect_high_risk  # noqa: E402

VALID_STATUSES = {"APPROVED", "APPROVED_WITH_CHANGES"}
VALID_PATHS = {"DUAL_AUTOMATED", "DUAL_TURN_BASED", "OPERATOR_FALLBACK"}


class PlanError(Exception):
    pass


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(description="Verify plan review artifacts")
    parser.add_argument("--plan", help="Verify one plan strictly")
    parser.add_argument("--all-strict", action="store_true", help="Require every plan to have an executable review")
    parser.add_argument(
        "--changed",
        action="store_true",
        help="Verify only plans changed vs the base ref (diff-scoped; strict per changed plan)",
    )
    return parser.parse_args()


def split_frontmatter(raw: str) -> tuple[dict[str, str], str, bool]:
    lines = raw.splitlines(keepends=True)
    if not lines or lines[0].strip() != "---":
        return {}, raw, False

    end_index = None
    for index in range(1, len(lines)):
        if lines[index].strip() == "---":
            end_index = index
            break

    if end_index is None:
        raise PlanError("frontmatter start found without closing ---")

    frontmatter_lines = lines[1:end_index]
    body = "".join(lines[end_index + 1 :])
    review = parse_review_block(frontmatter_lines)
    return review, body, True


def parse_review_block(frontmatter_lines: list[str]) -> dict[str, str]:
    review: dict[str, str] = {}
    in_review = False
    for line in frontmatter_lines:
        stripped = line.strip()
        if stripped == "review:":
            in_review = True
            continue
        if in_review:
            if not line.startswith("  "):
                break
            item = stripped
            if not item or ":" not in item:
                continue
            key, value = item.split(":", 1)
            review[key.strip()] = value.strip().strip('"').strip("'")
    return review


def compute_body_hash(body: str) -> str:
    return hashlib.sha256(body.encode("utf-8")).hexdigest()


def parse_reviewed_at(value: str) -> datetime:
    normalized = value.replace("Z", "+00:00")
    try:
        parsed = datetime.fromisoformat(normalized)
    except ValueError as exc:
        raise PlanError(f"invalid reviewed_at: {value}") from exc
    if parsed.tzinfo is None:
        raise PlanError("reviewed_at must include timezone")
    return parsed


def require(review: dict[str, str], key: str) -> str:
    value = review.get(key, "").strip()
    if not value:
        raise PlanError(f"missing review.{key}")
    return value


def verify_plan(path: Path, strict: bool) -> str:
    raw = path.read_text(encoding="utf-8")
    review, body, _ = split_frontmatter(raw)
    if not review:
        if strict:
            raise PlanError("missing review frontmatter block")
        return f"SKIP {path}: no review block; draft plans are not executable"

    status = require(review, "status")
    if status not in VALID_STATUSES:
        raise PlanError(f"review.status must be APPROVED or APPROVED_WITH_CHANGES, got {status}")

    review_path = require(review, "path")
    if review_path not in VALID_PATHS:
        raise PlanError(f"review.path must be one of {sorted(VALID_PATHS)}, got {review_path}")

    planner = require(review, "planner_model")
    reviewer = require(review, "reviewer_model")
    # Independence is the whole point of the review. Checking it only in the
    # writer would let a hand-edited frontmatter self-approve.
    if planner == reviewer:
        raise PlanError(
            f"planner_model and reviewer_model must differ, both are {planner!r}; "
            "self-approval is not review"
        )

    reviewed_at = parse_reviewed_at(require(review, "reviewed_at"))
    expected_hash = require(review, "plan_sha256")
    artifact = Path(require(review, "verdict_artifact"))
    ttl_days_raw = require(review, "ttl_days")

    try:
        ttl_days = int(ttl_days_raw)
    except ValueError as exc:
        raise PlanError(f"review.ttl_days must be integer, got {ttl_days_raw}") from exc

    if ttl_days <= 0 or ttl_days > 30:
        raise PlanError("review.ttl_days must be between 1 and 30")

    if review_path == "OPERATOR_FALLBACK":
        if not review.get("operator_fallback_reason", "").strip():
            raise PlanError("OPERATOR_FALLBACK requires operator_fallback_reason")
        # Same restriction the writer enforces: fallback is for low-risk plans.
        # Verifying it here closes the hand-written frontmatter path.
        risk = detect_high_risk(raw)
        if risk:
            raise PlanError(
                "OPERATOR_FALLBACK is not valid for a high-risk plan: "
                + ", ".join(risk)
                + "; review it through DUAL_TURN_BASED"
            )

    actual_hash = compute_body_hash(body)
    if actual_hash != expected_hash:
        raise PlanError(
            "O hash mudou; isso significa que o plano foi alterado depois do review. "
            f"expected {expected_hash}, got {actual_hash}"
        )

    if not artifact.exists():
        raise PlanError(f"verdict artifact not found: {artifact}")

    artifact_text = artifact.read_text(encoding="utf-8")
    if expected_hash not in artifact_text:
        raise PlanError(f"verdict artifact does not reference plan hash: {artifact}")

    now = datetime.now(timezone.utc)
    expires_at = reviewed_at.astimezone(timezone.utc) + timedelta(days=ttl_days)
    if now > expires_at:
        raise PlanError(f"review expired at {expires_at.isoformat()}")

    return f"OK {path}: {status} via {review_path}"


def main() -> int:
    args = parse_args()
    if args.plan:
        plans = [Path(args.plan)]
        strict = True
    elif os.environ.get("VPR_SCOPE_CHANGED") == "1":
        # Diff-scoped: only plans changed vs the base ref, resolved in the shell
        # wrapper. A plan touched by the diff MUST carry a valid review (strict).
        raw = os.environ.get("VPR_CHANGED_PLANS", "")
        plans = [Path(p) for p in raw.splitlines() if p.strip()]
        strict = True
        if not plans:
            print("verify-plan-review: no changed plans in diff; ok")
            return 0
    else:
        plans = sorted(Path(".cursor/plans").glob("*.plan.md"))
        strict = args.all_strict

    if not plans:
        print("verify-plan-review: no plan files found")
        return 0

    failed = False
    for plan in plans:
        try:
            print(verify_plan(plan, strict=strict))
        except PlanError as exc:
            failed = True
            print(f"FAIL {plan}: {exc}", file=sys.stderr)

    return 1 if failed else 0


if __name__ == "__main__":
    raise SystemExit(main())
PY
