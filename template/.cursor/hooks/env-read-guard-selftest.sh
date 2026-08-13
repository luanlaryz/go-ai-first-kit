#!/usr/bin/env bash
# Selftest for the env read guard.
set -uo pipefail

root="$(git rev-parse --show-toplevel 2>/dev/null || pwd)"
GUARD="$root/.cursor/hooks/env-read-guard.sh"
failures=0

command -v python3 >/dev/null 2>&1 || { echo "python3 is required for this selftest" >&2; exit 1; }
[ -x "$GUARD" ] || { echo "guard not executable: $GUARD" >&2; exit 1; }

permission() {
  python3 -c 'import json,sys
try:
    print(json.load(sys.stdin).get("permission", "<none>"))
except Exception:
    print("<invalid>")'
}

decide_path() {
  python3 -c 'import json,sys; print(json.dumps({"tool_name":"Read","tool_input":{"path": sys.argv[1]}}))' "$1" \
    | "$GUARD" 2>/dev/null \
    | permission
}

expect() {
  local label="$1" want="$2" path="$3"
  local got
  got="$(decide_path "$path")"
  if [ "$got" = "$want" ]; then
    printf 'ok   %-52s -> %s\n' "$label" "$got"
  else
    printf 'FAIL %-52s -> got=%s want=%s\n' "$label" "$got" "$want" >&2
    failures=$((failures + 1))
  fi
}

expect ".env na raiz requer aprovacao" ask ".env"
expect ".env em subdiretorio requer aprovacao" ask "/abs/path/test/integration/.env"
expect ".env.local requer aprovacao" ask ".env.local"
expect ".env.production requer aprovacao" ask "deploy/.env.production"
expect ".env.example liberado" allow "test/integration/.env.example"
expect ".env.sample liberado" allow "config/.env.sample"
expect ".env.template liberado" allow ".env.template"
expect "codigo Go liberado" allow "internal/config/config.go"
expect "arquivo com env no nome liberado" allow "internal/adapters/secrets/environment.go"
expect "runbook liberado" allow "docs/runbooks/agent-cloud-access.md"

bad="$(printf 'nao json' | "$GUARD" 2>/dev/null | permission)"
if [ "$bad" = "ask" ]; then
  printf 'ok   %-52s -> %s\n' "fail-closed em evento invalido" "$bad"
else
  printf 'FAIL %-52s -> got=%s want=ask\n' "fail-closed em evento invalido" "$bad" >&2
  failures=$((failures + 1))
fi

if [ "$failures" -ne 0 ]; then
  echo "env-read-guard-selftest: FAIL ($failures caso(s))" >&2
  exit 1
fi
echo "env-read-guard-selftest: ok"
