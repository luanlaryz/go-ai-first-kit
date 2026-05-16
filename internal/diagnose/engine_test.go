package diagnose_test

import (
	"context"
	"testing"

	"github.com/inventa-co/go-ai-first-kit/internal/diagnose"
	"github.com/inventa-co/go-ai-first-kit/internal/diagnose/categories"
)

func TestRunAgainstTemplate(t *testing.T) {
	report, err := diagnose.Run(context.Background(), "../../template", categories.Default())
	if err != nil {
		t.Fatalf("diagnose template: %v", err)
	}
	if report.FileCount == 0 {
		t.Fatal("expected scanned files")
	}
	if report.GlobalScore < 65 {
		t.Fatalf("expected template score >= 65, got %d", report.GlobalScore)
	}
	if len(report.Categories) != 7 {
		t.Fatalf("expected 7 categories, got %d", len(report.Categories))
	}
}
