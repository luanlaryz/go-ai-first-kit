package categories

import "github.com/inventa-co/go-ai-first-kit/internal/diagnose"

type openAPIChecker struct{}

func (openAPIChecker) Check(inv diagnose.Inventory) diagnose.CategoryResult {
	const pillar = "OpenAPI"
	candidates := []string{"docs/api/openapi.yaml", "docs/api/openapi.yml", "docs/api/openapi.json", "api/openapi.yaml", "api/openapi.yml", "api/openapi.json"}
	openAPIFile := ""
	for _, candidate := range candidates {
		if inv.FileExists(candidate) {
			openAPIFile = candidate
			break
		}
	}

	present := 0
	total := 3
	if openAPIFile != "" {
		present++
	}
	if inv.FileExists("skills/10-http-gin-openapi/SKILL.md") {
		present++
	}
	if inv.Contains("Makefile", "openapi") || inv.Contains(".github/workflows/ci.yml", "openapi") {
		present++
	}

	qualityChecks := 0
	qualityTotal := 3
	if openAPIFile != "" && (inv.Contains(openAPIFile, "openapi:", "paths:") || inv.Contains(openAPIFile, `"openapi"`, `"paths"`)) {
		qualityChecks++
	}
	if inv.Contains("skills/10-http-gin-openapi/SKILL.md", "OpenAPI") {
		qualityChecks++
	}
	if inv.Contains(".github/workflows/ci.yml", "openapi") || inv.Contains("Makefile", "openapi-lint") {
		qualityChecks++
	}

	findings := make([]diagnose.Finding, 0)
	if openAPIFile == "" {
		findings = append(findings, finding(pillar, diagnose.SeverityCritical, "Contrato OpenAPI ausente", "Adicionar docs/api/openapi.yaml ou documentar formalmente por que o projeto nao expoe HTTP."))
	}
	if openAPIFile != "" && qualityChecks == 0 {
		findings = append(findings, finding(pillar, diagnose.SeverityWarn, "OpenAPI sem estrutura minima", "Garantir campos openapi e paths no contrato.", inv.SnippetForPath(openAPIFile, "Contrato encontrado sem estrutura minima.")))
	}

	return diagnose.CategoryResult{
		Name:          pillar,
		Weight:        0.10,
		CoverageScore: diagnose.CoverageScore(present, total),
		QualityScore:  diagnose.CoverageScore(qualityChecks, qualityTotal),
		Summary:       "Contrato HTTP/OpenAPI e validacao de superficie publica",
		Findings:      findings,
	}
}
