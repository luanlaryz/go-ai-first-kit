#!/usr/bin/env bash
# Port size gate: prefer 1-3 methods per interface, warn >3, error >5.
# Allowlist: .guardrails/port-size-exceptions.yaml (path + symbol).
set -euo pipefail
cd "$(git rev-parse --show-toplevel 2>/dev/null || pwd)"
exec go run ./tools/guardrails portsize
