package categories

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/inventa-co/go-ai-first-kit/internal/diagnose"
)

func writeTestFile(t *testing.T, root, rel, content string) {
	t.Helper()
	full := filepath.Join(root, filepath.FromSlash(rel))
	if err := os.MkdirAll(filepath.Dir(full), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(full, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}

func newTestInventory(t *testing.T, root string) diagnose.Inventory {
	t.Helper()
	inv, err := diagnose.NewInventory(root)
	if err != nil {
		t.Fatal(err)
	}
	return inv
}

func hasFinding(findings []diagnose.Finding, title string) bool {
	for _, f := range findings {
		if f.Title == title {
			return true
		}
	}
	return false
}

func TestAIFirstFlagsMissingBacklog(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, "AGENTS.md", "# AGENTS\nSpec Driven Development\n")

	result := aiFirstChecker{}.Check(newTestInventory(t, root))

	if !hasFinding(result.Findings, "Backlog canonico ausente") {
		t.Fatalf("expected backlog finding, got %+v", result.Findings)
	}
}

func TestAIFirstRewardsBacklog(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, "AGENTS.md", "# AGENTS\nSpec Driven Development\n")

	without := aiFirstChecker{}.Check(newTestInventory(t, root))

	writeTestFile(t, root, "docs/backlog/Backlog.md", "# My App Backlog\n\n## Itens\n")
	writeTestFile(t, root, "skills/26-backlog-item-intake/SKILL.md", "# Backlog Item Intake\n")

	with := aiFirstChecker{}.Check(newTestInventory(t, root))

	if with.QualityScore <= without.QualityScore {
		t.Fatalf("expected backlog to improve quality: without=%d with=%d", without.QualityScore, with.QualityScore)
	}
	if with.CoverageScore <= without.CoverageScore {
		t.Fatalf("expected backlog to improve coverage: without=%d with=%d", without.CoverageScore, with.CoverageScore)
	}
	if hasFinding(with.Findings, "Backlog canonico ausente") {
		t.Fatal("did not expect backlog finding when backlog present")
	}
}
