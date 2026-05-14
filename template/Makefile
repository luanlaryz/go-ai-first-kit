GO ?= go
GOFMT ?= gofmt
STATICCHECK ?= $(shell command -v staticcheck 2>/dev/null)
STATICCHECK_VERSION ?= latest
GOVULNCHECK_VERSION ?= latest
GOFILES := $(shell find . -type f -name '*.go' -not -path './vendor/*')

.PHONY: setup bootstrap test test-security security-tests vulncheck secrets-check lint fmt fmt-check race coverage vet check-compliance check-pr-body check-file-size

setup:
	@command -v $(GO) >/dev/null 2>&1 || { echo "go not found; install Go 1.26.3+ from https://go.dev/dl/"; exit 1; }
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
