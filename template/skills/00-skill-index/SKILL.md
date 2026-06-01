---
name: {{PROJECT_SLUG}}-{{PROJECT_SLUG}}-go-skill-index
description: Entry point for Go agents working on {{PROJECT_SLUG}}. Use whenever you need to know which specialized skill governs a change or to verify global guardrails before starting or finishing any task.
---

# Codex Go Skill Index
Goal: ensure every change applies the right specialized skill and respects global guardrails for {{PROJECT_SLUG}}.

## When to Use
- Always load this index before starting work to route yourself to the correct specialized skill.
- Re-check before code review or PR merge to confirm no relevant skill was missed.

## Non-negotiables
1. **Spec Driven Development**: every feature traces to a spec in `specs/`; never implement behavior without normative coverage.
2. **Architecture Boundaries**: `pkg/` contains stable public contracts; `internal/` contains runtime details; `pkg/app` is the only bridge to `internal/runtime`.
3. **Context Propagation**: propagate `context.Context` as the standard boundary for execution, cancellation and lifecycle.
4. **Effective Go**: follow the 13-go skill for context, errors, interfaces, concurrency, and tests.
5. **Observability**: emit events, logs and hooks through public contracts; never leak internal types into public signatures.

## Do / Don't
- **Do**: Identify the skill ID(s) below that match the asset you modify and load them fully.
- **Do**: Record in PR description which skills were followed.
- **Don't**: Mix patterns from multiple skills without reconciling terminology.
- **Don't**: Introduce new libs or infra outside these skills without design review.

## Interfaces / Contracts
- Use this table to map work areas to skills:
  - `01-hexagonal-architecture`: layered architecture, ports, adapters, composition root wiring.
  - `02-tenant-auth-quotas`: multi-tenant context, auth, rate limits (inherited from infra context -- see note in skill).
  - `03-async-command-processing`: async patterns, command queues, workers (inherited from infra context).
  - `04-streaming-sse-ws`: streaming, SSE, WebSocket replay (inherited from infra context).
  - `05-{{PROJECT_SLUG}}-spec-architect`: spec refinement, dual-spec workflow, prompt generation for {{PROJECT_SLUG}}.
  - `06-multi-provider-model-routing`: model routing, provider config (inherited from infra context).
  - `07-postgres-pgx-migrate`: pgx usage, migrations (inherited from infra context).
  - `08-redis-cache-streams`: caching, locks, streams (inherited from infra context).
  - `09-sqs-fifo-idempotency`: FIFO queues, dedup, DLQ (inherited from infra context).
  - `10-http-gin-openapi`: REST handlers, middleware, OpenAPI (inherited from infra context).
  - `11-grpc-interceptors-health`: gRPC services, interceptors (inherited from infra context).
  - `12-observability-zap-otel-prom`: logging, tracing, metrics (inherited from infra context).
  - `13-go-idiomatic-effective-go`: Go idioms, context, errors, concurrency, tests, file size guidelines.
  - `14-config-viper-flags`: config loading, feature flags (inherited from infra context).
  - `15-devx-ci-precommit-changelog`: tooling, Makefile, CI, pre-commit, changelog.
  - `16-object-calisthenics`: reduce nesting, use strong types, keep functions small in domain and application layers.
  - `17-solid-go-ports`: SOLID principles with ports/adapters and strict dependency direction.
  - `18-security-owasp-api`: API security, input validation, secrets (inherited from infra context).
  - `19-prompt-injection-llm-safety`: LLM safety, prompt injection defense (inherited from infra context).
  - `20-testing-strategy-regression-load-containers`: testing pyramid, regression, conformance, contracts.
  - `21-documentation-open-source`: godoc, README, CONTRIBUTING, CHANGELOG, examples.
  - `22-release-versioning-governance`: SemVer, release readiness, changelog, release notes and distribution path.
  - `23-{{PROJECT_SLUG}}-sdd-autopilot`: interactive SDD autopilot for feature requests, dual-spec trails, diagnosis, reports and gates.
  - `24-agent-readiness-governance`: `@kodus/agent-readiness` analysis filtered for {{PROJECT_SLUG}}'s Go library/framework scope.
  - `25-{{PROJECT_SLUG}}-ultra-rigid-sdd`: ultra-rigid evidence-first SDD for phase planning, autopilot prompt generation, paired implementation and diagnosis specs, reconciliation, readiness/release gates, and framework-vs-application boundaries.
  - `26-backlog-item-intake`: governed backlog intake for requests, gaps, bugs, diagnostics, recommendations and ideas persisted in `docs/backlog/Backlog.md`.

## Checklists
**Before starting**
- [ ] Identify change scope and map to skill IDs above.
- [ ] Load each skill and required resource files.
- [ ] Confirm terminology alignment with `AGENTS.md` and relevant specs.

**During work**
- [ ] Keep architecture boundaries intact (`pkg/` vs `internal/` per `specs/020`).
- [ ] Apply Effective Go and observability patterns consistently.
- [ ] Respect file size guidelines from skill 13 and `specs/020` section 8.

**Before PR**
- [ ] Run relevant tests/linters per skill instructions.
- [ ] Update docs/diagrams referenced by skills if behavior changed.
- [ ] Capture skill checklist status in PR template.
- [ ] If touching core behavior, apply skill `20-testing-strategy-regression-load-containers`.

## Definition of Done
- Referenced skills applied with evidence (code/comments/tests) for every touched subsystem.
- No contradictions with non-negotiables.
- PR checklist attached and green.
- Documentation updated (or explicitly not needed).

## Minimal Examples
- Example PR note: `Skills: 01-hexagonal-architecture, 13-go-idiomatic-effective-go, 20-testing-strategy`. Include checklist results inline.
- Example routing decision: modifying `pkg/agent` + `internal/runtime` => load skills 13 + 17 + 20.
- Example testing routing: modifying conformance suites + memory backend => load skills 20 + 21.
- Example interactive request: user asks for a new feature from free text => load skills 05 + 23, then any domain-specific skills.
- Example readiness review: user asks to run or apply `@kodus/agent-readiness` => load skills 20 + 21 + 24, and skill 23 when it is part of an interactive trail.
- Example ultra-rigid phase planning: user asks for {{PROJECT_SLUG}} phase prompts, paired SDD specs, diagnosis gates, roadmap closure, or readiness confirmation => load skills 05 + 25, and skill 23 when it is part of an interactive trail.
- Example backlog intake: user asks to register a gap, bug, recommendation or diagnostic finding in the backlog => load skills 05 + 23 + 25 + 26 before editing `docs/backlog/Backlog.md`.
