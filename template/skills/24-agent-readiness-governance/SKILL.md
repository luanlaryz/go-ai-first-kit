---
name: agent-readiness-governance
description: Govern agent-readiness analysis for {{PROJECT_SLUG}} feature, refactor and PR workflows. Use when running, reading or applying @kodus/agent-readiness results, deciding which findings matter for a Go library/framework, creating readiness reports, attaching readiness evidence to PRs, updating docs, or deciding whether readiness regressions should block a change.
---

# Agent Readiness Governance

Goal: use `@kodus/agent-readiness` as an evidence input for `{{PROJECT_SLUG}}` without turning generic app/SaaS checks into inappropriate PR blockers for a Go library/framework.

## Required reading

Before applying this skill, read:

1. `AGENTS.md`
2. `skills/00-skill-index/SKILL.md`
3. `skills/21-documentation-open-source/SKILL.md`
4. `skills/20-testing-strategy-regression-load-containers/SKILL.md`
5. `skills/23-{{PROJECT_SLUG}}-sdd-autopilot/SKILL.md` when the work is part of an interactive autopilot trail
6. `docs/interactive-sdd-autopilot.md` when the work is part of an interactive autopilot trail
7. the current `agent-readiness` report if one already exists

When a readiness assessment requires a dedicated build or diagnosis spec that
does not exist yet, treat those paths as candidates to create through the
normal dual-spec decision. Do not require them as pre-existing reading.

If release, CI, security or OpenAPI files are changed, also route through the relevant domain skill from `skills/00-skill-index/SKILL.md`.

## Non-negotiables

1. Treat `agent-readiness` as an advisory and evidence source, not a universal gate.
2. Filter every finding through the `{{PROJECT_SLUG}}` context: Go library/framework, public API in `pkg/*`, docs/specs/reports, automation and agent workflows.
3. Do not block PRs on checks that are irrelevant to a library/framework.
4. Do block or request remediation when a change degrades agent operability in a relevant area.
5. Always classify findings as `worth_it_for_{{PROJECT_SLUG}}`, `optional_for_{{PROJECT_SLUG}}` or `out_of_scope_for_{{PROJECT_SLUG}}`.
6. Always produce a readiness report artifact for feature/refactor PRs when this skill is invoked.
7. Do not let a high raw score hide a relevant regression.
8. Do not let a low raw score force out-of-scope work.
9. When a finding is rejected as not relevant, record the reason.
10. When this skill is part of an interactive autopilot trail, the readiness report is supporting evidence; the phase/diagnosis report remains the gate source of truth.

## Classification rules

### `worth_it_for_{{PROJECT_SLUG}}`

Classify a finding as `worth_it_for_{{PROJECT_SLUG}}` when it improves one or more of:

- agent ability to understand repository purpose, architecture or constraints
- SDD/dual-spec workflow clarity
- Cursor/agent rules and automation reliability
- PR review quality and reproducibility
- testability, quick-check coverage, conformance or regression safety
- CI checks for library quality
- security hygiene relevant to a Go library/framework
- open-source contributor onboarding
- public docs, examples and API discoverability
- explicit evidence for generated changes

These findings may block a PR when the PR introduces or worsens the relevant deficiency.

### `optional_for_{{PROJECT_SLUG}}`

Classify a finding as `optional_for_{{PROJECT_SLUG}}` when it can help but should not block normal PRs unless the PR directly changes that area. Examples:

- additional devcontainer quality when local setup already works
- broader CI matrix beyond the current supported baseline
- dependency update automation beyond minimal safety
- extra templates or governance docs that improve contributor experience but are not required for the current change
- extra examples when existing examples already prove the modified capability

### `out_of_scope_for_{{PROJECT_SLUG}}`

Classify a finding as `out_of_scope_for_{{PROJECT_SLUG}}` when it is primarily for an application, SaaS product or deployed service rather than a Go library/framework. Do not block PRs on these unless the project explicitly opens a spec making them relevant.

Common non-blocking/out-of-scope examples:

- production deployment pipeline requirements
- hosting platform readiness
- app runtime monitoring dashboards outside local observability
- bundle-size or frontend asset analysis when no frontend product is changed
- product analytics instrumentation
- Kubernetes/infra hardening unrelated to library usage
- cloud deployment checks
- SaaS incident response processes
- Docker/devcontainer requirements when the repo already has a documented Go-only local flow and no containerized workflow is in scope

## Required workflow

### 1. Collect inputs

Gather:

- change summary or user request
- touched files
- current branch/PR context when available
- existing `agent-readiness` report, or command output if a new run was performed
- relevant specs, reports, docs and automation files

If no report exists and the task explicitly requires readiness assessment, run or request a report using the repository-approved method. Prefer JSON output if available, but do not depend on JSON if text is the only available artifact.

### 2. Normalize the report

Produce a normalized list of findings with:

- id or stable name
- original finding text
- source category/pillar if available
- severity/score impact if available
- affected files or repository areas
- initial classification
- final `{{PROJECT_SLUG}}` classification
- blocker status
- rationale

Use `resources/classification-rules.md`.

### 3. Decide PR impact

A finding may block a PR only when all are true:

- classification is `worth_it_for_{{PROJECT_SLUG}}`
- the PR introduces, worsens or leaves unaddressed a relevant deficiency in files it touches
- remediation is feasible inside the current scope or the PR must explicitly defer it via spec/report
- the finding is not contradicted by stronger repository evidence

A finding must not block a PR when:

- classification is `optional_for_{{PROJECT_SLUG}}` or `out_of_scope_for_{{PROJECT_SLUG}}`
- it is unrelated to changed surfaces
- it requires product/SaaS/deployment work outside repository scope
- it conflicts with existing non-goals or phase scope

### 4. Generate report artifacts

For every readiness assessment, create or update a report under one of:

- `docs/reports/agent-readiness/<YYYY-MM-DD>-<short-scope>.md`
- `docs/reports/agent-readiness/<pr-number-or-branch>.md`

If the repo does not yet have the directory, create it.

The report must use `resources/report-template.md`.

### 5. Update PR and documentation

When operating in a PR workflow:

- include a PR section using `resources/pr-template-snippet.md`
- attach or link the readiness report path
- list only relevant blockers
- list ignored findings with rationale

Update project docs only when:

- a recurring relevant gap is discovered
- a user-facing contributor workflow changes
- agent/autopilot rules change
- README/AGENTS/docs are contradicted by the readiness result

### 6. Feed autopilot decisions carefully

For interactive SDD autopilot:

- readiness findings can create a new improvement trail only if classified `worth_it_for_{{PROJECT_SLUG}}`
- readiness findings cannot override a phase report gate
- readiness findings cannot force work outside the active feature/refactor scope unless the user approves a new trail
- readiness report should be referenced as supporting evidence in the feature/refactor diagnosis report

## Required output

When this skill is used, respond with these sections:

1. **Readiness source**
2. **Scope interpreted for {{PROJECT_SLUG}}**
3. **Classified findings**
4. **PR blockers**
5. **Non-blocking recommendations**
6. **Out-of-scope findings ignored**
7. **Report artifact path**
8. **Docs/PR updates required**
9. **Autopilot impact**

## References

- Use `resources/classification-rules.md` for filtering rules.
- Use `resources/report-template.md` for the generated report shape.
- Use `resources/pr-template-snippet.md` for PR text.
