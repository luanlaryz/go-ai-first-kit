#!/usr/bin/env bash
# Parallel collision guard - beforeShellExecution hook wrapper.
# Reads the hook event JSON from stdin and delegates to the Python guard.
# Fail-open: any failure emits allow so a slow network never hard-blocks a push;
# hooks.json `failClosed` covers process crash/timeout and the definitive gate is
# scripts/check-parallel-collision.sh in pre-pr/CI.
set -uo pipefail

emit_allow() { printf '{"permission":"allow"}'; exit 0; }

command -v python3 >/dev/null 2>&1 || emit_allow
root="$(git rev-parse --show-toplevel 2>/dev/null || true)"
[ -n "$root" ] || emit_allow
[ -f "$root/.cursor/hooks/parallel-collision-guard.py" ] || emit_allow

python3 "$root/.cursor/hooks/parallel-collision-guard.py" || emit_allow
