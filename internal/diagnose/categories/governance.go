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
	qualityTotal := 5
	if inv.Contains(".github/PULL_REQUEST_TEMPLATE.md", "## Objetivo", "## Specs lidas", "## Checklist") {
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

	findings := make([]diagnose.Finding, 0)
	if missing := missingRequiredFiles(inv); len(missing) > 0 {
		findings = append(findings, finding(pillar, diagnose.SeverityCritical, "Arquivos obrigatorios ausentes", "Adicionar arquivos requeridos por scripts/check-compliance.sh.", diagnose.Snippet{Path: missing[0], Content: "Arquivo obrigatorio ausente."}))
	}
	if !hasDualSpecs(inv) {
		findings = append(findings, finding(pillar, diagnose.SeverityWarn, "Dual-spec ausente ou incompleta", "Garantir pares de spec de construcao e diagnostico em specs/."))
	}
	if !inv.Contains(".github/PULL_REQUEST_TEMPLATE.md", "## Gaps restantes") {
		findings = append(findings, finding(pillar, diagnose.SeverityWarn, "PR template incompleto", "Adicionar todas as secoes obrigatorias ao PULL_REQUEST_TEMPLATE.md.", inv.SnippetForPath(".github/PULL_REQUEST_TEMPLATE.md", "PR template ausente ou incompleto.")))
	}

	return diagnose.CategoryResult{
		Name:          pillar,
		Weight:        0.15,
		CoverageScore: diagnose.CoverageScore(present, total),
		QualityScore:  diagnose.CoverageScore(qualityChecks, qualityTotal),
		Summary:       "Compliance, dual-spec, PR template, CI e dependabot",
		Findings:      findings,
	}
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
