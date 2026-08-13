# Plan Review Verdict

Este template e a forma canonica do artefato em `automation/PLAN_REVIEWS/`. Copie e preencha quando
o veredito for escrito a mao; `scripts/review-plan.py` gera a mesma estrutura automaticamente.

## Identity

- Plan: PLAN_PATH
- Verdict artifact: VERDICT_ARTIFACT
- Review path: REVIEW_PATH
- Planner model: PLANNER_MODEL
- Reviewer model: REVIEWER_MODEL
- Reviewed at: REVIEWED_AT
- TTL days: TTL_DAYS
- Plan SHA-256: PLAN_SHA256

## Verdict

Status: APPROVED

Allowed statuses:

- `APPROVED`
- `APPROVED_WITH_CHANGES`
- `REJECTED`

## Required Changes

Record required changes when status is `APPROVED_WITH_CHANGES`. Use `None` when the status is `APPROVED` and no changes are required.

## Rejection Reasons

Record rejection reasons when status is `REJECTED`. Use `None` for approved plans.

## Operator Fallback

- Used: no
- Reason: None

## Checklist Result

- Contract coverage: PASS
- Architecture coverage: PASS
- Safety coverage: PASS
- Verification coverage: PASS

## Evidence

- Governing specs read: SPEC_LIST
- Skills applied: SKILL_LIST
- Files expected to change: FILE_LIST
- Validation commands: COMMAND_LIST
- Remaining gaps: GAP_LIST

## Append-Only Rule

After merge, do not edit this artifact. Create a versioned successor such as `PLAN_ID-v2.md` for a new review.

The plan hash must appear verbatim in this artifact: `scripts/verify-plan-review.sh` rejects a verdict that does not reference the hash it claims to approve.
