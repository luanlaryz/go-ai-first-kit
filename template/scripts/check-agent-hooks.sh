#!/usr/bin/env bash
# Agent hooks wiring gate.
#
# A security control that can be silently disabled is not a control: this gate
# fails when .cursor/hooks.json loses an expected guard, drops failClosed, or
# points at a script that is missing or not executable.
#
# Uses python3 rather than jq because python3 is already required by the guards
# themselves, so this gate adds no new dependency.
set -euo pipefail

HOOKS_FILE="${1:-.cursor/hooks.json}"

command -v python3 >/dev/null 2>&1 || {
  echo "check-agent-hooks: FAIL: python3 is required" >&2
  exit 1
}
[ -f "$HOOKS_FILE" ] || {
  echo "check-agent-hooks: FAIL: $HOOKS_FILE not found" >&2
  exit 1
}

python3 - "$HOOKS_FILE" <<'PY'
import json
import os
import sys

hooks_file = sys.argv[1]
# Hook commands are declared relative to the project root, which is the parent
# of the .cursor directory holding hooks.json.
root = os.path.dirname(os.path.dirname(os.path.abspath(hooks_file)))

# (event, script fragment, must declare failClosed)
REQUIRED = (
    ("beforeShellExecution", "agent-cloud-access-guard.sh", True),
    ("preToolUse", "env-read-guard.sh", True),
    # The guard itself is fail-open by design (its definitive gate is CI), but the
    # hook entry must still declare failClosed so a crash or timeout does not pass
    # unnoticed.
    ("beforeShellExecution", "governed-change-guard.sh", True),
    ("beforeShellExecution", "parallel-collision-guard.sh", True),
)

try:
    with open(hooks_file, encoding="utf-8") as handle:
        config = json.load(handle)
except json.JSONDecodeError as exc:
    print(f"check-agent-hooks: FAIL: {hooks_file} is not valid JSON ({exc})", file=sys.stderr)
    raise SystemExit(1)

hooks = config.get("hooks") or {}
failures = []

for event, script, want_fail_closed in REQUIRED:
    entries = [item for item in (hooks.get(event) or []) if script in str(item.get("command", ""))]
    if not entries:
        failures.append(f"{event} is missing hook {script}")
        continue

    entry = entries[0]
    if want_fail_closed and entry.get("failClosed") is not True:
        failures.append(f"{script} must declare failClosed: true")

    resolved = os.path.join(root, entry.get("command", ""))
    if not os.path.isfile(resolved):
        failures.append(f"hook script not found: {entry.get('command')}")
    elif not os.access(resolved, os.X_OK):
        failures.append(f"hook script not executable: {entry.get('command')}")

if failures:
    for failure in failures:
        print(f"check-agent-hooks: FAIL: {failure}", file=sys.stderr)
    print(f"check-agent-hooks: FAIL ({len(failures)} problema(s))", file=sys.stderr)
    raise SystemExit(1)

print("check-agent-hooks: ok")
PY
