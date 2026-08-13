#!/usr/bin/env bash
# SDD compliance gate. When a governed scope path (pkg/**, internal/**, api/**,
# migrations/**) changes in a diff, a corresponding spec / ADR / versioned change
# record must change too.
#
# Additionally, a newly ADDED construction spec (specs/NNN-*.md or
# specs/<sub>/NNN-*.md, non-diagnosis) requires its dual-spec diagnosis pair
# (NNN+1)-*-diagnosis.md in the SAME directory. Carve-out: specs/corrections/**
# and the meta-specs {000,001,010,020} have no mandatory pair.
#
# A plan in .cursor/plans/** does NOT count as a spec. PR-scoped: needs a base
# ref (BASE_REF, else origin/main...). In CI a missing base FAILS; locally
# without a base it is reported as skipped (never a silent gate pass).
#
# Declared coarseness: this gate only proves that SOME spec or ADR accompanies the
# diff, not that it is the RIGHT one. Correlating the change to the specific
# governing pair is done by scripts/check-governed-change.sh, which validates the
# dual-spec pair cited in the PR body against the backlog item. Duplicating that
# correlation here would mean guessing intent from paths.
set -euo pipefail
cd "$(git rev-parse --show-toplevel 2>/dev/null || pwd)"

fail() { printf 'check-sdd-compliance: FAIL: %s\n' "$*" >&2; exit 1; }

base="${BASE_REF:-}"
if [ -z "$base" ]; then
  for cand in origin/main origin/master main master; do
    if git rev-parse --verify --quiet "$cand" >/dev/null 2>&1; then base="$cand"; break; fi
  done
fi
if [ -z "$base" ]; then
  if [ -n "${CI:-}" ]; then fail "no base ref resolvable in CI; set BASE_REF (e.g. origin/main)"; fi
  printf 'check-sdd-compliance: skipped (no base ref; set BASE_REF to enable locally)\n'
  exit 0
fi

changed="$(git diff --name-only "$base...HEAD" 2>/dev/null || true)"
if [ -z "$changed" ]; then
  printf 'check-sdd-compliance: ok (no diff vs %s)\n' "$base"
  exit 0
fi

# Dual-spec: a newly added construction spec (any specs/ subdir) must ship its
# diagnosis pair in the SAME directory. Diff-scoped and namespace-aware.
added="$(git diff --name-only --diff-filter=A "$base...HEAD" 2>/dev/null || true)"
while IFS= read -r f; do
  [ -z "$f" ] && continue
  case "$f" in
    specs/corrections/*) ;;
    specs/*-diagnosis.md | specs/*/*-diagnosis.md) ;;
    specs/[0-9][0-9][0-9]-*.md | specs/*/[0-9][0-9][0-9]-*.md)
      base_name="$(basename "$f")"
      n=$((10#${base_name%%-*}))
      case "$n" in 0|1|10|20) continue ;; esac
      dir="$(dirname "$f")"
      pair="$(printf '%03d' $((n + 1)))"
      if ! ls "${dir}/${pair}-"*-diagnosis.md >/dev/null 2>&1; then
        fail "new construction spec $f requires its dual-spec pair ${dir}/${pair}-*-diagnosis.md"
      fi
      ;;
  esac
done < <(printf '%s\n' "$added")

if printf '%s\n' "$changed" | grep -Eq '^(pkg/|internal/|api/|migrations/)'; then
  if printf '%s\n' "$changed" | grep -Eiq '^(specs/|docs/decisions/)'; then
    printf 'check-sdd-compliance: ok (governed-scope change accompanied by spec/ADR)\n'
    exit 0
  fi
  printf 'changed files (vs %s):\n%s\n' "$base" "$changed" >&2
  fail "governed scope (pkg|internal|api|migrations) changed without a matching specs/** or docs/decisions/** change (a .cursor/plans/** file does not substitute a spec)"
fi

printf 'check-sdd-compliance: ok (no governed-scope contract change)\n'
exit 0
