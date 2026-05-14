---
name: {{PROJECT_SLUG}}-ultra-rigid-sdd
description: use for {{PROJECT_SLUG}} spec-driven development, autopilot prompt generation, phase planning, implementation specs, diagnosis specs, reconciliation phases, release/readiness gates, and any request that must follow the user's ultra-rigid evidence-first project standard. applies to {{PROJECT_SLUG}} framework parity, documentation, conformance, roadmap closure, and future sdd phases.
---

# {{PROJECT_TITLE}} Ultra-Rigid SDD

## Core rule

Always apply the ultra-rigid model when the user asks for {{PROJECT_SLUG}} phases, specs, diagnosis specs, Cursor/Codex/autopilot prompts, roadmap validation, parity work, documentation reconciliation, readiness confirmation, release/versioning governance, or any continuation of the {{PROJECT_SLUG}} SDD process.

The default posture is: **scope locked, evidence-first, no accidental capability expansion, paired feature spec + diagnosis spec, and explicit gates.**

## Project boundary

Treat `{{PROJECT_SLUG}}` as a reusable, versionable Go framework/library for agents, sub-agents, tools, memory, workflows, HTTP exposure, observability, provider adapters, conformance, and examples.

Do not treat `{{PROJECT_SLUG}}` as the product application.

Do not put product-domain responsibilities into the framework core:
- no {{DOMAIN_ACTOR_TITLE}} Portal or {{EXTERNAL_SYSTEM}}-specific logic
- no queues or business workers
- no auth/RBAC of the final product
- no domain repositories such as {{DOMAIN_ENTITY_SET}}
- no hosted {{UPSTREAM_OPS_NAME}}/control plane work unless explicitly authorized
- no application payload formats as framework contracts unless intentionally generalized

The consumer application may use {{PROJECT_SLUG}}, but must remain responsible for external systems, queues, product auth, domain rules, and production deployment composition.

## Mandatory invocation workflow

When this skill is triggered:

1. **Restate the locked scope in one short paragraph.**
2. **Classify the request type:** feature phase, diagnosis phase, paired SDD phase, reconciliation, readiness confirmation, release/versioning, prompt generation, or audit.
3. **Check for prior completion:** if the user says a roadmap or phase is complete, do not reopen it unless evidence shows a real blocker.
4. **Generate only the smallest sufficient scope.**
5. **Always use the ultra-rigid output contract.**
6. **Include docs/demo/report updates whenever the task affects public behavior or discovery.**
7. **End with a binary gate or next-step decision.**

For detailed templates, read `references/ultra-rigid-output-contract.md` when generating prompts/specs, and read `references/sdd-phase-template.md` when creating a full SDD phase.

## Ultra-rigid defaults

### Always enforce

- Do not open capability nova without explicit user authorization.
- Do not reopen completed roadmaps by narrative speculation.
- Do not redesign runtime for reconciliation, diagnosis, polish, or readiness tasks.
- Do not mix framework readiness with application readiness.
- Do not treat examples/helpers as substitutes for public framework contracts.
- Do not treat smoke tests as proof of full parity unless the diagnosis explicitly says so.
- Do not persist a full runtime operation context automatically.
- Do not bypass docs, feature matrix, reports, or examples when affected.
- Do not skip the diagnosis spec.
- Do not skip explicit `PASS / PARTIAL / FAIL / BLOCKED` classification.

### Evidence order

Use this precedence when reasoning:

1. Exported contracts in `pkg/*`, real HTTP/OpenAPI surface, real behavior.
2. Automated tests, conformance, smoke checks, race checks, benchmarks when applicable.
3. Normative specs and diagnosis specs.
4. Reports, docs, feature matrix, README, examples.
5. Historical notes and chat context.

If lower-precedence artifacts disagree with code/tests, mark them as drift and propose reconciliation, not capability expansion.

## Required SDD pairing

For any implementation phase, produce both:

1. Feature/implementation spec.
2. Diagnosis spec.

Also produce prompts in this order:

1. Prompt to create the feature spec.
2. Prompt to implement the feature spec.
3. Prompt to create the diagnosis spec.
4. Prompt to execute the diagnosis.

Each prompt must include:
- task
- mandatory file path
- objective
- locked scope
- prohibited scope
- required artifacts
- validation gates
- exact output format
- rejection criteria

## Required phase sections

Every feature spec must include:

1. objetivo
2. motivacao
3. pergunta principal
4. escopo
5. requisitos de design
6. blocos obrigatorios da fase
7. requisitos funcionais
8. decisoes obrigatorias de modelagem
9. criterios de aceitacao
10. evidencia obrigatoria
11. fora do escopo
12. perguntas secundarias

Every diagnosis spec must include:

1. objetivo
2. pergunta principal
3. escopo
4. checklist obrigatorio
5. criterios de classificacao
6. saida obrigatoria
7. criterios de aceitacao

## Mandatory gate language

Use explicit binary decisions. Examples:

- `READY FOR PHASE N`
- `NOT READY FOR PHASE N`
- `FRAMEWORK READY FOR GO APP REWRITE`
- `FRAMEWORK NOT READY FOR GO APP REWRITE`
- `READY FOR STABLE RELEASE GOVERNANCE`
- `NOT READY FOR STABLE RELEASE GOVERNANCE`
- `ROADMAP COMPLETE`
- `ROADMAP NOT COMPLETE`

Never end with only "recommended" or "looks good".

## Validation commands

Default validation gates:

```bash
go test ./...
go test -race ./...
go vet ./...
```

Add these when applicable:

```bash
go run ./cmd/openapi-bundle -check
```

For examples/demos:

```bash
go test ./test/smoke/...
go test ./examples/...
go build ./examples/<example-name>/
```

For storage changes, require restart-proof tests.
For streaming changes, require replay/resume tests.
For public API changes, require semver impact notes.
For concurrency changes, require race checks and at least one concurrent scenario.

If a command cannot be run, require the implementing agent to state why and classify the gate as `BLOCKED` or `PARTIAL`, not `PASS`.

## Documentation and demo rules

Whenever public behavior changes, require updates to affected artifacts, typically:

- `README.md`
- `specs/010-feature-matrix.md`
- `docs/*.md`
- `docs/reports/*.md`
- `examples/*/README.md`
- reference consumer or starter example when applicable
- OpenAPI bundle when HTTP surface changes
- `CHANGELOG.md` or release docs when API surface changes

Do not allow implementation-only completion when discovery artifacts become stale.

## Rejection criteria

A generated response is invalid if it:

- opens capability nova outside the requested scope
- reopens a completed roadmap without hard evidence
- omits paired feature and diagnosis specs for an SDD phase
- omits exact file paths
- omits validation commands
- omits docs/demo/report updates when applicable
- confuses {{PROJECT_SLUG}} framework readiness with application readiness
- treats product-domain concerns as {{PROJECT_SLUG}} core
- gives a generic plan without gate language
- lacks `PASS / PARTIAL / FAIL / BLOCKED` criteria

## Standard concise response when acknowledging a completed roadmap

Use this stance:

> Treat the completed roadmap as closed. Do not reopen the main parity track. If residual issues exist, propose only a small reconciliation or readiness-confirmation phase with locked scope, evidence-first checks, and no new capability.

## Standard concise response when generating a future phase

Use this stance:

> Generate one minimal SDD phase, with feature spec + diagnosis spec + four autopilot prompts. Lock the scope, declare prohibited scope, list required artifacts, require validation gates, require docs/demo updates, and include rejection criteria.
