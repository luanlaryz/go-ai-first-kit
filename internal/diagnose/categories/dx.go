package categories

import "github.com/inventa-co/go-ai-first-kit/internal/diagnose"

type dxChecker struct{}

func (dxChecker) Check(inv diagnose.Inventory) diagnose.CategoryResult {
	const pillar = "DX"
	required := []string{
		"Makefile",
		".pre-commit-config.yaml",
		".editorconfig",
		".gitignore",
		"scripts/check-compliance.sh",
		"scripts/check-file-size.sh",
	}
	present := countFiles(inv, required)
	qualityChecks := 0
	qualityTotal := 5
	if inv.Contains("Makefile", "setup:", "fmt:", "vet:", "lint:", "test:", "coverage:") {
		qualityChecks++
	}
	if inv.Contains(".pre-commit-config.yaml", "scripts/check-secrets.sh") || inv.Contains(".pre-commit-config.yaml", "make secrets-check") {
		qualityChecks++
	}
	if inv.Contains(".editorconfig", "end_of_line", "insert_final_newline") {
		qualityChecks++
	}
	if inv.Contains(".gitignore", ".env", "coverage") {
		qualityChecks++
	}
	if inv.FileExists("scripts/check-compliance.sh") && inv.FileExists("scripts/check-pr-body.sh") {
		qualityChecks++
	}

	findings := make([]diagnose.Finding, 0)
	if !inv.FileExists("Makefile") {
		findings = append(findings, finding(pillar, diagnose.SeverityCritical, "Makefile ausente", "Adicionar Makefile com alvos setup, fmt, vet, lint, test e coverage."))
	}
	if !inv.FileExists(".pre-commit-config.yaml") {
		findings = append(findings, finding(pillar, diagnose.SeverityWarn, "Pre-commit ausente", "Adicionar .pre-commit-config.yaml com checagem de secrets e higiene basica."))
	}

	coverage := diagnose.CoverageScore(present, len(required))
	quality := diagnose.CoverageScore(qualityChecks, qualityTotal)
	return diagnose.CategoryResult{
		Name:          pillar,
		Weight:        0.10,
		CoverageScore: coverage,
		QualityScore:  quality,
		Summary:       "Tooling local, formatacao, hooks e scripts de compliance",
		Findings:      findings,
	}
}
