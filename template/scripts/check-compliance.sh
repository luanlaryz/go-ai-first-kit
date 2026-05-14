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
)

for file in "${required_files[@]}"; do
	assert_file "$file"
done

[[ -x "scripts/check-compliance.sh" ]] || fail "scripts/check-compliance.sh must be executable"
[[ -x "scripts/check-secrets.sh" ]] || fail "scripts/check-secrets.sh must be executable"
[[ -x "scripts/check-pr-body.sh" ]] || fail "scripts/check-pr-body.sh must be executable"
[[ -x "scripts/create-pr-from-template.sh" ]] || fail "scripts/create-pr-from-template.sh must be executable"
assert_dir "test/security"

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
assert_grep "^## Specs lidas$" ".github/PULL_REQUEST_TEMPLATE.md"
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

log "ok"
