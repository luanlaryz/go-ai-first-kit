#!/usr/bin/env bash
# Agent cloud access guard - beforeShellExecution hook wrapper.
# Reads the hook event JSON from stdin and delegates to the Python classifier.
# Fail-closed: any failure emits `ask`, never `allow`. This is the opposite of
# governed-change-guard.sh, which is fail-open because its definitive gate is CI;
# here there is no downstream gate, so a broken guard must not open production.
set -uo pipefail

emit_ask() {
  printf '{"permission":"ask","user_message":"Guard de credencial indisponivel; decisao segura = ask."}'
  exit 0
}

command -v python3 >/dev/null 2>&1 || emit_ask
root="$(git rev-parse --show-toplevel 2>/dev/null || true)"
[ -n "$root" ] || emit_ask
[ -f "$root/.cursor/hooks/agent-cloud-access-guard.py" ] || emit_ask

CURSOR_PROJECT_ROOT="$root" python3 "$root/.cursor/hooks/agent-cloud-access-guard.py" || emit_ask
