---
name: {{PROJECT_SLUG}}-sdd-autopilot
description: Operate {{PROJECT_SLUG}}'s interactive SDD autopilot for feature requests, evolutions, bugs, refactors and operational documentation. Use when a user asks Cursor to take an initial request through requirements, dual-spec, implementation, diagnosis, report review, gates, state updates and completion without breaking the fixed roadmap autopilot.
---

# {{PROJECT_SLUG}} SDD Autopilot

Goal: conduct an interactive feature request through Spec Driven Development until completion or a stop condition, while preserving the existing `phase_autopilot`.

## When to Use

Use this skill when the user asks to:

- implement a feature or evolution from an initial request
- turn a vague request into specs and implementation
- continue an interactive SDD trail
- diagnose or unblock an interactive SDD trail
- run the `interactive_sdd_autopilot`

Do not use this skill for the fixed roadmap execution. Use `automation/RUNBOOK.md` for `phase_autopilot`.

## Required Reading

Before editing, read:

1. `AGENTS.md`
2. `skills/00-skill-index/SKILL.md`
3. `skills/05-{{PROJECT_SLUG}}-spec-architect/SKILL.md`
4. `automation/INTERACTIVE_AUTOPILOT.md`
5. `automation/INTERACTIVE_RUNBOOK.md`
6. `automation/INTERACTIVE_STATE.json`
7. `automation/STOP_CONDITIONS.md`

Then read only the existing specs, skills and docs that govern the requested
domain. A build or diagnosis spec that does not yet exist is a candidate to
create after the spec decision, not required reading before that point.

## Non-Negotiables

1. Treat the user request as intake, not as a spec.
2. Do not implement behavior outside approved specs.
3. Every new trail must have one build spec and one diagnosis spec.
4. If a material requirement is ambiguous, ask before editing.
5. The report is the source of truth for advancement.
6. Advance only when `classification = PASS` and the decision allows the next step.
7. Stop immediately on any stop condition.
8. Do not mutate `automation/ROADMAP.json` or `automation/PHASE_STATE.json` for interactive trails.
9. Do not create product capability, `pkg/*` API, runtime feature, {{UPSTREAM_OPS_NAME}} or hosted dependency unless governed by a separate approved product spec.

## Interactive Workflow

### 1. Intake

Normalize the request:

- original request
- request type: `feature`, `evolution`, `bug`, `refactor`, `docs` or `investigation`
- intended outcome
- known scope
- ambiguous points
- candidate specs
- candidate skills

If the request is materially incomplete, ask the smallest useful set of questions and stop until answered.

### 2. Spec Decision

Use `skills/05-{{PROJECT_SLUG}}-spec-architect/SKILL.md`.

Decide:

- amend existing specs when the request belongs to an existing module or contract
- create a new dual-spec trail when the request introduces a new bounded context, primary contract or diagnostic domain

Record the decision and the governing sections.

### 3. Dual-Spec Gate

Before implementation, ensure:

- build spec defines objective, scope, non-goals, contracts or observable behavior, tests and acceptance criteria
- diagnosis spec defines signals, failure modes, commands, evidence, classification, decision and report path

If either spec is missing or incomplete, create or refine it first.

### 4. Implementation

Implement only the scoped behavior covered by the specs.

Rules:

- preserve `pkg/` and `internal/` boundaries from `specs/020-repository-architecture.md`
- update tests when behavior changes
- update docs when public or operational behavior changes
- do not change fixed roadmap files to make the interactive trail pass
- keep evidence traceable from request to specs, code, tests, docs and report

### 5. Diagnosis

Run the diagnosis spec.

Minimum checks when applicable:

- targeted tests for changed packages
- `go test ./...`
- `go test -race ./...`
- OpenAPI check when HTTP/OpenAPI changes
- documentation and rules checks when governance changes
- verification that no forbidden scope entered

If a required check cannot run, record the technical reason in the report.

### 6. Report Review

Read the generated report completely.

Extract:

- `classification`
- `decision`
- tests executed
- skipped checks
- gaps
- stop conditions

Never advance based only on tests, narrative or confidence.

### 7. Advance, Retry, Block or Complete

Advance only when:

- report exists and matches the current trail step
- `classification = PASS`
- decision explicitly allows advancement
- required checks passed or were explicitly not applicable
- no stop condition is active
- `automation/INTERACTIVE_STATE.json` is coherent

Block when the report is `PARTIAL`, `FAIL`, `BLOCKED`, missing, ambiguous or depends on human decision.

Use retries only within the configured retry limit.

## State Rules

Use `automation/INTERACTIVE_STATE.json` for interactive trails.

Never use interactive state to represent the fixed roadmap. Never use `automation/PHASE_STATE.json` to represent an interactive feature request.

When updating state, preserve:

- request id
- request type
- current step
- governing specs
- build spec
- diagnosis spec
- report path
- retry count
- latest classification
- latest decision
- blocked status and reason
- completion criteria

## Stop Conditions

Stop immediately when:

- requirements remain materially ambiguous
- build spec is missing for a new trail
- diagnosis spec is missing for a new trail
- report is missing or inconsistent
- classification is not `PASS`
- decision does not allow advancement
- tests required by the diagnosis spec fail
- implementation would exceed the approved spec
- a change would mutate fixed roadmap state to bypass a gate
- a human architecture decision is required

## Final Response Format

Always answer with:

- objetivo
- specs lidas
- skills aplicadas
- arquivos alterados
- comandos executados
- testes executados
- report lido
- classificacao
- decisao
- estado atualizado
- gaps restantes
- proxima etapa ou motivo de bloqueio

## Minimal Prompts

Start an interactive trail:

```text
Execute o interactive_sdd_autopilot para esta solicitacao.
Comece por intake, levante requisitos, decida entre amendar spec existente ou abrir nova trilha dual-spec, implemente apenas depois de specs suficientes, execute diagnostico, leia o report e avance somente quando o gate passar.
```

Continue an interactive trail:

```text
Continue o interactive_sdd_autopilot a partir de automation/INTERACTIVE_STATE.json.
Leia o estado, specs e report atuais, preserve o phase_autopilot e avance somente se o report permitir.
```
