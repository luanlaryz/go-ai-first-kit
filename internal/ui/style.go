package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/inventa-co/go-ai-first-kit/internal/diagnose"
)

type Styles struct {
	noColor bool
	title   lipgloss.Style
	ok      lipgloss.Style
	warn    lipgloss.Style
	crit    lipgloss.Style
	muted   lipgloss.Style
}

func NewStyles(noColor bool) Styles {
	return Styles{
		noColor: noColor,
		title:   lipgloss.NewStyle().Bold(true),
		ok:      lipgloss.NewStyle().Foreground(lipgloss.Color("2")).Bold(true),
		warn:    lipgloss.NewStyle().Foreground(lipgloss.Color("3")).Bold(true),
		crit:    lipgloss.NewStyle().Foreground(lipgloss.Color("1")).Bold(true),
		muted:   lipgloss.NewStyle().Foreground(lipgloss.Color("8")),
	}
}

func (s Styles) Band(band diagnose.Band) string {
	text := string(band)
	if s.noColor {
		return text
	}
	switch band {
	case diagnose.BandOK:
		return s.ok.Render(text)
	case diagnose.BandWarn:
		return s.warn.Render(text)
	default:
		return s.crit.Render(text)
	}
}

func (s Styles) Severity(severity diagnose.Severity) string {
	text := string(severity)
	if s.noColor {
		return text
	}
	switch severity {
	case diagnose.SeverityCritical:
		return s.crit.Render(text)
	case diagnose.SeverityWarn:
		return s.warn.Render(text)
	default:
		return s.muted.Render(text)
	}
}

func (s Styles) Score(score int) string {
	text := fmt.Sprintf("%3d", score)
	if s.noColor {
		return text
	}
	switch diagnose.BandForScore(score) {
	case diagnose.BandOK:
		return s.ok.Render(text)
	case diagnose.BandWarn:
		return s.warn.Render(text)
	default:
		return s.crit.Render(text)
	}
}
