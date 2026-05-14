## Agent Readiness

- Report: `<docs/reports/agent-readiness/...>`
- Source: `<agent-readiness command/report>`
- {{PROJECT_TITLE}}-filtered result: `<READY | READY_WITH_WARNINGS | NOT_READY>`
- PR impact: `<BLOCKS_PR | DOES_NOT_BLOCK_PR>`

### Relevant blockers

- [ ] None
- [ ] `<blocker id>` - `<required fix>`

### Optional recommendations

- `<recommendation>`

### Out-of-scope findings ignored

- `<finding>` - ignored because `<reason>`

### Notes

This readiness check filters generic `agent-readiness` output through the `{{PROJECT_SLUG}}` library/framework scope. Out-of-scope app/SaaS/deployment findings do not block this PR.
