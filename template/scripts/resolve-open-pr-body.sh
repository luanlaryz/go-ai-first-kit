#!/usr/bin/env bash
# Resolve, via gh, the body of the OPEN PR for the current branch (local use).
# Usage: resolve-open-pr-body.sh <out-file>
# Exit codes:
#   0 = body written to <out-file> (single open PR found)
#   2 = provably NO open PR for the branch (the only legitimate skip)
#   3 = gh unavailable (not installed)
#   4 = could not prove (unauthenticated/network/detached HEAD/ambiguous PR)
set -euo pipefail

out="${1:?usage: resolve-open-pr-body.sh <out-file>}"

command -v gh >/dev/null 2>&1 || exit 3

branch="$(git rev-parse --abbrev-ref HEAD 2>/dev/null || true)"
if [ -z "$branch" ] || [ "$branch" = "HEAD" ]; then
  exit 4
fi

if ! numbers="$(gh pr list --head "$branch" --state open --json number --jq '.[].number' 2>/dev/null)"; then
  exit 4
fi

count="$(printf '%s\n' "$numbers" | sed '/^$/d' | wc -l | tr -d ' ')"
if [ "$count" = "0" ]; then
  exit 2
fi
if [ "$count" != "1" ]; then
  exit 4
fi

number="$(printf '%s\n' "$numbers" | sed '/^$/d' | head -n1)"
gh pr view "$number" --json body --jq '.body // ""' > "$out" 2>/dev/null || exit 4
exit 0
