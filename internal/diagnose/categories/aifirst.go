package categories

import (
	"strings"

	"github.com/inventa-co/go-ai-first-kit/internal/diagnose"
)

type aiFirstChecker struct{}

func (aiFirstChecker) Check(inv diagnose.Inventory) diagnose.CategoryResult {
	const pillar = "AI-first"
	requiredFiles := []string{
		"AGENTS.md",
		"skills/00-skill-index/SKILL.md",
		"skills/26-backlog-item-intake/SKILL.md",
		"automation/AUTOPILOT.md",
		"automation/RUNBOOK.md",
		"automation/INTERACTIVE_AUTOPILOT.md",
		"automation/INTERACTIVE_RUNBOOK.md",
		"docs/ai/prompts/diagnose-task.md",
		"docs/ai/briefs/ai-diagnosis-brief.md",
		"docs/backlog/Backlog.md",
	}
	requiredDirs := []string{".cursor/rules", ".cursor/hooks", "skills", "docs/ai/prompts", "docs/ai/briefs"}
	present := countFiles(inv, requiredFiles) + countDirs(inv, requiredDirs)
	total := len(requiredFiles) + len(requiredDirs)

	skillCount := 0
	for _, file := range inv.FilesWithPrefix("skills") {
		if strings.HasSuffix(file, "SKILL.md") {
			skillCount++
		}
	}
	qualityChecks := 0
	qualityTotal := 6
	if inv.Contains("AGENTS.md", "Spec Driven Development") || inv.Contains("AGENTS.md", "spec governante") {
		qualityChecks++
	}
	if inv.Contains("AGENTS.md", "dual-spec") || inv.Contains("AGENTS.md", "spec de diagnostico") {
		qualityChecks++
	}
	if skillCount >= 20 {
		qualityChecks++
	}
	if inv.CountMarkdownFiles("docs/ai/prompts") > 0 && inv.CountMarkdownFiles("docs/ai/briefs") > 0 {
		qualityChecks++
	}
	if inv.FileExists("automation/INTERACTIVE_STATE.json") && inv.FileExists("automation/PHASE_STATE.json") {
		qualityChecks++
	}
	if inv.FileExists("docs/backlog/Backlog.md") && inv.FileExists("skills/26-backlog-item-intake/SKILL.md") {
		qualityChecks++
	}

	findings := make([]diagnose.Finding, 0)
	if !inv.FileExists("AGENTS.md") {
		findings = append(findings, finding(pillar, diagnose.SeverityCritical, "AGENTS.md ausente", "Adicionar constituicao operacional para agentes com source of truth e workflow obrigatorio."))
	}
	if skillCount < 20 {
		findings = append(findings, finding(pillar, diagnose.SeverityWarn, "Indice de skills incompleto", "Garantir skills suficientes e indexadas para roteamento de agentes."))
	}
	if !inv.FileExists("docs/ai/prompts/diagnose-task.md") {
		findings = append(findings, finding(pillar, diagnose.SeverityWarn, "Prompt de diagnostico ausente", "Adicionar prompt estatico de diagnostico em docs/ai/prompts/diagnose-task.md."))
	}
	if !inv.FileExists("docs/backlog/Backlog.md") {
		findings = append(findings, finding(pillar, diagnose.SeverityWarn, "Backlog canonico ausente", "Adicionar docs/backlog/Backlog.md e a skill 26-backlog-item-intake para intake governado de backlog."))
	}

	return diagnose.CategoryResult{
		Name:          pillar,
		Weight:        0.20,
		CoverageScore: diagnose.CoverageScore(present, total),
		QualityScore:  diagnose.CoverageScore(qualityChecks, qualityTotal),
		Summary:       "Instrucoes para agentes, skills, prompts, briefs, autopilots e backlog",
		Findings:      findings,
	}
}
