# Review Contract

## Review Paths

### DUAL_AUTOMATED

Use when the environment can run planner and reviewer models independently.

Requirements:

- Planner and reviewer must be from different model families.
- Reviewer must inspect the plan, governing specs, relevant skills, and validation commands.
- Verdict artifact must include the reviewer model identifier.

Required for:

- `pkg/**` (public API surface)
- `internal/**` behavior changes
- published contracts under `api/**`
- `migrations/**`
- auth, security and secret handling changes
- concurrency, queue, DLQ, replay and ordering semantics
- ADRs in `docs/decisions/`
- release-readiness and hardening tracks

### DUAL_TURN_BASED

Use when the operator can switch models in the IDE or orchestration layer.

Requirements:

- Planner model and reviewer model must be recorded.
- Reviewer must use the final plan content and current hash.
- The verdict must state which changes were required or why none were required.

### OPERATOR_FALLBACK

Use only when dual-model execution is unavailable and the plan is low risk, or during bootstrap of this guardrail.

Requirements:

- `operator_fallback_reason` must be present.
- Human reviewer must complete `review_checklist.md`.
- Review artifact must state why a model reviewer was unavailable.

Prohibited for:

- public API surface in `pkg/**`;
- published contracts under `api/**`;
- migrations and persisted data;
- auth, security, secrets, or tenant isolation;
- concurrency, queue semantics, DLQ, replay, ordering and idempotency;
- ADR creation or modification.

`scripts/review-plan.py` refuses `OPERATOR_FALLBACK` when the plan text matches any high-risk marker, so this list is enforced, not merely documented.

## Verdict Semantics

- `APPROVED`: executable while hash and TTL remain valid.
- `APPROVED_WITH_CHANGES`: executable only after required changes are incorporated and hash is recalculated.
- `REJECTED`: not executable; create a revised plan and new review.
- `DRAFT_NOT_EXECUTABLE`: never executable.

## Drift Rule

Any plan body change after review invalidates the verdict. The next reviewer must compare the old and new plan content, then issue a new versioned artifact.

## TTL Rule

Default TTL is 7 days. A TTL longer than 7 days and up to 30 days must include a risk justification in the verdict artifact. A TTL longer than 30 days is not valid.
