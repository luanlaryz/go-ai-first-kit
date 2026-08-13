#!/usr/bin/env bash
# Governed change enforcement gate.
#
# When the diff touches governed scope (pkg/**, internal/**, api/**,
# migrations/**), the PR body must declare the quartet: a backlog item
# (BLG-NNNN with an implementable status), a dual-spec pair (construction NNN +
# diagnosis NNN+1), a plan with an approved review, and the regression coverage
# obligation. Cross-checked against docs/backlog/Backlog.md, specs/ and
# scripts/verify-plan-review.sh.
#
# Security: the PR body is resolved by precedence
#   $PR_BODY (file) > .pull_request.body from $GITHUB_EVENT_PATH (JSON)
#   > gh pr view of the branch's open PR (local only, fail-closed) > empty
# and is NEVER interpolated into a shell command. PR-scoped (BASE_REF). On push
# in CI (no PR context): explicit skip. Locally without a body: only skips with
# PROOF (via gh) that no PR is open; gh missing/erroring with governed scope
# touched = FAIL. On pull_request with no extractable body: FAIL.
set -euo pipefail
cd "$(git rev-parse --show-toplevel 2>/dev/null || pwd)"

base="${BASE_REF:-}"
if [ -z "$base" ]; then
  for cand in origin/main origin/master main master; do
    if git rev-parse --verify --quiet "$cand" >/dev/null 2>&1; then base="$cand"; break; fi
  done
fi
if [ -z "$base" ]; then
  if [ -n "${CI:-}" ] && [ "${GITHUB_EVENT_NAME:-}" = "pull_request" ]; then
    echo "check-governed-change: FAIL: no base ref resolvable in CI; set BASE_REF (e.g. origin/main)" >&2
    exit 1
  fi
  echo "check-governed-change: skipped (no base ref; set BASE_REF to enable locally)"
  exit 0
fi

changed="$(git diff --name-only "$base...HEAD" 2>/dev/null || true)"
explicit_body="${1:-${PR_BODY:-}}"
pr_body_file="$explicit_body"

gh_status=""
gh_tmp=""
cleanup_gh_tmp() { if [ -n "$gh_tmp" ]; then rm -f "$gh_tmp"; fi; }
trap cleanup_gh_tmp EXIT

# In pull_request CI, prefer the live body by PR number: it works with a detached
# HEAD and avoids a stale payload when the body is edited mid-run. Never
# overrides an explicit PR_BODY/$1.
if [ -z "$explicit_body" ] || [ ! -s "${explicit_body}" ]; then
  if [ "${GITHUB_EVENT_NAME:-}" = "pull_request" ]; then
    live_tmp="$(mktemp)"
    if ./scripts/resolve-pr-body-by-number.sh "$live_tmp"; then
      pr_body_file="$live_tmp"
      gh_tmp="$live_tmp"
      gh_status="resolved-by-number"
    else
      rm -f "$live_tmp"
    fi
  fi
fi

if { [ -z "$pr_body_file" ] || [ ! -s "$pr_body_file" ]; } \
  && { [ -z "${GITHUB_EVENT_PATH:-}" ] || [ ! -f "${GITHUB_EVENT_PATH:-}" ]; } \
  && [ -z "${GITHUB_EVENT_NAME:-}${CI:-}" ]; then
  gh_tmp="$(mktemp)"
  rc=0
  ./scripts/resolve-open-pr-body.sh "$gh_tmp" || rc=$?
  case "$rc" in
    0) pr_body_file="$gh_tmp"; gh_status="resolved" ;;
    2) gh_status="no-pr" ;;
    3) gh_status="gh-missing" ;;
    *) gh_status="gh-error" ;;
  esac
fi

CGC_CHANGED="$changed" CGC_PR_BODY_FILE="$pr_body_file" CGC_GH_STATUS="$gh_status" python3 - <<'PY'
from __future__ import annotations

import datetime
import fnmatch
import json
import os
import re
import subprocess
import sys
from pathlib import Path

