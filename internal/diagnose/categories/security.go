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
		".cursor/hooks/agent-cloud-access-guard.sh",
		".cursor/hooks/env-read-guard.sh",
		"scripts/check-agent-hooks.sh",
		"docs/runbooks/agent-cloud-access.md",
	}
	present := countFiles(inv, required)
	if inv.DirExists("test/security") {
		present++
	}
	total := len(required) + 1

	qualityChecks := 0
	qualityTotal := 6
	if inv.Contains("Makefile", "secrets-check:", "vulncheck:", "test-security:") {
		qualityChecks++
	}
	// A guard that can be silently unwired is not a control, so the wiring gate
	// counts as much as the guard itself.
	if inv.Contains(".cursor/hooks.json", "agent-cloud-access-guard.sh", "env-read-guard.sh", "failClosed") {
		qualityChecks++
	}
	if inv.Contains("Makefile", "check-agent-hooks:") {
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
	if !inv.Contains(".cursor/hooks.json", "failClosed") {
		findings = append(findings, finding(pillar, diagnose.SeverityWarn, "Hooks de credencial sem fail-closed", "Registrar os guards de acesso a nuvem e leitura de .env em .cursor/hooks.json com failClosed: true e validar com make check-agent-hooks.", inv.SnippetForPath(".cursor/hooks.json", "hooks.json ausente ou sem guard fail-closed.")))
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
