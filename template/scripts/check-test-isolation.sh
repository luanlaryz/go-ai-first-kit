#!/usr/bin/env bash
# Conformance suite isolation.
#
# The conformance suite is the executable contract. A PR that changes the
# contract test together with the feature it validates can always be made green:
# weaken the test, ship the feature. So the two must not travel in the same
# commit or PR.
#
# This does NOT conflict with the rule that behavior changes ship with tests:
# unit tests next to the code, fixtures and the security baseline stay free. Only
# the conformance tree is scope-isolated.
#
# Configurable via TEST_ISOLATION_PROTECTED_PREFIX.
#
# The per-commit loop ignores merge commits: CI materializes refs/pull/N/merge,
# and diffing against the second parent of a stale tip falsely lists feature
# files already on the base. The range union (base...HEAD) remains the
# authoritative mixed-scope check.
set -euo pipefail
cd "$(git rev-parse --show-toplevel)"

PROTECTED_PREFIX="${TEST_ISOLATION_PROTECTED_PREFIX:-test/conformance/}"
FEATURE_PREFIXES=("pkg/" "internal/" "api/" "migrations/" "cmd/")

base="${BASE_REF:-}"
if [ -z "$base" ]; then
  for cand in origin/main origin/master main master; do
    if git rev-parse --verify --quiet "$cand" >/dev/null 2>&1; then base="$cand"; break; fi
  done
fi

if [ -z "$base" ]; then
  if [ -n "${CI:-}" ]; then
    echo "check-test-isolation: FAIL: no BASE_REF in CI" >&2
    exit 1
  fi
  echo "check-test-isolation: skipped (no base ref; set BASE_REF)"
  exit 0
fi

is_protected() {
  local f="$1"
  [[ "$f" == "$PROTECTED_PREFIX"* ]]
}

is_feature() {
  local f="$1"
  for p in "${FEATURE_PREFIXES[@]}"; do
    [[ "$f" == "$p"* ]] && return 0
  done
  return 1
}

check_file_set() {
  local label="$1"
  shift
  local files=("$@")
  local has_prot=false has_feat=false
  for f in "${files[@]}"; do
    [ -z "$f" ] && continue
    if is_protected "$f"; then has_prot=true; fi
    if is_feature "$f"; then has_feat=true; fi
  done
  if $has_prot && $has_feat; then
    echo "check-test-isolation: FAIL: $label mixes $PROTECTED_PREFIX with feature scope" >&2
    echo "check-test-isolation: split the conformance change into its own commit/PR" >&2
    return 1
  fi
  return 0
}

changed=$(git diff --name-only "$base...HEAD" 2>/dev/null || true)
if [ -z "${CI:-}" ]; then
  wt=$(git diff --name-only HEAD 2>/dev/null || true)
  unstaged=$(git diff --name-only 2>/dev/null || true)
  untracked=$(git ls-files --others --exclude-standard 2>/dev/null || true)
  changed=$(printf '%s\n%s\n%s\n%s' "$changed" "$wt" "$unstaged" "$untracked" | sed '/^$/d' | sort -u)
fi
if [ -z "$changed" ]; then
  echo "check-test-isolation: ok (no changes since $base)"
  exit 0
fi
changed_arr=()
while IFS= read -r line; do [ -n "$line" ] && changed_arr+=("$line"); done <<< "$changed"

if ! check_file_set "PR range" "${changed_arr[@]}"; then
  exit 1
fi

for c in $(git rev-list --no-merges "$base"..HEAD 2>/dev/null); do
  files=$(git diff-tree --no-commit-id --name-only -r "$c" 2>/dev/null || true)
  arr=()
  while IFS= read -r line; do [ -n "$line" ] && arr+=("$line"); done <<< "$files"
  if ! check_file_set "commit $c" "${arr[@]}"; then
    exit 1
  fi
done
all_count=$(git rev-list --count "$base"..HEAD 2>/dev/null || echo 0)
nm_count=$(git rev-list --count --no-merges "$base"..HEAD 2>/dev/null || echo 0)
if [ "$all_count" -gt "$nm_count" ] 2>/dev/null; then
  echo "check-test-isolation: note: ignored $((all_count - nm_count)) merge commit(s) in per-commit loop (range union still enforced)"
fi

echo "check-test-isolation: ok (no mixed conformance + feature changes since $base)"
exit 0
