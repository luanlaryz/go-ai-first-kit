# Model Roles

## Planner

The planner proposes an executable path. It must:

- read governing specs, ADRs, skills, and local code;
- identify touched files and validation commands;
- keep the plan in `DRAFT_NOT_EXECUTABLE` until reviewed;
- avoid inventing contracts that are absent from the repository.

## Reviewer

The reviewer challenges the plan. It must:

- verify that the plan is inside approved contracts;
- search for gaps, contradictions, missing tests, and bypasses;
- decide `APPROVED`, `APPROVED_WITH_CHANGES`, or `REJECTED`;
- record hash, model identity, TTL, review path, and remaining gaps.

## Operator

The operator owns fallback decisions. It must:

- prefer `DUAL_AUTOMATED` when available;
- use `DUAL_TURN_BASED` when a model switch is available;
- use `OPERATOR_FALLBACK` only for bootstrap, emergency, or low-risk plans;
- reject fallback for restricted surfaces.

## Family Separation

`DUAL_AUTOMATED` requires different model families for planner and reviewer. The point is independence of failure modes: a reviewer that shares the planner's blind spots does not review, it agrees.

Same-family planner/reviewer pairs are allowed only under `DUAL_TURN_BASED` with an explicit reason in the verdict artifact. `scripts/review-plan.py` refuses an identical planner and reviewer identifier outright, because that is self-approval.
