# Agent Readiness Report - <scope>

Date: <YYYY-MM-DD>
Scope: <feature/refactor/PR/branch>
Source: <agent-readiness command/report path>

## 1. Executive summary

- Raw readiness score: `<score or unknown>`
- {{PROJECT_TITLE}}-filtered result: `<READY | READY_WITH_WARNINGS | NOT_READY>`
- PR impact: `<BLOCKS_PR | DOES_NOT_BLOCK_PR>`
- Relevant blockers: `<count>`
- Optional recommendations: `<count>`
- Out-of-scope findings ignored: `<count>`

## 2. Scope interpreted for {{PROJECT_SLUG}}

Describe why this assessment matters for a Go library/framework and which areas were considered in scope.

## 3. Source evidence

List:

- report path or command used
- timestamp
- format (`text`, `json`, or `manual summary`)
- relevant commit/branch/PR if known

## 4. Classified findings

| Finding | Original category | {{PROJECT_TITLE}} classification | PR impact | Rationale |
| --- | --- | --- | --- | --- |
| `<id/name>` | `<category>` | `worth_it_for_{{PROJECT_SLUG}}` | `BLOCKS_PR` | `<why>` |
| `<id/name>` | `<category>` | `optional_for_{{PROJECT_SLUG}}` | `DOES_NOT_BLOCK_PR` | `<why>` |
| `<id/name>` | `<category>` | `out_of_scope_for_{{PROJECT_SLUG}}` | `OUT_OF_SCOPE` | `<why>` |

## 5. PR blockers

List only findings that block the current PR.

For each blocker:

- finding
- affected files
- required fix
- acceptance evidence

## 6. Non-blocking recommendations

List findings worth considering later, grouped by theme.

## 7. Out-of-scope findings ignored

List findings intentionally ignored for this PR with rationale.

## 8. Documentation updates

State whether updates are required for:

- `AGENTS.md`
- `README.md`
- `docs/*`
- `automation/*`
- `.cursor/rules/*`
- `skills/*`
- PR description

## 9. Autopilot impact

State whether this report:

- creates no new autopilot work
- adds tasks to the current trail
- requires a separate SDD trail
- blocks advancement

## 10. Final decision

Use exactly one:

- `READY`
- `READY_WITH_WARNINGS`
- `NOT_READY`

Then explain in one paragraph.
