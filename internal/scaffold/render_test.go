package scaffold

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRenderTemplate(t *testing.T) {
	target := t.TempDir()
	result, err := Render(context.Background(), target, Params{
		ProjectSlug:        "myapp",
		ProjectTitle:       "My App",
		ModulePath:         "github.com/acme/myapp",
		ProjectDescription: "Example app",
		AuthorName:         "Acme",
		LicenseName:        "MIT",
		UpstreamName:       "none",
	}, Options{Force: true})
	if err != nil {
		t.Fatalf("render failed: %v", err)
	}
	if len(result.Files) == 0 {
		t.Fatal("expected rendered files")
	}
	assertFile(t, target, "go.mod")
	assertFile(t, target, "AGENTS.md")
	assertFile(t, target, ".cursor/rules/myapp.mdc")
	assertNoPlaceholder(t, filepath.Join(target, "AGENTS.md"))
	info, err := os.Stat(filepath.Join(target, "scripts/check-compliance.sh"))
	if err != nil {
		t.Fatalf("stat script: %v", err)
	}
	if info.Mode()&0o111 == 0 {
		t.Fatal("expected generated shell script to be executable")
	}
}

func assertFile(t *testing.T, root, rel string) {
	t.Helper()
	if _, err := os.Stat(filepath.Join(root, filepath.FromSlash(rel))); err != nil {
		t.Fatalf("expected %s: %v", rel, err)
	}
}

func assertNoPlaceholder(t *testing.T, path string) {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	if strings.Contains(string(data), "{{PROJECT_SLUG}}") {
		t.Fatalf("placeholder not rendered in %s", path)
	}
}
