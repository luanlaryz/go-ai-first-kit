package categories

import (
	"testing"

	"github.com/inventa-co/go-ai-first-kit/internal/diagnose"
)

func docsResultForChangelog(t *testing.T, changelog string) diagnose.CategoryResult {
	t.Helper()
	root := t.TempDir()
	writeTestFile(t, root, "README.md", "# App\n\n## Uso\n")
	writeTestFile(t, root, "CONTRIBUTING.md", "# Contributing\n\nPR workflow\n")
	writeTestFile(t, root, "doc.go", "package app\n")
	writeTestFile(t, root, "examples/main.go", "package main\n")
	writeTestFile(t, root, "CHANGELOG.md", changelog)
	return docsChecker{}.Check(newTestInventory(t, root))
}

func TestDocsRewardsKeepAChangelogFormat(t *testing.T) {
	without := docsResultForChangelog(t, "# Changelog\n\n## [Unreleased]\n")
	with := docsResultForChangelog(t, "# Changelog\n\nThe format is based on Keep a Changelog and adheres to Semantic Versioning.\n\n## [Unreleased]\n")

	if with.QualityScore <= without.QualityScore {
		t.Fatalf("expected keep-a-changelog format to improve quality: without=%d with=%d", without.QualityScore, with.QualityScore)
	}
}
