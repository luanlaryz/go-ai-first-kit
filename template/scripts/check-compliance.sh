#!/usr/bin/env bash

set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

log() {
	printf 'check-compliance: %s\n' "$*"
}

fail() {
	printf 'check-compliance: FAIL: %s\n' "$*" >&2
	exit 1
}

assert_file() {
	local file="$1"
	[[ -f "$file" ]] || fail "missing required file: $file"
}

assert_dir() {
	local dir="$1"
	[[ -d "$dir" ]] || fail "missing required directory: $dir"
}

assert_grep() {
	local pattern="$1"
	local file="$2"
	grep -Eq "$pattern" "$file" || fail "expected pattern [$pattern] in $file"
}

required_files=(
	"docs/ai/ai-contribution-contract.md"
	"docs/ai/task-input-format.md"
	"docs/ai/compliance-exceptions.md"
	".cursor/rules/{{PROJECT_SLUG}}.mdc"
	".github/PULL_REQUEST_TEMPLATE.md"
	".pre-commit-config.yaml"
	".github/workflows/ci.yml"
	".github/dependabot.yml"
	"SECURITY.md"
	"Makefile"
	"scripts/check-compliance.sh"
	"scripts/check-secrets.sh"
	"scripts/check-pr-body.sh"
	"scripts/create-pr-from-template.sh"
	"docs/ai/prompts/implement-task.md"
	"docs/ai/prompts/review-change.md"
	"docs/ai/prompts/diagnose-task.md"
	"docs/ai/prompts/remediation-bugfix.md"
	"docs/ai/briefs/ai-implementation-brief.md"
	"docs/ai/briefs/ai-diagnosis-brief.md"
	"docs/ai/briefs/ai-remediation-bugfix-brief.md"
	".cursor/hooks.json"
	".cursor/hooks/agent-cloud-access-guard.sh"
	".cursor/hooks/agent-cloud-access-guard.py"
	".cursor/hooks/cloud-access-config.json"
	".cursor/hooks/env-read-guard.sh"
	".cursor/hooks/env-read-guard.py"
	".cursor/hooks/governed-change-guard.sh"
	".cursor/hooks/parallel-collision-guard.sh"
	".cursor/rules/agent-cloud-access.mdc"
	".cursor/rules/governed-change-enforcement.mdc"
	".cursor/rules/dual-model-plan-review.mdc"
	".cursor/rules/parallel-session-safety.mdc"
	".cursor/rules/pre-pr-before-push.mdc"
	"scripts/check-agent-hooks.sh"
	"scripts/check-governed-change.sh"
	"scripts/check-sdd-compliance.sh"
	"scripts/check-parallel-collision.sh"
	"scripts/check-parallel-collision-selftest.sh"
	"scripts/check-architecture-boundaries.sh"
	"scripts/check-test-isolation.sh"
	"scripts/verify-plan-review.sh"
	"scripts/review-plan.py"
	"scripts/lib/plan_risk.py"
	"scripts/resolve-open-pr-body.sh"
	"scripts/resolve-pr-body-by-number.sh"
	"automation/PLAN_REVIEWS/README.md"
	".guardrails/README.md"
	".guardrails/function-size-exceptions.yaml"
	".guardrails/port-size-exceptions.yaml"
	".guardrails/ignored-error-exceptions.yaml"
	".guardrails/public-route-exceptions.yaml"
	".guardrails/package-exceptions.yaml"
	".guardrails/governed-change-exceptions.yaml"
	"tools/guardrails/main.go"
	"docs/runbooks/agent-cloud-access.md"
	"skills/27-dual-model-plan-review/SKILL.md"
	"skills/28-plan-review-autopilot/SKILL.md"
	"skills/29-governed-change-workflow/SKILL.md"
	"skills/30-third-party-service-integrations/SKILL.md"
	"skills/31-parallel-session-coordination/SKILL.md"
)

for file in "${required_files[@]}"; do
	assert_file "$file"
done

[[ -x "scripts/check-compliance.sh" ]] || fail "scripts/check-compliance.sh must be executable"
[[ -x "scripts/check-secrets.sh" ]] || fail "scripts/check-secrets.sh must be executable"
[[ -x "scripts/check-pr-body.sh" ]] || fail "scripts/check-pr-body.sh must be executable"
[[ -x "scripts/create-pr-from-template.sh" ]] || fail "scripts/create-pr-from-template.sh must be executable"
assert_dir "test/security"