GOVERNED = re.compile(r"^(pkg/|internal/|api/|migrations/)")
ALLOWLIST = Path(".guardrails/governed-change-exceptions.yaml")
BACKLOG = Path("docs/backlog/Backlog.md")
# Statuses under which implementation may legitimately be in flight.
VALID_STATUS = {"ready_for_implementation", "in_progress", "done"}
# Regression evidence must point at the project's versioned test tree.
TEST_ROOT = os.environ.get("GOVERNED_TEST_ROOT", "test/")


def out(msg: str) -> None:
    print(f"check-governed-change: {msg}")


def fail(msg: str) -> int:
    print(f"check-governed-change: FAIL: {msg}", file=sys.stderr)
    return 1


REQUIRED_ALLOWLIST_FIELDS = ("path", "rule", "justification", "owner", "expires_at", "ref")


def load_allowlist_globs() -> tuple[list[str], list[str], list[str]]:
    """Return (active_globs, expired_globs, schema_errors) from the YAML allowlist.

    Enforces the same field contract as .guardrails/README.md and the Go
    guardrails tool. An entry missing owner, justification or ref is rejected
    rather than silently honoured: an exception nobody owns is not an exception,
    it is a permanent hole.
    """
    if not ALLOWLIST.exists():
        return [], [], []

    active: list[str] = []
    expired: list[str] = []
    errors: list[str] = []
    today = datetime.date.today()
    entries: list[dict[str, str]] = []
    current: dict[str, str] | None = None

    for lineno, raw in enumerate(ALLOWLIST.read_text(encoding="utf-8").splitlines(), start=1):
        line = raw.strip()
        if not line or line.startswith("#") or line.startswith("exceptions:") or line.startswith("version:"):
            continue
        if line.startswith("- "):
            current = {"_line": str(lineno)}
            entries.append(current)
            line = line[2:].strip()
            if not line:
                continue
        if current is None:
            errors.append(f"{ALLOWLIST}:{lineno}: field outside of a list item")
            continue
        key, _, value = line.partition(":")
        if not _:
            errors.append(f"{ALLOWLIST}:{lineno}: malformed line")
            continue
        current[key.strip()] = value.strip().strip('"').strip("'")

    for entry in entries:
        lineno = entry.get("_line", "?")
        missing = [f for f in REQUIRED_ALLOWLIST_FIELDS if not entry.get(f)]
        if missing:
            errors.append(f"{ALLOWLIST}:{lineno}: entry missing required field(s): {', '.join(missing)}")
            continue
        path_val = entry["path"]
        try:
            exp_date = datetime.date.fromisoformat(entry["expires_at"])
        except ValueError:
            errors.append(f"{ALLOWLIST}:{lineno}: invalid expires_at {entry['expires_at']!r} (want YYYY-MM-DD)")
            expired.append(path_val)
            continue
        (active if exp_date >= today else expired).append(path_val)

    return active, expired, errors


def matches(path: str, pattern: str) -> bool:
    if not pattern:
        return False
    if path == pattern or path.startswith(pattern.rstrip("/") + "/"):
        return True
    return fnmatch.fnmatch(path, pattern) or fnmatch.fnmatch(path, pattern.rstrip("/") + "/**")


def resolve_body() -> tuple[str | None, str]:
    pr_body_file = os.environ.get("CGC_PR_BODY_FILE", "").strip()
    if pr_body_file and Path(pr_body_file).is_file():
        return Path(pr_body_file).read_text(encoding="utf-8"), "PR_BODY"
    event_path = os.environ.get("GITHUB_EVENT_PATH", "")
    if event_path and Path(event_path).is_file():
        try:
            data = json.loads(Path(event_path).read_text(encoding="utf-8"))
        except (ValueError, OSError):
            return None, "event-unreadable"
        pr = data.get("pull_request") or {}
        body = pr.get("body")
        if isinstance(body, str):
            return body, "GITHUB_EVENT_PATH"
    return None, "none"


