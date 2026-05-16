package ui

import (
	"fmt"
	"strings"

	"github.com/inventa-co/go-ai-first-kit/internal/diagnose"
	"github.com/inventa-co/go-ai-first-kit/internal/version"
)

func RenderTerminal(report diagnose.Report, noColor bool) string {
	styles := NewStyles(noColor)
	var b strings.Builder
	fmt.Fprintf(&b, "gakit diagnose · %s\n", report.ProjectSlug)
	fmt.Fprintf(&b, "path: %s\n", report.Path)
	fmt.Fprintf(&b, "scanned: %d files · elapsed: %s\n\n", report.FileCount, report.Elapsed.Round(1_000_000))
	fmt.Fprintf(&b, "%-24s %7s %-10s %6s  %s\n", "Pilar", "Score", "Faixa", "Peso", "Justificativa")
	fmt.Fprintf(&b, "%s\n", strings.Repeat("-", 86))
	for _, category := range report.Categories {
		fmt.Fprintf(&b, "%-24s %7s %-10s %5.0f%%  %s\n",
			category.Name,
			styles.Score(category.Score),
			styles.Band(diagnose.BandForScore(category.Score)),
			category.Weight*100,
			truncate(category.Summary, 44),
		)
	}
	fmt.Fprintf(&b, "%s\n", strings.Repeat("-", 86))
	fmt.Fprintf(&b, "%-24s %7s %-10s %5s  %s\n\n",
		"Score Global Ponderado",
		styles.Score(report.GlobalScore),
		styles.Band(report.Band),
		"100%",
		"",
	)
	if len(report.Improvements) == 0 {
		b.WriteString("Itens a melhorar\nNenhum item critico detectado.\n")
		return b.String()
	}
	b.WriteString("Itens a melhorar\n")
	b.WriteString(strings.Repeat("-", 16) + "\n")
	for index, item := range report.Improvements {
		fmt.Fprintf(&b, "%d. [%s][%s] %s\n", index+1, styles.Severity(item.Severity), item.Pillar, item.Recommendation)
		if len(item.Evidence) > 0 {
			snippet := item.Evidence[0]
			if snippet.Path != "" {
				fmt.Fprintf(&b, "   exemplo: %s", snippet.Path)
				if snippet.StartLine > 0 {
					fmt.Fprintf(&b, ":%d", snippet.StartLine)
				}
				fmt.Fprintln(&b)
			}
			if strings.TrimSpace(snippet.Content) != "" {
				fmt.Fprintf(&b, "   %s\n", indentOneLine(snippet.Content))
			}
		}
		b.WriteString("\n")
	}
	return b.String()
}

func RenderMarkdown(report diagnose.Report) string {
	var b strings.Builder
	fmt.Fprintf(&b, "# gakit diagnose - %s\n\n", report.ProjectSlug)
	fmt.Fprintf(&b, "- Path: `%s`\n", report.Path)
	fmt.Fprintf(&b, "- Scanned at: `%s`\n", report.ScannedAt.Format("2006-01-02T15:04:05Z07:00"))
	fmt.Fprintf(&b, "- Files scanned: `%d`\n", report.FileCount)
	fmt.Fprintf(&b, "- Elapsed: `%s`\n", report.Elapsed.Round(1_000_000))
	fmt.Fprintf(&b, "- gakit version: `%s`\n", version.String())
	fmt.Fprintf(&b, "- Global weighted score: `%d` (`%s`)\n\n", report.GlobalScore, report.Band)
	b.WriteString("| Pilar | Score | Faixa | Peso | Justificativa |\n")
	b.WriteString("| --- | ---: | --- | ---: | --- |\n")
	for _, category := range report.Categories {
		fmt.Fprintf(&b, "| %s | %d | %s | %.0f%% | %s |\n", category.Name, category.Score, diagnose.BandForScore(category.Score), category.Weight*100, escapeTable(category.Summary))
	}
	b.WriteString("\n## Itens a melhorar\n\n")
	if len(report.Improvements) == 0 {
		b.WriteString("Nenhum item critico detectado.\n")
		return b.String()
	}
	for index, item := range report.Improvements {
		fmt.Fprintf(&b, "%d. **[%s][%s] %s**\n\n", index+1, item.Severity, item.Pillar, item.Title)
		fmt.Fprintf(&b, "   Recomendacao: %s\n\n", item.Recommendation)
		for _, snippet := range item.Evidence {
			if snippet.Path == "" {
				continue
			}
			fmt.Fprintf(&b, "   Evidencia: `%s", snippet.Path)
			if snippet.StartLine > 0 {
				fmt.Fprintf(&b, ":%d-%d", snippet.StartLine, snippet.EndLine)
			}
			b.WriteString("`\n\n")
			if strings.TrimSpace(snippet.Content) != "" {
				b.WriteString("```text\n")
				b.WriteString(snippet.Content)
				if !strings.HasSuffix(snippet.Content, "\n") {
					b.WriteString("\n")
				}
				b.WriteString("```\n\n")
			}
		}
	}
	return b.String()
}

func truncate(value string, max int) string {
	if len(value) <= max {
		return value
	}
	if max <= 1 {
		return value[:max]
	}
	return value[:max-1] + "…"
}

func indentOneLine(value string) string {
	value = strings.ReplaceAll(value, "\n", " / ")
	return truncate(value, 100)
}

func escapeTable(value string) string {
	return strings.ReplaceAll(value, "|", "\\|")
}
