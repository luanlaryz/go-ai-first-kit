#!/usr/bin/env bash
# Observability gate.
#
# Progressive by design: the stack is not imposed on an empty project. What the
# gate enforces is coherence - if a telemetry dependency is declared in go.mod,
# it must actually be used in production code, and stdlib printing is never an
# acceptable substitute for the structured logger.
#
# Target state is described in skills/12-observability-zap-otel-prom/SKILL.md.
set -euo pipefail
cd "$(git rev-parse --show-toplevel 2>/dev/null || pwd)"

fail() { printf 'check-observability: FAIL: %s\n' "$*" >&2; exit 1; }
note() { printf 'check-observability: %s\n' "$*"; }

# Production Go files under internal/: excludes tests and vendored trees.
production_files() {
	find internal -type f -name '*.go' \
		! -name '*_test.go' \
		! -path '*/vendor/*' \
		! -path '*/testdata/*' 2>/dev/null || true
}

if [ ! -d internal ]; then
	note "ok (no internal/ tree yet; nothing to instrument)"
	exit 0
fi

files="$(production_files)"
if [ -z "$files" ]; then
	note "ok (no production Go files under internal/ yet)"
	exit 0
fi

# Only DIRECT dependencies count as adoption. A transitive `// indirect` entry is
# something another library pulled in, not a decision to instrument with it, and
# demanding wiring for it would be a false positive.
declared() { grep -- "$1" go.mod 2>/dev/null | grep -qv '// indirect'; }
used() { printf '%s\n' "$files" | xargs grep -l -- "$1" >/dev/null 2>&1; }

adopted=0

# 1. Structured logger: declared means wired.
if declared 'go.uber.org/zap'; then
	used 'go.uber.org/zap' || fail "go.uber.org/zap is in go.mod but never imported in internal/** production code"
	adopted=$((adopted + 1))
fi

# 2. Tracing: declared means instrumented, and the HTTP edge must be wrapped
#    exactly once rather than per handler.
if declared 'go.opentelemetry.io'; then
	used 'go.opentelemetry.io' || fail "OpenTelemetry is in go.mod but not used in internal/** (tracing instrumentation required)"
	adopted=$((adopted + 1))
	if declared 'otelhttp' || used 'otelhttp'; then
		used 'otelhttp' || fail "otelhttp declared but not used; wrap the HTTP handler at the edge, not per route"
	fi
fi

# 3. Metrics: declared means a metrics surface exists.
if declared 'github.com/prometheus/client_golang'; then
	used 'prometheus' || fail "Prometheus client is in go.mod but no metrics are registered in internal/**"
	adopted=$((adopted + 1))
fi

# 4. No stdlib log / fmt printing in internal/** production code. This applies
#    regardless of which telemetry stack the project adopted: unstructured output
#    is invisible to any of them.
pattern='(log\.(Printf|Println|Print|Fatal|Fatalf|Fatalln|Panic)|fmt\.(Print|Println|Printf))\('
offenders="$(printf '%s\n' "$files" | xargs grep -En -- "$pattern" 2>/dev/null || true)"
if [ -n "$offenders" ]; then
	printf '%s\n' "$offenders" >&2
	fail "stdlib log/fmt printing found in internal/** production code; use the structured logger"
fi

if [ "$adopted" -eq 0 ]; then
	note "ok (no telemetry dependency declared yet; stdlib printing check passed)"
else
	note "ok ($adopted telemetry layer(s) declared and wired; no stdlib printing)"
fi
