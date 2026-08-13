# Review Checklist

Use this checklist for every executable plan review.

## Contract Coverage

- [ ] Governing specs and ADRs are named.
- [ ] Relevant skills are named.
- [ ] No feature is outside an approved spec.
- [ ] Required contract updates (OpenAPI or equivalent) are included for transport changes.
- [ ] Required retry, DLQ, audit, metrics, alarms, and replay are included for queue flows.

## Architecture Coverage

- [ ] `pkg/*` stays free of dependencies on `internal/*`, except the permitted public bridge.
- [ ] Public API changes are intentional and spec-backed, not incidental.
- [ ] Ports and use cases remain explicit.
- [ ] Wiring stays in the composition root (`cmd/*`).
- [ ] No new runtime or unsupported library is introduced without a decision record.

## Safety Coverage

- [ ] Tenant isolation by `{{DOMAIN_ACTOR}}` identity is preserved where relevant.
- [ ] `context.Context` and correlation identifiers propagate across boundaries.
- [ ] Secrets are not exposed in logs, artifacts, or review documents.
- [ ] Prompt-injection and tool safety risks are considered for agent/tool plans.
- [ ] Auth and security changes include explicit tests or diagnostics.

## Verification Coverage

- [ ] Validation commands are concrete.
- [ ] Expected outcomes are classifiable as `PASS`, `PARTIAL`, `FAIL`, or `BLOCKED`.
- [ ] Known gaps are listed.
- [ ] `APPROVED_WITH_CHANGES` changes are explicit and testable.
- [ ] Plan hash and review artifact are recorded.
