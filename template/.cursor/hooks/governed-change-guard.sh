#!/usr/bin/env bash
# Governed change guard - beforeShellExecution hook wrapper.
# Reads the hook event JSON from stdin and delegates to the Python guard.
# Fail-open: any failure emits allow so the guard never hard-blocks on its own
# bug; hooks.json `failClosed` covers process crash/timeout. The definitive gate
# is scripts/check-governed-change.sh in pre-pr/CI.
set -uo pipefail

emit_allow() { printf '{"permission":"allow"}'; exit 0; }

command -v python3 >/dev/null 2>&1 || emit_allow
root="$(git rev-parse --show-toplevel 2>/dev/null || true)"
[ -n "$root" ] || emit_allow
[ -f "$root/.cursor/hooks/governed-change-guard.py" ] || emit_allow

python3 "$root/.cursor/hooks/governed-change-guard.py" || emit_allow
