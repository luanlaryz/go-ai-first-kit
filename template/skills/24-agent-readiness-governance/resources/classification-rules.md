# {{PROJECT_TITLE}} Agent Readiness Classification Rules

## Purpose

Filter generic `agent-readiness` findings through the scope of `{{PROJECT_SLUG}}`: a Go library/framework with SDD, automation, specs, reports, examples and public API contracts.

## Classification decision tree

### Step 1: Is the finding about agent ability to modify this repository safely?

If yes, classify as `worth_it_for_{{PROJECT_SLUG}}` unless it is already fully covered by existing evidence.

Typical examples:

- missing or stale agent instructions
- unclear architecture boundaries
- missing spec/report convention
- missing PR checklist for generated changes
- missing tests for changed public behavior
- missing CI gate for tests
- missing docs for a public or operational surface

### Step 2: Is it helpful but not required for the current change?

If yes, classify as `optional_for_{{PROJECT_SLUG}}`.

Typical examples:

- additional examples beyond minimum coverage
- optional devcontainer when normal Go setup is documented
- extra documentation polish not required for the touched surface
- expanded CI matrix when base CI is green

### Step 3: Is it app/SaaS/deployment specific?

If yes, classify as `out_of_scope_for_{{PROJECT_SLUG}}` unless the repository has a spec that explicitly brings it into scope.

Typical examples:

- production deployment platform
- cloud infrastructure monitoring
- product analytics
- frontend bundle checks
- SaaS incident-management workflows
- hosted environment operations

## Blocker policy

### Block PR

Block or request remediation when a `worth_it_for_{{PROJECT_SLUG}}` finding applies directly to changed surfaces and the PR would leave agents less able to reason about, test or safely modify the project.

Examples:

- PR changes public API but docs/OpenAPI/specs are not updated.
- PR changes automation but `AGENTS.md` or runbooks are now misleading.
- PR changes behavior with no test, smoke or diagnosis evidence.
- PR changes release, security or PR governance and leaves contradictions.

### Do not block PR

Do not block when:

- the finding is `optional_for_{{PROJECT_SLUG}}`
- the finding is `out_of_scope_for_{{PROJECT_SLUG}}`
- the finding is unrelated to the files changed
- the finding is historical and does not affect the current PR
- remediation would require a separate feature/spec not approved in the current trail

## Required rationale phrases

Use direct labels in reports:

- `BLOCKS_PR`: must be fixed or explicitly waived by spec/report.
- `DOES_NOT_BLOCK_PR`: advisory only.
- `OUT_OF_SCOPE`: should not be worked in this PR.
- `SEPARATE_TRAIL_REQUIRED`: relevant but too large for current PR.

## {{PROJECT_TITLE}}-specific relevant areas

Readiness findings are especially relevant when touching:

- `AGENTS.md`
- `.cursor/rules/*`
- `automation/*`
- `skills/*`
- `specs/*`
- `docs/*`
- `docs/reports/*`
- `.github/workflows/*`
- `README.md`
- `pkg/*` public API
- `api/openapi/*` and `api/openapi.yaml`
- `test/*`
- `examples/*`

## {{PROJECT_TITLE}}-specific non-goals

Do not introduce these merely to satisfy readiness score:

- hosted control plane
- {{UPSTREAM_OPS_NAME}}-style managed services
- SaaS dashboards
- production deployment platform
- Kubernetes deployment automation
- mandatory container workflow
- application analytics
- frontend performance tooling unless the changed area is a frontend app
