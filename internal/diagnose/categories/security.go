package categories

import "github.com/inventa-co/go-ai-first-kit/internal/diagnose"

type securityChecker struct{}

func (securityChecker) Check(inv diagnose.Inventory) diagnose.CategoryResult {
	const pillar = "Security"
	required := []string{
		"SECURITY.md",
		"test/security/security_baseline_test.go",
		"scripts/check-secrets.sh",
		"skills/18-security-owasp-api/SKILL.md",
		"skills/19-prompt-injection-llm-safety/SKILL.md",
	}
	present := countFiles(inv, required)
	if inv.DirExists("test/security") {
		present++
	}
	total := len(required) + 1

	qualityChecks := 0
	qualityTotal := 4
	if inv.Contains("Makefile", "secrets-check:", "vulncheck:", "test-security:") {
		qualityChecks++
	}
	if inv.Contains(".pre-commit-config.yaml", "make secrets-check") || inv.Contains(".pre-commit-config.yaml", "check-secrets.sh") {
		qualityChecks++
	}
	if inv.Contains("SECURITY.md", "make test-security") {
		qualityChecks++
	}
	if len(inv.FilesWithPrefix("test/security")) > 0 {
		qualityChecks++
	}

	findings := make([]diagnose.Finding, 0)
	if !inv.FileExists("SECURITY.md") {
		findings = append(findings, finding(pillar, diagnose.SeverityCritical, "SECURITY.md ausente", "Adicionar politica de seguranca e comando de validacao esperado."))
	}
	if !inv.Contains("Makefile", "vulncheck:") {
		findings = append(findings, finding(pillar, diagnose.SeverityWarn, "govulncheck nao exposto", "Adicionar alvo vulncheck ao Makefile e CI quando aplicavel.", inv.SnippetForPath("Makefile", "Makefile ausente ou sem vulncheck.")))
	}

	return diagnose.CategoryResult{
		Name:          pillar,
		Weight:        0.15,
		CoverageScore: diagnose.CoverageScore(present, total),
		QualityScore:  diagnose.CoverageScore(qualityChecks, qualityTotal),
		Summary:       "Politica, secrets, vulncheck, testes e prompt-injection safety",
		Findings:      findings,
	}
}
