GO ?= go
GOFMT ?= gofmt
STATICCHECK ?= $(shell command -v staticcheck 2>/dev/null)
VERSION ?= dev
COMMIT ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo none)
DATE ?= $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
LDFLAGS := -X github.com/inventa-co/go-ai-first-kit/internal/version.Version=$(VERSION) -X github.com/inventa-co/go-ai-first-kit/internal/version.Commit=$(COMMIT) -X github.com/inventa-co/go-ai-first-kit/internal/version.Date=$(DATE)
GOFILES := $(shell find . -type f -name '*.go' -not -path './vendor/*' -not -path './template/*')
PACKAGES := . ./cmd/... ./internal/...

.PHONY: build install test lint vet fmt fmt-check dist

build:
	$(GO) build -ldflags "$(LDFLAGS)" -o bin/gakit ./cmd/gakit

install:
	$(GO) install -ldflags "$(LDFLAGS)" ./cmd/gakit

test:
	$(GO) test $(PACKAGES)

lint:
ifeq ($(STATICCHECK),)
	@echo "staticcheck not found; install with: go install honnef.co/go/tools/cmd/staticcheck@latest"
	@exit 1
endif
	$(STATICCHECK) ./...

vet:
	$(GO) vet $(PACKAGES)

fmt:
	$(GOFMT) -w $(GOFILES)

fmt-check:
	@unformatted="$$(gofmt -l $(GOFILES))"; \
	if [ -n "$$unformatted" ]; then \
		echo "unformatted Go files:"; \
		echo "$$unformatted"; \
		exit 1; \
	fi

dist: build
