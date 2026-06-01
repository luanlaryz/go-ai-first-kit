package scaffold

import (
	"context"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
	"unicode/utf8"
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
	assertFile(t, target, "skills/26-backlog-item-intake/SKILL.md")
	assertFile(t, target, "skills/26-backlog-item-intake/references/backlog-classification-rules.md")
	assertFile(t, target, "docs/backlog/Backlog.md")
	assertFile(t, target, "docs/decisions/README.md")
	assertFile(t, target, "docs/decisions/0001-record-architecture-decisions.md")
	assertFile(t, target, "docs/release-versioning-policy.md")
	assertFile(t, target, "docs/release-notes-policy.md")
	assertFile(t, target, "docs/release-checklist.md")
	assertFile(t, target, "docs/plans/.gitkeep")
	assertFile(t, target, "docs/reports/.gitkeep")
	assertFile(t, target, "CHANGELOG.md")
	assertNoPlaceholder(t, filepath.Join(target, "AGENTS.md"))
	assertNoUnresolvedPlaceholders(t, target)
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

func assertNoUnresolvedPlaceholders(t *testing.T, root string) {
	t.Helper()
	re := regexp.MustCompile(`\{\{[A-Z_]+\}\}`)
	err := filepath.Walk(root, func(p string, info os.FileInfo, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if info.IsDir() {
			return nil
		}
		data, err := os.ReadFile(p)
		if err != nil {
			return err
		}
		if !utf8.Valid(data) {
			return nil
		}
		if loc := re.Find(data); loc != nil {
			rel, _ := filepath.Rel(root, p)
			t.Errorf("unresolved placeholder %q in %s", string(loc), rel)
		}
		return nil
	})
	if err != nil {
		t.Fatalf("walk %s: %v", root, err)
	}
}
