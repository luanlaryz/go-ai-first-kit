package ui

import (
	"strings"
	"testing"
	"time"

	"github.com/inventa-co/go-ai-first-kit/internal/diagnose"
)

func TestRenderTerminalWrapsLongCategorySummary(t *testing.T) {
	summary := "Tooling local, formatacao, hooks e scripts de compliance para times grandes sem perder detalhes"
	report := diagnose.Report{
		ProjectSlug: "example",
		Path:        "/tmp/example",
		GlobalScore: 92,
		Band:        diagnose.BandOK,
		Categories: []diagnose.CategoryResult{
			{
				Name:    "DX",
				Weight:  0.10,
				Score:   92,
				Summary: summary,
			},
			{
				Name:    "AI-first",
				Weight:  0.20,
				Score:   100,
				Summary: "Instrucoes para agentes, skills, prompts, briefs e autopilots",
			},
		},
	}

	got := RenderTerminal(report, true)

	if strings.Contains(got, "…") {
		t.Fatalf("terminal output should wrap long summaries instead of truncating them:\n%s", got)
	}

	wantFirstLine := "Tooling local, formatacao, hooks e scripts de compliance para\n"
	if !strings.Contains(got, wantFirstLine) {
		t.Fatalf("terminal output missing wrapped first summary line %q:\n%s", wantFirstLine, got)
	}

	wantContinuation := strings.Repeat(" ", terminalJustificationIndent) + "times grandes sem perder detalhes\n"
	if !strings.Contains(got, wantContinuation) {
		t.Fatalf("terminal output missing aligned continuation line %q:\n%s", wantContinuation, got)
	}

	separator := strings.Repeat("-", terminalJustificationIndent+terminalJustificationWidth)
	wantSeparators := len(report.Categories) + 1
	if gotSeparators := strings.Count(got, separator+"\n"); gotSeparators != wantSeparators {
		t.Fatalf("terminal output should render %d table separators, got %d:\n%s", wantSeparators, gotSeparators, got)
	}
}

func TestRenderCorrectionPlanPromptIncludesDualSpecAndFindings(t *testing.T) {
	report := diagnose.Report{
		ProjectSlug: "example",
		Path:        "/tmp/example",
		ScannedAt:   time.Date(2026, 5, 16, 10, 0, 0, 0, time.UTC),
		FileCount:   12,
		GlobalScore: 30,
		Band:        diagnose.BandCritical,
		Improvements: []diagnose.Finding{
			{
				Pillar:         "Security",
				Severity:       diagnose.SeverityCritical,
				Title:          "Politica de seguranca ausente",
				Recommendation: "Adicionar politica de seguranca e comando de validacao esperado.",
				Evidence: []diagnose.Snippet{
					{
						Path:      "SECURITY.md",
						StartLine: 1,
						EndLine:   1,
						Content:   "Arquivo obrigatorio ausente.",
					},
				},
			},
		},
	}

	got := RenderCorrectionPlanPrompt(report)

	for _, want := range []string{
		"Crie um plano de correcao priorizado",
		"Produza somente plano; nao implemente",
		"Spec de implementacao",
		"Spec de diagnostico",
		"[CRIT][Security] Politica de seguranca ausente",
		"Adicionar politica de seguranca e comando de validacao esperado.",
		"`SECURITY.md:1-1`",
		"Arquivo obrigatorio ausente.",
		"Proxima acao recomendada sem executar a correcao.",
	} {
		if !strings.Contains(got, want) {
			t.Fatalf("correction plan prompt missing %q:\n%s", want, got)
		}
	}
}

func TestRenderCorrectionPlanPromptHandlesNoFindings(t *testing.T) {
	report := diagnose.Report{
		ProjectSlug: "example",
		Path:        "/tmp/example",
		ScannedAt:   time.Date(2026, 5, 16, 10, 0, 0, 0, time.UTC),
		FileCount:   12,
		GlobalScore: 92,
		Band:        diagnose.BandOK,
	}

	got := RenderCorrectionPlanPrompt(report)

	for _, want := range []string{
		"Nenhum item critico ou warning foi detectado.",
		"plano leve de confirmacao",
		"manutencao dos gates existentes",
	} {
		if !strings.Contains(got, want) {
			t.Fatalf("correction plan prompt without findings missing %q:\n%s", want, got)
		}
	}
}
