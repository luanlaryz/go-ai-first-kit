package categories

import "github.com/inventa-co/go-ai-first-kit/internal/diagnose"

func Default() []diagnose.Checker {
	return []diagnose.Checker{
		dxChecker{},
		aiFirstChecker{},
		securityChecker{},
		hexagonalChecker{},
		openAPIChecker{},
		docsChecker{},
		governanceChecker{},
	}
}

func finding(pillar string, severity diagnose.Severity, title, recommendation string, evidence ...diagnose.Snippet) diagnose.Finding {
	return diagnose.Finding{
		Pillar:         pillar,
		Severity:       severity,
		Title:          title,
		Recommendation: recommendation,
		Evidence:       evidence,
	}
}

func countFiles(inv diagnose.Inventory, files []string) int {
	count := 0
	for _, file := range files {
		if inv.FileExists(file) {
			count++
		}
	}
	return count
}

func countDirs(inv diagnose.Inventory, dirs []string) int {
	count := 0
	for _, dir := range dirs {
		if inv.DirExists(dir) {
			count++
		}
	}
	return count
}
