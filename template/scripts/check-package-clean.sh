#!/usr/bin/env bash
# Packaging cleanliness gate. Forbidden artifacts must not be git-tracked nor
# present in a delivery package dir ($PACKAGE_DIR / dist / release). Gitignored
# workspace build outputs are NOT scanned - the target is what ships and what is
# committed, not what a developer has locally.
# Allowlist: .guardrails/package-exceptions.yaml (validated path prefixes).
set -euo pipefail
cd "$(git rev-parse --show-toplevel 2>/dev/null || pwd)"

fail() { printf 'check-package-clean: FAIL: %s\n' "$*" >&2; exit 1; }

# Validated allowed path prefixes; the tool also enforces allowlist expiry.
allowed="$(go run ./tools/guardrails allowlist-paths .guardrails/package-exceptions.yaml)"

is_allowed() {
  local f="$1" pref
  while IFS= read -r pref; do
    [ -n "$pref" ] || continue
    [[ "$f" == "$pref"* ]] && return 0
  done <<< "$allowed"
  return 1
}

violations=()

# 1. git-tracked artifacts that should never be committed.
while IFS= read -r f; do
  [ -n "$f" ] || continue
  is_allowed "$f" && continue
  case "$f" in
    */node_modules/*|node_modules/*) violations+=("tracked node_modules: $f") ;;
    */coverage/*|coverage/*) violations+=("tracked coverage: $f") ;;
    *.coverage.out|coverage.out) violations+=("tracked coverage profile: $f") ;;
    *.test) violations+=("tracked compiled test binary: $f") ;;
    bin/*|*/bin/*) violations+=("tracked build output: $f") ;;
  esac
done < <(git ls-files)

# 2. delivery package staging dirs.
for d in "${PACKAGE_DIR:-}" dist release; do
  [ -n "$d" ] && [ -d "$d" ] || continue
  while IFS= read -r entry; do
    [ -n "$entry" ] && violations+=("package dir $d contains forbidden artifact: $entry")
  done < <(find "$d" \( -name .git -o -name node_modules -o -name coverage -o -name '*.test' -o -name 'coverage.out' \) -print 2>/dev/null)
done

if [ ${#violations[@]} -gt 0 ]; then
  printf '%s\n' "${violations[@]}" >&2
  fail "${#violations[@]} forbidden packaging artifact(s) found"
fi

printf 'check-package-clean: ok\n'
