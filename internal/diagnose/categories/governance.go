package categories

import (
	"strings"

	"github.com/inventa-co/go-ai-first-kit/internal/diagnose"
)

var RequiredComplianceFiles = []string{
	"docs/ai/ai-contribution-contract.md",
	"docs/ai/task-input-format.md",
	"docs/ai/compliance-exceptions.md",
	".cursor/rules/{{PROJECT_SLUG}}.mdc",
	".github/PULL_REQUEST_TEMPLATE.md",
	".pre-commit-config.yaml",
	".github/workflows/ci.yml",
	".github/dependabot.yml",
	"SECURITY.md",
	"Makefile",
	"scripts/check-compliance.sh",
	"scripts/check-secrets.sh",
	"scripts/check-pr-body.sh",
	"scripts/create-pr-from-template.sh",
	"docs/ai/prompts/implement-task.md",
	"docs/ai/prompts/review-change.md",
	"docs/ai/prompts/diagnose-task.md",
	"docs/ai/prompts/remediation-bugfix.md",
	"docs/ai/briefs/ai-implementation-brief.md",
	"docs/ai/briefs/ai-diagnosis-brief.md",
	"docs/ai/briefs/ai-remediation-bugfix-brief.md",
	".cursor/hooks.json",
	".cursor/hooks/agent-cloud-access-guard.sh",
	".cursor/hooks/agent-cloud-access-guard.py",
	".cursor/hooks/cloud-access-config.json",
	".cursor/hooks/env-read-guard.sh",
	".cursor/hooks/env-read-guard.py",
	".cursor/hooks/governed-change-guard.sh",
	".cursor/hooks/parallel-collision-guard.sh",
	".cursor/rules/agent-cloud-access.mdc",
	".cursor/rules/governed-change-enforcement.mdc",
	".cursor/rules/dual-model-plan-review.mdc",
	".cursor/rules/parallel-session-safety.mdc",
	".cursor/rules/pre-pr-before-push.mdc",
	"scripts/check-agent-hooks.sh",
	"scripts/check-governed-change.sh",
	"scripts/check-sdd-compliance.sh",
	"scripts/check-parallel-collision.sh",
	"scripts/check-parallel-collision-selftest.sh",
	"scripts/check-architecture-boundaries.sh",
	"scripts/check-test-isolation.sh",
	"scripts/verify-plan-review.sh",
	"scripts/review-plan.py",
	"scripts/lib/plan_risk.py",
	"scripts/resolve-open-pr-body.sh",
	"scripts/resolve-pr-body-by-number.sh",
	"automation/PLAN_REVIEWS/README.md",
	".guardrails/README.md",
	".guardrails/function-size-exceptions.yaml",
	".guardrails/port-size-exceptions.yaml",
	".guardrails/ignored-error-exceptions.yaml",
	".guardrails/public-route-exceptions.yaml",
	".guardrails/package-exceptions.yaml",
	".guardrails/governed-change-exceptions.yaml",
	"tools/guardrails/main.go",
	"docs/runbooks/agent-cloud-access.md",
	"skills/27-dual-model-plan-review/SKILL.md",
	"skills/28-plan-review-autopilot/SKILL.md",
	"skills/29-governed-change-workflow/SKILL.md",
	"skills/30-third-party-service-integrations/SKILL.md",
	"skills/31-parallel-session-coordination/SKILL.md",
}

type governanceChecker struct{}

