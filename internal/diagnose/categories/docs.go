package categories

import "github.com/inventa-co/go-ai-first-kit/internal/diagnose"

type docsChecker struct{}

func (docsChecker) Check(inv diagnose.Inventory) diagnose.CategoryResult {
	const pillar = "Documentacao"
	required := []string{"README.md", "CONTRIBUTING.md", "CHANGELOG.md", "doc.go", "skills/21-documentation-open-source/SKILL.md"}
	present := countFiles(inv, required)
	if inv.DirExists("examples") {
		present++
	}
	total := len(required) + 1

	qualityChecks := 0
	qualityTotal := 5
	if inv.Contains("README.md", "##") && (inv.Contains("README.md", "Instal") || inv.Contains("README.md", "Uso")) {
		qualityChecks++
	}
	if inv.Contains("CONTRIBUTING.md", "PR") || inv.Contains("CONTRIBUTING.md", "contrib") {
		qualityChecks++
	}
	if inv.Contains("CHANGELOG.md", "## [Unreleased]") {
		qualityChecks++
	}
	if len(inv.FilesWithPrefix("examples")) > 0 {
		qualityChecks++
	}
	if inv.FileExists("doc.go") {
		qualityChecks++
	}

	findings := make([]diagnose.Finding, 0)
	if !inv.FileExists("README.md") {
		findings = append(findings, finding(pillar, diagnose.SeverityCritical, "README.md ausente", "Adicionar README com descricao, instalacao e uso."))
	}
	if !inv.FileExists("CHANGELOG.md") {
		findings = append(findings, finding(pillar, diagnose.SeverityWarn, "CHANGELOG.md ausente", "Adicionar changelog com secao [Unreleased]."))
	}
	if !inv.DirExists("examples") {
		findings = append(findings, finding(pillar, diagnose.SeverityWarn, "Examples ausentes", "Adicionar examples/ com pelo menos um caminho executavel ou documentado."))
	}

	return diagnose.CategoryResult{
		Name:          pillar,
		Weight:        0.15,
		CoverageScore: diagnose.CoverageScore(present, total),
		QualityScore:  diagnose.CoverageScore(qualityChecks, qualityTotal),
		Summary:       "README, contribuicao, changelog, doc.go e examples",
		Findings:      findings,
	}
}
