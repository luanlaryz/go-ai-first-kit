package categories

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/inventa-co/go-ai-first-kit/internal/diagnose"
)

func TestSecurityCheckerFlagsMissingPolicy(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "Makefile"), []byte("test:\n\tgo test ./...\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	inv, err := diagnose.NewInventory(root)
	if err != nil {
		t.Fatal(err)
	}
	result := securityChecker{}.Check(inv)
	if result.CoverageScore >= 50 {
		t.Fatalf("expected low coverage, got %d", result.CoverageScore)
	}
	if len(result.Findings) == 0 {
		t.Fatal("expected findings")
	}
}
