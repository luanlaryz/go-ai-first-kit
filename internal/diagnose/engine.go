package diagnose

import (
	"context"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

func Run(ctx context.Context, root string, checkers []Checker) (Report, error) {
	start := time.Now()
	inv, err := NewInventory(root)
	if err != nil {
		return Report{}, err
	}

	categories := make([]CategoryResult, 0, len(checkers))
	for _, checker := range checkers {
		select {
		case <-ctx.Done():
			return Report{}, ctx.Err()
		default:
		}
		result := checker.Check(inv)
		result.Score = CategoryScore(result.CoverageScore, result.QualityScore)
		categories = append(categories, result)
	}

	report := Report{
		ProjectSlug: projectSlug(inv),
		Path:        inv.Root,
		ScannedAt:   start,
		Elapsed:     time.Since(start),
		FileCount:   inv.FileCount(),
		Categories:  categories,
	}
	report.GlobalScore = weightedScore(categories)
	report.Band = BandForScore(report.GlobalScore)
	report.Improvements = improvements(categories)
	return report, nil
}

func weightedScore(categories []CategoryResult) int {
	var totalWeight float64
	var total float64
	for _, category := range categories {
		totalWeight += category.Weight
		total += float64(category.Score) * category.Weight
	}
	if totalWeight == 0 {
		return 0
	}
	return clampScore(int(total / totalWeight))
}

func improvements(categories []CategoryResult) []Finding {
	var findings []Finding
	for _, category := range categories {
		findings = append(findings, category.Findings...)
	}
	sort.SliceStable(findings, func(i, j int) bool {
		return severityRank(findings[i].Severity) > severityRank(findings[j].Severity)
	})
	return findings
}

func severityRank(severity Severity) int {
	switch severity {
	case SeverityCritical:
		return 3
	case SeverityWarn:
		return 2
	default:
		return 1
	}
}

func projectSlug(inv Inventory) string {
	if inv.FileExists("go.mod") {
		if text, err := inv.Read("go.mod"); err == nil {
			for _, line := range strings.Split(text, "\n") {
				fields := strings.Fields(line)
				if len(fields) == 2 && fields[0] == "module" {
					parts := strings.Split(strings.TrimSpace(fields[1]), "/")
					return parts[len(parts)-1]
				}
			}
		}
	}
	return filepath.Base(inv.Root)
}
