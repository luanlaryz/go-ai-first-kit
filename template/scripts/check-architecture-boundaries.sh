#!/usr/bin/env bash
# Hexagonal boundary gate. Production-only, and every root is optional:
# domain imports stdlib only; application forbids adapters/frameworks/SDKs and
# concrete telemetry; external SDKs only in adapters or the composition root;
# pkg/ must not depend on internal/ except through the declared public bridge.
set -euo pipefail
cd "$(git rev-parse --show-toplevel 2>/dev/null || pwd)"
exec go run ./tools/guardrails architecture
