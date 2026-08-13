---
name: 28-plan-review-autopilot
description: Prepares auditable plan review artifacts and review frontmatter without bypassing the dual-model review gate. Use when a `.cursor/plans/*.plan.md` file is blocked by missing or invalid `review:` metadata.
---

# Plan Review Autopilot

Goal: help the agent prepare the plan review gate while preserving the contract in `skills/27-dual-model-plan-review/SKILL.md`.

## When To Use

- A plan in `.cursor/plans/*.plan.md` is blocked because `review:` is missing, stale, expired, or invalid.
- The user asks to prepare, generate, or repair review artifacts.
- You need to explain why a plan remains `DRAFT_NOT_EXECUTABLE`.

## Required Reading

1. `AGENTS.md`
2. `skills/27-dual-model-plan-review/SKILL.md`
3. `skills/27-dual-model-plan-review/resources/review_contract.md`
4. `automation/PLAN_REVIEWS/README.md`

## Workflow

1. Confirm the plan path and that the user wants review preparation, not functional execution.
2. Run:

```bash
make review-plan PLAN=.cursor/plans/example.plan.md
```

3. If no reviewer API is configured, use the safe turn-based preparation:

```bash
make review-plan-turn-based PLAN=.cursor/plans/example.plan.md
```

4. Explain that a pending artifact is not approval. The plan remains blocked until a different reviewer model/operator reviews the final plan body.
5. After independent turn-based review, record it explicitly:

```bash
python3 scripts/review-plan.py --plan .cursor/plans/example.plan.md --mode DUAL_TURN_BASED --operator-reviewed --planner-model "<planner-model>" --reviewer-model "<different-model>"
```

6. Validate before execution:

```bash
make verify-plan-reviews
```

## Mode Rules

- `DUAL_AUTOMATED`: requires `PLAN_REVIEW_API_KEY`, `PLAN_REVIEW_PLANNER_MODEL`, and a different `PLAN_REVIEW_REVIEWER_MODEL`. The script refuses to fake approval when provider automation is unavailable.
- `DUAL_TURN_BASED`: safe default when there is no reviewer API. Creates a pending artifact unless `--operator-reviewed` and a different reviewer model are supplied.
- `OPERATOR_FALLBACK`: last resort only. It is blocked for high-risk plans and requires `--operator-fallback-reason`.

CI never needs `PLAN_REVIEW_API_KEY`: the gate `scripts/verify-plan-review.sh` validates artifacts, it does not produce them.

## Explain It Simply

Tell newer contributors: the script prepares the paperwork for review. It does not make an unsafe plan executable by itself. A plan becomes executable only when the review block, artifact, hash, TTL, and reviewer evidence all pass verification.

Use clear messages:

- "Este plano ainda nao e executavel."
- "O review nao foi criado porque falta reviewer model independente."
- "Para seguir com seguranca, rode este comando..."
- "Nao use OPERATOR_FALLBACK neste caso porque o plano toca area de alto risco."
- "O hash mudou; isso significa que o plano foi alterado depois do review."

## Non-Negotiables

- Do not change the plan body while preparing review.
- Do not execute a functional plan before review verification passes.
- Do not use the same model as planner and reviewer.
- Do not use `OPERATOR_FALLBACK` for public API surface, published contracts, migrations, auth, security, secrets, concurrency, queues, DLQ, replay, idempotency, ordering, or ADRs.
- Do not remove or weaken `scripts/verify-plan-review.sh`.

## Definition Of Done

- Artifact exists under `automation/PLAN_REVIEWS/`.
- Approved plans have valid `review:` frontmatter and unchanged body hash.
- Pending artifacts are clearly marked as non-executable.
- `make verify-plan-reviews` passes before any plan execution.
- Final response lists files changed, commands run, results, risks, next steps, applied skills, and gaps.