func (governanceChecker) Check(inv diagnose.Inventory) diagnose.CategoryResult {
	const pillar = "Governanca"
	present := countRequiredFiles(inv)
	total := len(RequiredComplianceFiles)
	if inv.DirExists("specs") {
		present++
		total++
	}

	qualityChecks := 0
	qualityTotal := 11
	if inv.Contains(".github/PULL_REQUEST_TEMPLATE.md", "## Objetivo", "## Specs lidas", "## Checklist") {
		qualityChecks++
	}
	if inv.Contains(".github/PULL_REQUEST_TEMPLATE.md", "## Backlog", "## Autopilot", "## Regressao") {
		qualityChecks++
	}
	if hasGovernedChangeGate(inv) {
		qualityChecks++
	}
	if hasPlanReviewGate(inv) {
		qualityChecks++
	}
	if hasExpiringAllowlists(inv) {
		qualityChecks++
	}
	if inv.Contains(".github/workflows/ci.yml", "make check-compliance", "make check-pr-body") {
		qualityChecks++
	}
	if inv.Contains(".github/dependabot.yml", "package-ecosystem: gomod", "package-ecosystem: github-actions") {
		qualityChecks++
	}
	if hasDualSpecs(inv) {
		qualityChecks++
	}
	if inv.Contains("docs/ai/compliance-exceptions.md", "## Excecoes ativas") || inv.Contains("docs/ai/compliance-exceptions.md", "## Exceções ativas") {
		qualityChecks++
	}
	if inv.FileExists("docs/decisions/README.md") {
		qualityChecks++
	}
	if inv.FileExists("docs/release-versioning-policy.md") {
		qualityChecks++
	}

	findings := make([]diagnose.Finding, 0)
	if missing := missingRequiredFiles(inv); len(missing) > 0 {
		findings = append(findings, finding(pillar, diagnose.SeverityCritical, "Arquivos obrigatorios ausentes", "Adicionar arquivos requeridos por scripts/check-compliance.sh.", diagnose.Snippet{Path: missing[0], Content: "Arquivo obrigatorio ausente."}))
	}
	if !hasDualSpecs(inv) {
		findings = append(findings, finding(pillar, diagnose.SeverityWarn, "Dual-spec ausente ou incompleta", "Garantir pares de spec de construcao e diagnostico em specs/."))
	}
	if !inv.DirExists("docs/decisions") {
		findings = append(findings, finding(pillar, diagnose.SeverityWarn, "Sistema de ADR ausente", "Adicionar docs/decisions/ com README e ADR seed para registrar decisoes de arquitetura e governanca."))
	}
	if !inv.FileExists("docs/release-versioning-policy.md") {
		findings = append(findings, finding(pillar, diagnose.SeverityWarn, "Politica de release ausente", "Adicionar docs/release-versioning-policy.md alinhada a skill 22-release-versioning-governance."))
	}
	if !inv.Contains(".github/PULL_REQUEST_TEMPLATE.md", "## Gaps restantes") {
		findings = append(findings, finding(pillar, diagnose.SeverityWarn, "PR template incompleto", "Adicionar todas as secoes obrigatorias ao PULL_REQUEST_TEMPLATE.md.", inv.SnippetForPath(".github/PULL_REQUEST_TEMPLATE.md", "PR template ausente ou incompleto.")))
	}
	if !hasGovernedChangeGate(inv) {
		findings = append(findings, finding(pillar, diagnose.SeverityWarn, "Gate de mudanca governada ausente", "Adicionar scripts/check-governed-change.sh e o alvo make correspondente para exigir backlog, dual-spec e plano revisado em escopo governado."))
	}
	if !hasPlanReviewGate(inv) {
		findings = append(findings, finding(pillar, diagnose.SeverityWarn, "Gate de review de plano ausente", "Adicionar scripts/verify-plan-review.sh e automation/PLAN_REVIEWS/ para que plano so execute com veredito registrado."))
	}
	if !hasExpiringAllowlists(inv) {
		findings = append(findings, finding(pillar, diagnose.SeverityWarn, "Allowlist de guardrail sem prazo", "Adicionar .guardrails/*.yaml com expires_at obrigatorio para que divida tecnica tenha vencimento."))
	}

	return diagnose.CategoryResult{
		Name:          pillar,
		Weight:        0.15,
		CoverageScore: diagnose.CoverageScore(present, total),
		QualityScore:  diagnose.CoverageScore(qualityChecks, qualityTotal),
		Summary:       "Compliance, dual-spec, PR template, CI, dependabot, ADR e release",
		Findings:      findings,
	}
}

// hasGovernedChangeGate reports whether the governed-change trio is enforced by
// an executable gate wired into the Makefile, not merely documented.
func hasGovernedChangeGate(inv diagnose.Inventory) bool {
	return inv.FileExists("scripts/check-governed-change.sh") &&
		inv.Contains("Makefile", "check-governed-change")
}

// hasPlanReviewGate reports whether plan execution is gated by a recorded
// verdict: the verifier, the artifact directory and a Makefile entry point.
func hasPlanReviewGate(inv diagnose.Inventory) bool {
	return inv.FileExists("scripts/verify-plan-review.sh") &&
		inv.FileExists("automation/PLAN_REVIEWS/README.md") &&
		inv.Contains("Makefile", "verify-plan-reviews")
}

// hasExpiringAllowlists reports whether guardrail exceptions carry an expiry
// contract. An allowlist without expiry turns technical debt into permission, so
// any file that declares entries must also declare expires_at for them.
func hasExpiringAllowlists(inv diagnose.Inventory) bool {
	found := false
	for _, file := range inv.FilesWithPrefix(".guardrails") {
		if !strings.HasSuffix(file, ".yaml") {
			continue
		}
		found = true
		if !inv.Contains(file, "exceptions:") {
			return false
		}
		// An empty allowlist is compliant. One with entries is only compliant
		// when every entry can expire.
		if inv.Contains(file, "- path:") && !inv.Contains(file, "expires_at:") {
			return false
		}
	}
	return found && inv.FileExists(".guardrails/README.md")
}

func countRequiredFiles(inv diagnose.Inventory) int {
	count := 0
	for _, file := range RequiredComplianceFiles {
		if inv.FileExists(file) || renderedRuleExists(inv, file) {
			count++
		}
	}
	return count
}

func missingRequiredFiles(inv diagnose.Inventory) []string {
	var missing []string
	for _, file := range RequiredComplianceFiles {
		if inv.FileExists(file) || renderedRuleExists(inv, file) {
			continue
		}
		missing = append(missing, file)
	}
	return missing
}

func renderedRuleExists(inv diagnose.Inventory, file string) bool {
	if file != ".cursor/rules/{{PROJECT_SLUG}}.mdc" {
		return false
	}
	for _, candidate := range inv.FilesWithPrefix(".cursor/rules") {
		if candidate != ".cursor/rules/{{PROJECT_SLUG}}.mdc" && strings.HasSuffix(candidate, ".mdc") {
			return true
		}
	}
	return false
}

func hasDualSpecs(inv diagnose.Inventory) bool {
	buildSpecs := 0
	diagnosisSpecs := 0
	for _, file := range inv.FilesWithPrefix("specs") {
		if !strings.HasSuffix(file, ".md") {
			continue
		}
		if containsDiagnosisName(file) {
			diagnosisSpecs++
		} else {
			buildSpecs++
		}
	}
	return buildSpecs > 0 && diagnosisSpecs > 0
}

func containsDiagnosisName(file string) bool {
	file = strings.ToLower(file)
	for _, marker := range []string{"diagnosis", "diagnostico", "diagnóstico"} {
		if strings.Contains(file, marker) {
			return true
		}
	}
	return false
}
