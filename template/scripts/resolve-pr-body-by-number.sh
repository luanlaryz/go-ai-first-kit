#!/usr/bin/env bash
# Resolve the live body of a PR by number from $GITHUB_EVENT_PATH (CI).
# Usage: resolve-pr-body-by-number.sh <out-file>
# Exit codes:
#   0 = body written to <out-file>
#   2 = event path missing / no pull_request.number
#   3 = no transport available (neither token+curl nor gh)
#   4 = failed to read the body (auth/network/PR not found)
# Does not use the local branch, so it works with a detached HEAD on
# refs/pull/N/merge. Preference: GitHub API via curl+token (Actions) > gh CLI.
set -euo pipefail

out="${1:?usage: resolve-pr-body-by-number.sh <out-file>}"

if [ -z "${GITHUB_EVENT_PATH:-}" ] || [ ! -f "${GITHUB_EVENT_PATH}" ]; then
  exit 2
fi

number="$(python3 -c '
import json, os, sys
try:
    d = json.load(open(os.environ["GITHUB_EVENT_PATH"]))
except (OSError, ValueError):
    sys.exit(2)
pr = d.get("pull_request") or {}
n = pr.get("number")
if not isinstance(n, int) or n <= 0:
    sys.exit(2)
print(n)
' 2>/dev/null)" || exit 2

[ -n "$number" ] || exit 2

token="${GH_TOKEN:-${GITHUB_TOKEN:-}}"
api="${GITHUB_API_URL:-https://api.github.com}"
repo="${GITHUB_REPOSITORY:-}"

fetch_via_api() {
  [ -n "$token" ] || return 1
  [ -n "$repo" ] || return 1
  command -v curl >/dev/null 2>&1 || return 1
  curl -fsSL \
    -H "Authorization: Bearer ${token}" \
    -H "Accept: application/vnd.github+json" \
    -H "X-GitHub-Api-Version: 2022-11-28" \
    "${api}/repos/${repo}/pulls/${number}" \
    | python3 -c 'import json,sys; b=json.load(sys.stdin).get("body"); sys.stdout.write(b if isinstance(b,str) else "")' \
    > "$out"
}

fetch_via_gh() {
  command -v gh >/dev/null 2>&1 || return 1
  gh pr view "$number" --json body --jq '.body // ""' > "$out"
}

if fetch_via_api || fetch_via_gh; then
  [ -s "$out" ] || exit 4
  exit 0
fi

if [ -z "$token" ] && ! command -v gh >/dev/null 2>&1; then
  echo "resolve-pr-body-by-number: FAIL: no GITHUB_TOKEN/GH_TOKEN and gh not installed" >&2
  exit 3
fi
echo "resolve-pr-body-by-number: FAIL: could not fetch body for PR #${number}" >&2
exit 4
