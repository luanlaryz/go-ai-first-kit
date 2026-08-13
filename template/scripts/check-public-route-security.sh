#!/usr/bin/env bash
# Public route security gate: public routes must sit behind auth, tenant
# resolution and rate limiting, and the published contract must declare security
# schemes. No-ops until the project has an HTTP adapter or a contract file.
# Allowlist: .guardrails/public-route-exceptions.yaml.
set -euo pipefail
cd "$(git rev-parse --show-toplevel 2>/dev/null || pwd)"
exec go run ./tools/guardrails public-route
