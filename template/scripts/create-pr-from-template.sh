#!/usr/bin/env bash

set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

TEMPLATE_FILE=".github/PULL_REQUEST_TEMPLATE.md"

usage() {
	cat <<'EOF'
Usage:
  scripts/create-pr-from-template.sh [gh pr create args...]

Examples:
  scripts/create-pr-from-template.sh --base main --title "docs: update report"
  scripts/create-pr-from-template.sh --base main --title "feat: x" --body-file pr-body.md

Behavior:
  - If --body-file is provided, the file is validated before calling gh pr create.
  - If --body-file is not provided, the script copies the repository PR template
    to a temporary file, opens it in $EDITOR when interactive, validates it, and
    then uses it as the PR body.
  - Passing --body is rejected to keep the workflow anchored to the repository
    template.
EOF
}

require_command() {
	command -v "$1" >/dev/null 2>&1 || {
		printf 'create-pr-from-template: FAIL: missing required command: %s\n' "$1" >&2
		exit 1
	}
}

require_command gh
require_command python3

body_file=""
gh_args=()

while [[ $# -gt 0 ]]; do
	case "$1" in
	--body)
		printf 'create-pr-from-template: FAIL: --body is not allowed; use --body-file or the repository template flow\n' >&2
		exit 1
		;;
	--body-file)
		gh_args+=("$1")
		shift
		[[ $# -gt 0 ]] || {
			printf 'create-pr-from-template: FAIL: --body-file requires a path\n' >&2
			exit 1
		}
		body_file="$1"
		gh_args+=("$1")
		;;
	-h|--help)
		usage
		exit 0
		;;
	*)
		gh_args+=("$1")
		;;
	esac
	shift
done

temp_file=""
cleanup() {
	if [[ -n "$temp_file" && -f "$temp_file" ]]; then
		rm -f "$temp_file"
	fi
}
trap cleanup EXIT

if [[ -z "$body_file" ]]; then
	temp_file="$(mktemp)"
	cp "$TEMPLATE_FILE" "$temp_file"
	body_file="$temp_file"

	if [[ -t 1 ]]; then
		editor="${EDITOR:-}"
		if [[ -z "$editor" ]]; then
			printf 'create-pr-from-template: FAIL: EDITOR is not set; pass --body-file with a filled template\n' >&2
			exit 1
		fi
		"$editor" "$body_file"
	else
		printf 'create-pr-from-template: FAIL: non-interactive use requires --body-file with a filled template\n' >&2
		exit 1
	fi

	gh_args+=("--body-file" "$body_file")
fi

./scripts/check-pr-body.sh --body-file "$body_file"
gh pr create "${gh_args[@]}"
