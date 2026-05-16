package ui

import (
	"strings"
	"testing"

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
