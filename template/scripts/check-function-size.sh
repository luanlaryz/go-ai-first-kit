#!/usr/bin/env bash
# Function size gate: warn >40, error >80 lines.
# Allowlist: .guardrails/function-size-exceptions.yaml (path + symbol).
set -euo pipefail
cd "$(git rev-parse --show-toplevel 2>/dev/null || pwd)"
exec go run ./tools/guardrails funcsize
