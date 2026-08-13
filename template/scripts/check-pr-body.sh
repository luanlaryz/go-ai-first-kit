#!/usr/bin/env bash

set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

fail() {
	printf 'check-pr-body: FAIL: %s\n' "$*" >&2
	exit 1
}

usage() {
	cat <<'EOF'
Usage:
  scripts/check-pr-body.sh --body-file <path>
  scripts/check-pr-body.sh < body.md

When running in GitHub Actions on pull_request events, the script also reads the
PR body automatically from GITHUB_EVENT_PATH.
EOF
}

is_dependabot_pull_request_event() {
	local event_path="$1"

	python3 - "$event_path" <<'PY'
import json
import sys

with open(sys.argv[1], "r", encoding="utf-8") as fh:
    payload = json.load(fh)

pull_request = payload.get("pull_request") or {}
user_login = (pull_request.get("user") or {}).get("login", "")
sender_login = (payload.get("sender") or {}).get("login", "")
head_ref = (pull_request.get("head") or {}).get("ref", "")

if (
    user_login in {"dependabot[bot]", "app/dependabot"}
    or sender_login == "dependabot[bot]"
    or head_ref.startswith("dependabot/")
):
    sys.exit(0)

sys.exit(1)
PY
}

body_file=""
temp_file=""

cleanup() {
	if [[ -n "$temp_file" && -f "$temp_file" ]]; then
		rm -f "$temp_file"
	fi
}
trap cleanup EXIT

while [[ $# -gt 0 ]]; do
	case "$1" in
	--body-file)
		shift
		[[ $# -gt 0 ]] || fail "--body-file requires a path"
		body_file="$1"
		;;
	-h|--help)
		usage
		exit 0
		;;
	*)
		fail "unknown argument: $1"
		;;
	esac
	shift
done

if [[ -n "$body_file" ]]; then
	[[ -f "$body_file" ]] || fail "body file not found: $body_file"
elif [[ -n "${GITHUB_EVENT_PATH:-}" && -f "${GITHUB_EVENT_PATH:-}" ]]; then
	if is_dependabot_pull_request_event "$GITHUB_EVENT_PATH"; then
		printf 'check-pr-body: skipped for Dependabot generated dependency PR\n'
		exit 0
	fi

	temp_file="$(mktemp)"
	python3 - "$GITHUB_EVENT_PATH" "$temp_file" <<'PY'
import json
import pathlib
import sys

event_path = pathlib.Path(sys.argv[1])
out_path = pathlib.Path(sys.argv[2])
with event_path.open("r", encoding="utf-8") as fh:
    payload = json.load(fh)
body = payload.get("pull_request", {}).get("body", "")
out_path.write_text(body, encoding="utf-8")
PY
	body_file="$temp_file"
elif [[ ! -t 0 ]]; then
	temp_file="$(mktemp)"
	cat >"$temp_file"
	body_file="$temp_file"
else
	usage
	fail "no PR body source provided"
fi

[[ -s "$body_file" ]] || fail "PR body is empty"

required_headings=(
	"## Objetivo"
	"## Backlog"
	"## Specs lidas"
	"## Autopilot"
	"## Regressao"
	"## Skills aplicadas"
	"## Arquivos alterados"
	"## Impacto em docs e superficie publica"
	"## Comandos executados"
	"## Testes executados"
	"## Excecoes formais"
	"## Gaps restantes"
	"## Checklist"
)

for heading in "${required_headings[@]}"; do
	grep -Fqx "$heading" "$body_file" || fail "missing required heading: $heading"
done

forbidden_literals=(
	"Descreva o resultado observavel desta mudanca."
	"- [ ] \`AGENTS.md\`"
	"- [ ] spec(s) governante(s) citadas abaixo"
	"- \`specs/...\`"
	"- \`skills/...\`"
	"- \`BLG-NNNN\` (\`docs/backlog/Backlog.md\`), com \`Status\` em \`ready_for_implementation\`, \`in_progress\` ou \`done\`"
	"- Plano: \`.cursor/plans/<id>.plan.md\`"
	"- \`BLG-NNNN\` + cenario alvo em \`test/<pacote>::<cenario>\`"
	"- liste os arquivos principais alterados neste PR"
	"- \`none\`, ou descreva o impacto relevante quando existir"
	"- liste comandos executados alem dos checks padrao, quando houver"
	"- descreva gaps, tradeoffs ou itens fora do escopo"
)

for literal in "${forbidden_literals[@]}"; do
	if grep -Fqx -- "$literal" "$body_file"; then
		fail "template placeholder not replaced: $literal"
	fi
done

printf 'check-pr-body: ok\n'
