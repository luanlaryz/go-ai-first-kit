GO ?= go
GOFMT ?= gofmt
STATICCHECK ?= $(shell command -v staticcheck 2>/dev/null)
STATICCHECK_VERSION ?= latest
GOVULNCHECK_VERSION ?= latest
GOFILES := $(shell find . -type f -name '*.go' -not -path './vendor/*')

PLAN ?=
PR_BODY ?=

.PHONY: setup bootstrap test test-security security-tests vulncheck secrets-check lint fmt fmt-check race coverage vet check-compliance check-pr-body check-file-size check-agent-hooks hook-selftests env-keys check-sdd-compliance check-governed-change verify-plan-reviews verify-plan-reviews-all review-plan review-plan-turn-based check-parallel-collision check-parallel-collision-selftest \
	check-architecture-boundaries check-function-size check-port-size check-ignored-errors \
	check-observability check-package-clean check-public-route-security check-test-isolation \
	guardrails pre-pr

setup:
	@command -v $(GO) >/dev/null 2>&1 || { echo "go not found; install Go 1.26.4+ from https://go.dev/dl/"; exit 1; }
	$(GO) mod download
	$(GO) install honnef.co/go/tools/cmd/staticcheck@$(STATICCHECK_VERSION)

bootstrap: setup

test:
	$(GO) test ./...

test-security: security-tests secrets-check vulncheck

security-tests:
	$(GO) test ./test/security/...

vulncheck:
	$(GO) run golang.org/x/vuln/cmd/govulncheck@$(GOVULNCHECK_VERSION) ./...

secrets-check:
	./scripts/check-secrets.sh

lint:
ifeq ($(STATICCHECK),)
	@echo "staticcheck not found; install with: go install honnef.co/go/tools/cmd/staticcheck@$(STATICCHECK_VERSION)"
	@exit 1
endif
	$(STATICCHECK) ./...

fmt:
	$(GOFMT) -w $(GOFILES)

fmt-check:
	@unformatted="$$(gofmt -l $(GOFILES))"; \
	if [ -n "$$unformatted" ]; then \
		echo "unformatted Go files:"; \
		echo "$$unformatted"; \
		exit 1; \
	fi

race:
	$(GO) test -race ./...

coverage:
	$(GO) test -coverprofile=coverage.out ./...

vet:
	$(GO) vet ./...

check-compliance:
	./scripts/check-compliance.sh

check-pr-body:
	./scripts/check-pr-body.sh

check-file-size:
	./scripts/check-file-size.sh

check-agent-hooks:
	./scripts/check-agent-hooks.sh

hook-selftests:
	./.cursor/hooks/agent-cloud-access-guard-selftest.sh
	./.cursor/hooks/env-read-guard-selftest.sh

env-keys:
	./scripts/env-keys.sh

check-sdd-compliance:
	./scripts/check-sdd-compliance.sh

check-parallel-collision:
	./scripts/check-parallel-collision.sh

check-parallel-collision-selftest:
	./scripts/check-parallel-collision-selftest.sh

check-architecture-boundaries:
	./scripts/check-architecture-boundaries.sh

check-function-size:
	./scripts/check-function-size.sh

check-port-size:
	./scripts/check-port-size.sh

check-ignored-errors:
	./scripts/check-ignored-errors.sh

check-observability:
	./scripts/check-observability.sh

check-package-clean:
	./scripts/check-package-clean.sh

check-public-route-security:
	./scripts/check-public-route-security.sh

check-test-isolation:
	./scripts/check-test-isolation.sh

# Aggregated structural gates. Failure in any one fails the target: an aggregate
# that hides a failing gate is worse than no aggregate.
guardrails: check-compliance check-agent-hooks check-architecture-boundaries check-file-size \
	check-function-size check-port-size check-ignored-errors check-observability \
	check-package-clean check-public-route-security check-test-isolation \
	check-sdd-compliance check-parallel-collision check-parallel-collision-selftest \
	hook-selftests

# Mandatory before git push or PR creation. Mirrors CI plus the governance gates.
pre-pr: fmt-check lint vet guardrails test race test-security verify-plan-reviews check-governed-change
	@if [ -n "$(PR_BODY)" ]; then ./scripts/check-pr-body.sh --body-file $(PR_BODY); fi
	@echo "pre-pr: ok"

check-governed-change:
	./scripts/check-governed-change.sh $(PR_BODY)

# Diff-scoped: only plans touched by this branch must carry a valid review, so an
# unrelated plan whose TTL lapsed on the base branch does not block the PR.
verify-plan-reviews:
	./scripts/verify-plan-review.sh --changed

verify-plan-reviews-all:
	./scripts/verify-plan-review.sh --all-strict

review-plan:
	@test -n "$(PLAN)" || { echo "usage: make review-plan PLAN=.cursor/plans/<id>.plan.md"; exit 1; }
	python3 scripts/review-plan.py --plan $(PLAN)

review-plan-turn-based:
	@test -n "$(PLAN)" || { echo "usage: make review-plan-turn-based PLAN=.cursor/plans/<id>.plan.md"; exit 1; }
	python3 scripts/review-plan.py --plan $(PLAN) --mode DUAL_TURN_BASED