def section(body: str, header: str) -> str:
    # Capture text from a `## header` line until the next `## ` heading.
    pattern = re.compile(rf"^##\s+{re.escape(header)}.*$", re.MULTILINE)
    m = pattern.search(body)
    if not m:
        return ""
    start = m.end()
    nxt = re.search(r"^##\s+", body[start:], re.MULTILINE)
    return body[start: start + nxt.start()] if nxt else body[start:]


def backlog_block(item: str) -> str | None:
    """Return the item block from the single-file backlog.

    The kit keeps one backlog file, so the item is delimited by its own `###`
    heading and the next `###`/`##` heading.
    """
    if not BACKLOG.is_file():
        return None
    text = BACKLOG.read_text(encoding="utf-8")
    m = re.search(rf"^###\s+{re.escape(item)}\b.*$", text, re.MULTILINE)
    if not m:
        return None
    start = m.end()
    nxt = re.search(r"^#{2,3}\s+", text[start:], re.MULTILINE)
    return text[start: start + nxt.start()] if nxt else text[start:]


def item_status(block: str) -> str | None:
    m = re.search(r"^-\s*\*\*Status\*\*:\s*(.+)$", block, re.MULTILINE)
    if not m:
        return None
    return m.group(1).strip().strip("`").strip()


def main() -> int:
    changed = [c for c in os.environ.get("CGC_CHANGED", "").splitlines() if c.strip()]
    active_globs, expired_globs, allowlist_errors = load_allowlist_globs()
    if allowlist_errors:
        for e in allowlist_errors:
            print(f"check-governed-change: - {e}", file=sys.stderr)
        return fail("governed-change allowlist does not satisfy the required schema")

    governed = [f for f in changed if GOVERNED.match(f)]
    # Expired exception used by a changed file is a hard failure.
    for f in governed:
        for g in expired_globs:
            if matches(f, g):
                return fail(f"allowlist entry for '{g}' is expired but matches changed file '{f}'")
    governed = [f for f in governed if not any(matches(f, g) for g in active_globs)]

    if not governed:
        out("ok (no governed-scope change)")
        return 0

    body, source = resolve_body()
    event_name = os.environ.get("GITHUB_EVENT_NAME", "")
    gh_status = os.environ.get("CGC_GH_STATUS", "")
    if body is None or not body.strip():
        if event_name == "pull_request":
            return fail("pull_request touched governed scope but PR body is empty/unreadable")
        if gh_status == "no-pr":
            out("skipped (governed scope changed but no open PR for this branch; proven via gh)")
            return 0
        if gh_status == "gh-missing":
            return fail(
                "governed scope changed and gh is not installed; "
                "run 'make pre-pr PR_BODY=<body.md>' or install gh"
            )
        if gh_status in ("gh-error", "resolved"):
            return fail(
                "governed scope changed but the branch's open PR body could not be resolved via gh; "
                "run 'make pre-pr PR_BODY=<body.md>'"
            )
        out("skipped (governed scope changed but no PR body; provide PR_BODY locally or open a PR)")
        return 0

    errors: list[str] = []

    # 1) Backlog: ## Backlog -> BLG-NNNN with an implementable status.
    # Scoped strictly to the section. The sections are mandatory in the PR
    # template, so falling back to the whole body would let an id cited under
    # '## Regressao' satisfy '## Backlog' and defeat the structure.
    backlog_sec = section(body, "Backlog")
    item_m = re.search(r"BLG-\d{4}", backlog_sec)
    item = item_m.group(0) if item_m else None
    block = backlog_block(item) if item else None
    if not item:
        errors.append("missing BLG-NNNN in '## Backlog' section of the PR body")
    elif block is None:
        errors.append(f"{item} not found as a '### {item}' entry in {BACKLOG}")
    else:
        status = item_status(block)
        if status not in VALID_STATUS:
            errors.append(f"{item} status must be in {sorted(VALID_STATUS)}, got '{status or 'none'}'")

    # 2) Specs: ## Specs lidas -> construction NNN + diagnosis NNN+1, both existing.
    specs_sec = section(body, "Specs lidas")
    spec_paths = re.findall(r"specs/\d{3}-[\w./-]*?\.md", specs_sec)
    diagnosis = [p for p in spec_paths if p.endswith("-diagnosis.md")]
    construction = [p for p in spec_paths if not p.endswith("-diagnosis.md")]
    pair_c = pair_d = ""
    if not construction or not diagnosis:
        errors.append(
            "'## Specs lidas' must cite a construction spec (specs/NNN-...md) and its diagnosis "
            "spec (specs/(NNN+1)-...-diagnosis.md)"
        )
    else:
        ok_pair = False
        for c in construction:
            cn = int(re.search(r"specs/(\d{3})", c).group(1))
            for d in diagnosis:
                dn = int(re.search(r"specs/(\d{3})", d).group(1))
                if dn == cn + 1 and Path(c).exists() and Path(d).exists():
                    ok_pair = True
                    pair_c, pair_d = c, d
                    break
            if ok_pair:
                break
        if not ok_pair:
            errors.append(
                "no valid dual-spec pair found (need specs/NNN + specs/(NNN+1)-...-diagnosis.md, "
                "both existing on disk)"
            )
        elif block is not None:
            # Cross-link: the backlog item should reference the same pair.
            c_num = pair_c.rsplit("/", 1)[-1].split("-")[0]
            d_num = pair_d.rsplit("/", 1)[-1].split("-")[0]
            if c_num not in block or d_num not in block:
                errors.append(
                    f"{item} should reference the dual-spec pair {pair_c} / {pair_d} "
                    "(Spec governante / Spec de diagnostico)"
                )

    # 3) Autopilot: ## Autopilot -> plan with an approved review.
    autopilot_sec = section(body, "Autopilot")
    plan_m = re.search(r"\.cursor/plans/[\w./-]+\.plan\.md", autopilot_sec)
    if not plan_m:
        errors.append("'## Autopilot' must cite the plan path (.cursor/plans/<id>.plan.md)")
    else:
        plan_path = plan_m.group(0)
        if not Path(plan_path).exists():
            errors.append(f"plan not found on disk: {plan_path}")
        else:
            res = subprocess.run(
                ["scripts/verify-plan-review.sh", "--plan", plan_path],
                capture_output=True, text=True,
            )
            if res.returncode != 0:
                detail = (res.stderr or res.stdout).strip().splitlines()
                tail = detail[-1] if detail else "review verification failed"
                errors.append(f"plan {plan_path} has no valid approved review: {tail}")

    # 4) Regressao: ## Regressao -> backlog item + target scenario in the test tree.
    regression_sec = section(body, "Regressao") or section(body, "Regressão")
    if not regression_sec.strip():
        errors.append(
            f"'## Regressao' must declare BLG-NNNN and the target scenario under {TEST_ROOT}"
        )
    else:
        reg_item_m = re.search(r"BLG-\d{4}", regression_sec)
        if not reg_item_m:
            errors.append("'## Regressao' must cite BLG-NNNN for regression coverage")
        elif item and reg_item_m.group(0) != item:
            errors.append("'## Regressao' BLG must match '## Backlog' BLG")
        # Require a concrete package::scenario reference, not just the substring
        # "test/": a bare mention proves nothing about what will be covered.
        scenario = re.compile(rf"{re.escape(TEST_ROOT)}[\w./-]+::[\w./-]+")
        if not scenario.search(regression_sec):
            errors.append(f"'## Regressao' must reference {TEST_ROOT}<package>::<scenario>")

    if errors:
        for e in errors:
            print(f"check-governed-change: - {e}", file=sys.stderr)
        return fail(
            "governed-scope change is missing the required quartet "
            "(backlog + dual-spec + reviewed plan + regressao)"
        )

    out(f"ok (governed change linked to {item} + dual-spec + reviewed plan + regressao; body via {source})")
    return 0


sys.exit(main())
PY
