#!/usr/bin/env bash

set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

fail() {
	printf 'check-secrets: FAIL: %s\n' "$*" >&2
	exit 1
}

command -v git >/dev/null 2>&1 || fail "git not found"

tmp_file="$(mktemp)"
cleanup() {
	rm -f "$tmp_file"
}
trap cleanup EXIT

patterns=(
	'AKIA[0-9A-Z]{16}'
	'gh[pousr]_[A-Za-z0-9_]{36,}'
	'github_pat_[A-Za-z0-9_]{22,}'
	'([Aa][Pp][Ii][_-]?[Kk][Ee][Yy]|[Aa][Cc][Cc][Ee][Ss][Ss][_-]?[Tt][Oo][Kk][Ee][Nn]|[Aa][Uu][Tt][Hh][_-]?[Tt][Oo][Kk][Ee][Nn]|[Ss][Ee][Cc][Rr][Ee][Tt]|[Pp][Aa][Ss][Ss][Ww][Oo][Rr][Dd]|[Pp][Aa][Ss][Ss][Ww][Dd])[[:space:]]*[:=][[:space:]]*["'\'']?[A-Za-z0-9_./+=-]{16,}'
	'-----BEGIN (RSA |DSA |EC |OPENSSH |PGP )?PRIVATE KEY-----'
)

skip_file() {
	case "$1" in
	docs/reports/agent-readiness/*.raw.json)
		return 0
		;;
	*)
		return 1
		;;
	esac
}

findings=0

while IFS= read -r -d '' file; do
	[[ -f "$file" ]] || continue
	if skip_file "$file"; then
		continue
	fi

	for pattern in "${patterns[@]}"; do
		if LC_ALL=C grep -IEn -- "$pattern" "$file" >"$tmp_file"; then
			while IFS= read -r line; do
				printf 'check-secrets: potential secret in %s:%s\n' "$file" "$line" >&2
				findings=$((findings + 1))
			done <"$tmp_file"
		fi
	done
done < <(git ls-files -z --cached --others --exclude-standard)

if (( findings > 0 )); then
	fail "potential secrets found; remove the value, rotate it if real, and document only redacted evidence"
fi

printf 'check-secrets: ok\n'
