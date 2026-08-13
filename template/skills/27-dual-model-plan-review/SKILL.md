---
name: 27-dual-model-plan-review
description: Enforces auditable dual-model review for executable plans. Use when creating, revising, approving, or executing plans, specs, diagnostics, autopilot work, migrations, hardening, parity, security, or release-readiness tracks.
---

# Dual Model Plan Review

Goal: keep implementation plans executable only after an independent, recorded review proves the plan is within contract, current, and safe to run.

## When to Use

- Creating or editing `.cursor/plans/*.plan.md`.
- Executing a plan produced by an agent or operator.
- Planning changes to specs, ADRs, automation, migrations, published contracts, auth, security, `pkg/*` or `internal/*`.
- Running autopilot or governance tracks where plan drift could bypass gates.

## Required Reading

1. `AGENTS.md`
2. `skills/00-skill-index/SKILL.md`
3. `automation/STOP_CONDITIONS.md`
4. `automation/PLAN_REVIEWS/README.md`
5. `resources/review_contract.md`

## Non-Negotiables

- A plan starts as `DRAFT_NOT_EXECUTABLE`.
- Execution requires `APPROVED` or `APPROVED_WITH_CHANGES` in frontmatter and an artifact in `automation/PLAN_REVIEWS/`.
- `plan_sha256` must match the current plan body hash at execution time.
- Review artifacts are append-only after merge; later reviews use `-v2`, `-v3`, and so on.
- `OPERATOR_FALLBACK` is last resort and must include `operator_fallback_reason`.
- `OPERATOR_FALLBACK` is prohibited for public API surface, published contracts, migrations, auth, security, concurrency and ADR changes.
- Review state (`APPROVED`, `REJECTED`) is separate from phase diagnosis (`PASS`, `FAIL`, `BLOCKED`).

## Workflow

1. Draft the plan with status `DRAFT_NOT_EXECUTABLE`.
2. Select the strongest eligible review path:
   - `DUAL_AUTOMATED`: planner and reviewer are from different model families.
   - `DUAL_TURN_BASED`: operator switches model in a later turn.
   - `OPERATOR_FALLBACK`: human reviewer completes the checklist and records a reason.
3. Complete `resources/review_checklist.md`.
4. Write the verdict using `resources/verdict_template.md`.
5. Add or update the plan `review:` frontmatter with the same hash and artifact path.
6. Run `make verify-plan-reviews` before execution.
7. If verification fails, classify the plan as `BLOCKED` until reviewed again.

Preparation of the artifact and frontmatter is automated by `skills/28-plan-review-autopilot/SKILL.md`; it prepares the gate without approving anything on its own.

## Verdict Contract

Every executable plan frontmatter must include:

```yaml
review:
  status: APPROVED
  path: DUAL_TURN_BASED
  planner_model: "<planner-model>"
  reviewer_model: "<reviewer-model-from-another-family>"
  reviewed_at: 2026-01-31T09:00:00-03:00
  plan_sha256: 64_HEX_CHARACTERS
  verdict_artifact: automation/PLAN_REVIEWS/plan-id.md
  ttl_days: 7
```

When `path` is `OPERATOR_FALLBACK`, include:

```yaml
  operator_fallback_reason: bootstrap-do-proprio-guardrail
```

The executable gate is `scripts/verify-plan-review.sh`: it validates artifacts (frontmatter, hash, TTL, verdict file), never calls a model, and therefore needs no API key in CI.

## Do / Don't

- Do prefer a reviewer model from a different family for consequential work.
- Do reject plans with unapproved specs, missing contract updates, missing retry/DLQ requirements, or unclear tenant isolation.
- Do regenerate the hash after incorporating required changes.
- Do create a new versioned review artifact instead of editing an old verdict after merge.
- Don't execute a plan just because a human said "go" if the hash or TTL fails.
- Don't use `OPERATOR_FALLBACK` for high-risk surfaces.
- Don't collapse plan review verdicts into phase gate results.

## Checklists

Before review:

- [ ] The plan cites the governing specs or ADRs.
- [ ] The plan states touched files or module boundaries.
- [ ] The plan lists validation commands.
- [ ] The plan starts as non-executable until reviewed.

During review:

- [ ] Apply `resources/review_checklist.md`.
- [ ] Record required changes if verdict is `APPROVED_WITH_CHANGES`.
- [ ] Record rejection reasons if verdict is `REJECTED`.
- [ ] Compute and record `plan_sha256`.

After review:

- [ ] Add `review:` frontmatter to the plan.
- [ ] Store the verdict under `automation/PLAN_REVIEWS/`.
- [ ] Run `make verify-plan-reviews`.
- [ ] Include verdict path, hash, models, and gaps in the final response.

## Definition of Done

- Plan has a valid review frontmatter block.
- Verdict artifact exists and includes the same hash.
- `make verify-plan-reviews` passes.
- Any `APPROVED_WITH_CHANGES` changes are incorporated before execution.
- Final output lists skills applied, files changed, commands/tests, review verdict, and remaining gaps.

## Resources

- `resources/review_contract.md`: review paths, eligibility, and blocked surfaces.
- `resources/review_checklist.md`: mandatory review checklist.
- `resources/verdict_template.md`: copyable review artifact template.
- `resources/model_roles.md`: planner/reviewer role matrix.