# Governance and security gates are only controls if they can actually run.
for script in \
	scripts/check-agent-hooks.sh \
	scripts/check-governed-change.sh \
	scripts/check-sdd-compliance.sh \
	scripts/check-parallel-collision.sh \
	scripts/check-parallel-collision-selftest.sh \
	scripts/check-architecture-boundaries.sh \
	scripts/check-function-size.sh \
	scripts/check-port-size.sh \
	scripts/check-ignored-errors.sh \
	scripts/check-observability.sh \
	scripts/check-package-clean.sh \
	scripts/check-public-route-security.sh \
	scripts/check-test-isolation.sh \
	scripts/verify-plan-review.sh \
	scripts/resolve-open-pr-body.sh \
	scripts/resolve-pr-body-by-number.sh \
	scripts/env-keys.sh \
	.cursor/hooks/agent-cloud-access-guard.sh \
	.cursor/hooks/agent-cloud-access-guard-selftest.sh \
	.cursor/hooks/env-read-guard.sh \
	.cursor/hooks/env-read-guard-selftest.sh \
	.cursor/hooks/governed-change-guard.sh \
	.cursor/hooks/parallel-collision-guard.sh; do
	[[ -x "$script" ]] || fail "$script must be executable"
done

shopt -s nullglob

prompt_files=(docs/ai/prompts/*.md)
brief_files=(docs/ai/briefs/*.md)
(( ${#prompt_files[@]} > 0 )) || fail "docs/ai/prompts must contain at least one markdown file"
(( ${#brief_files[@]} > 0 )) || fail "docs/ai/briefs must contain at least one markdown file"

security_tests=($(find test/security -type f -name '*_test.go'))
(( ${#security_tests[@]} > 0 )) || fail "test/security must contain at least one *_test.go file"

for dir in $(find pkg -type f -name '*.go' ! -name '*_test.go' -exec dirname {} \; | sort -u); do
	assert_file "$dir/doc.go"
	tests=("$dir"/*_test.go)
	(( ${#tests[@]} > 0 )) || fail "missing *_test.go in $dir"
done

for target in lint vet check-compliance check-pr-body test-security; do
	assert_grep "^${target}:" "Makefile"
	assert_grep "run: make ${target}" ".github/workflows/ci.yml"
done

# Governance gates must exist as targets AND run in CI: a target nobody invokes
# is documentation, not a gate.
for target in check-agent-hooks hook-selftests guardrails verify-plan-reviews check-governed-change; do
	assert_grep "^${target}:" "Makefile"
	assert_grep "run: make ${target}" ".github/workflows/ci.yml"
done

for target in pre-pr check-sdd-compliance check-parallel-collision check-architecture-boundaries \
	check-function-size check-port-size check-ignored-errors check-observability \
	check-package-clean check-public-route-security check-test-isolation review-plan env-keys; do
	assert_grep "^${target}:" "Makefile"
done

assert_grep "^secrets-check:" "Makefile"
assert_grep "make secrets-check" ".pre-commit-config.yaml"
assert_grep "package-ecosystem: gomod" ".github/dependabot.yml"
assert_grep "package-ecosystem: github-actions" ".github/dependabot.yml"
assert_grep "make test-security" "SECURITY.md"

placeholder_patterns=(
	'\\bplaceholder\\b'
	'\\bTODO\\b'
	'\\bTBD\\b'
)

operational_files=(
	"Makefile"
	".github/workflows/ci.yml"
	".github/PULL_REQUEST_TEMPLATE.md"
	".pre-commit-config.yaml"
	".cursor/rules/{{PROJECT_SLUG}}.mdc"
	"docs/ai/ai-contribution-contract.md"
	"docs/ai/task-input-format.md"
	"docs/ai/compliance-exceptions.md"
	"${prompt_files[@]}"
	"${brief_files[@]}"
)

for file in "${operational_files[@]}"; do
	for pattern in "${placeholder_patterns[@]}"; do
		if grep -Eiq "$pattern" "$file"; then
			fail "forbidden placeholder marker found in $file"
		fi
	done
done

assert_grep "^## Objetivo$" ".github/PULL_REQUEST_TEMPLATE.md"
assert_grep "^## Backlog$" ".github/PULL_REQUEST_TEMPLATE.md"
assert_grep "^## Specs lidas$" ".github/PULL_REQUEST_TEMPLATE.md"
assert_grep "^## Autopilot$" ".github/PULL_REQUEST_TEMPLATE.md"
assert_grep "^## Regressao$" ".github/PULL_REQUEST_TEMPLATE.md"
assert_grep "^## Arquivos alterados$" ".github/PULL_REQUEST_TEMPLATE.md"
assert_grep "^## Impacto em docs e superficie publica$" ".github/PULL_REQUEST_TEMPLATE.md"
assert_grep "^## Comandos executados$" ".github/PULL_REQUEST_TEMPLATE.md"
assert_grep "^## Testes executados$" ".github/PULL_REQUEST_TEMPLATE.md"
assert_grep "^## Excecoes formais$" ".github/PULL_REQUEST_TEMPLATE.md"
assert_grep "^## Gaps restantes$" ".github/PULL_REQUEST_TEMPLATE.md"
assert_grep "^## Checklist$" ".github/PULL_REQUEST_TEMPLATE.md"
assert_grep "PULL_REQUEST_TEMPLATE.md" "AGENTS.md"
assert_grep "scripts/create-pr-from-template.sh" "AGENTS.md"
assert_grep "check-pr-body" ".github/workflows/ci.yml"
assert_grep "^# AI Contribution Contract$" "docs/ai/ai-contribution-contract.md"
assert_grep "^# Task Input Format$" "docs/ai/task-input-format.md"
assert_grep "spec governante" "docs/ai/ai-contribution-contract.md"
assert_grep "docs/ai/compliance-exceptions.md" "docs/ai/ai-contribution-contract.md"
assert_grep "arquivos alterados" "docs/ai/ai-contribution-contract.md"
assert_grep "testes executados" "docs/ai/ai-contribution-contract.md"
assert_grep "gaps restantes" "docs/ai/ai-contribution-contract.md"
assert_grep "^## Campos obrigatorios$" "docs/ai/task-input-format.md"
assert_grep '`Specs lidas`' "docs/ai/task-input-format.md"
assert_grep '`Impacto esperado em docs ou superficie publica`' "docs/ai/task-input-format.md"
assert_grep "^Specs lidas:$" "docs/ai/task-input-format.md"
assert_grep "^Impacto esperado em docs ou superficie publica:$" "docs/ai/task-input-format.md"
assert_grep "docs/ai/compliance-exceptions.md" "docs/ai/task-input-format.md"
assert_grep "^- arquivos alterados$" "docs/ai/task-input-format.md"
assert_grep "^- testes executados$" "docs/ai/task-input-format.md"
assert_grep "^- gaps restantes$" "docs/ai/task-input-format.md"
assert_grep "^# AI Compliance Exceptions$" "docs/ai/compliance-exceptions.md"
assert_grep "^## Estado atual$" "docs/ai/compliance-exceptions.md"
assert_grep "^## Formato de registro$" "docs/ai/compliance-exceptions.md"
assert_grep "^## Excecoes ativas$" "docs/ai/compliance-exceptions.md"

# Every guardrail allowlist must carry the exceptions section: an allowlist that
# is absent or malformed is a hard error in the gate, never "no exceptions".
for allowlist in .guardrails/*.yaml; do
	assert_grep "^exceptions:" "$allowlist"
done

# Security hooks must stay wired and fail-closed.
assert_grep "agent-cloud-access-guard.sh" ".cursor/hooks.json"
assert_grep "env-read-guard.sh" ".cursor/hooks.json"
assert_grep "governed-change-guard.sh" ".cursor/hooks.json"
assert_grep "parallel-collision-guard.sh" ".cursor/hooks.json"
assert_grep "failClosed" ".cursor/hooks.json"
assert_grep "skills/29-governed-change-workflow" "AGENTS.md"
assert_grep "skills/31-parallel-session-coordination" "AGENTS.md"

log "ok"
