#!/usr/bin/env bash
# Env read guard - preToolUse hook wrapper for Read.
# Fail-closed: any failure emits `ask`, never `allow`.
set -uo pipefail

emit_ask() {
  printf '{"permission":"ask","user_message":"Guard de .env indisponivel; decisao segura = ask."}'
  exit 0
}

command -v python3 >/dev/null 2>&1 || emit_ask
root="$(git rev-parse --show-toplevel 2>/dev/null || true)"
[ -n "$root" ] || emit_ask
[ -f "$root/.cursor/hooks/env-read-guard.py" ] || emit_ask

python3 "$root/.cursor/hooks/env-read-guard.py" || emit_ask
