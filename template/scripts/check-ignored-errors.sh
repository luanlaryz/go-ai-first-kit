#!/usr/bin/env bash
# Ignored error gate: `_ = call(...)` is rejected when the callee name suggests
# audit, persistence, queue or IO work, where discarding the error hides failure.
# Allowlist: .guardrails/ignored-error-exceptions.yaml (path + symbol).
set -euo pipefail
cd "$(git rev-parse --show-toplevel 2>/dev/null || pwd)"
exec go run ./tools/guardrails ignored-errors
