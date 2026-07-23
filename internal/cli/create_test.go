package cli

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCreateNonInteractive(t *testing.T) {
	target := filepath.Join(t.TempDir(), "myapp")
	out, err := executeRoot(t,
		"create", target,
		"--slug", "myapp",
		"--title", "My App",
		"--module", "github.com/acme/myapp",
		"--description", "Example app",
		"--author", "Acme",
		"--non-interactive",
	)
	if err != nil {
		t.Fatalf("create failed: %v\n%s", err, out)
	}
	assertProjectCreated(t, target)
}

func TestCreateWithTemplateDefault(t *testing.T) {
	target := filepath.Join(t.TempDir(), "myapp")
	out, err := executeRoot(t,
		"create", target,
		"--template", "default",
		"--slug", "myapp",
		"--title", "My App",
		"--module", "github.com/acme/myapp",
		"--description", "Example app",
		"--author", "Acme",
		"--non-interactive",
	)
	if err != nil {
		t.Fatalf("create failed: %v\n%s", err, out)
	}
	assertProjectCreated(t, target)
}

func TestCreateWithUnknownTemplateFails(t *testing.T) {
	target := filepath.Join(t.TempDir(), "myapp")
	_, err := executeRoot(t,
		"create", target,
		"--template", "api",
		"--slug", "myapp",
		"--title", "My App",
		"--module", "github.com/acme/myapp",
		"--description", "Example app",
		"--author", "Acme",
		"--non-interactive",
	)
	if err == nil {
		t.Fatal("expected error for unknown template")
	}
	want := `template "api" not found; available templates: default`
	if err.Error() != want {
		t.Fatalf("expected error %q, got %q", want, err.Error())
	}
}

func assertProjectCreated(t *testing.T, target string) {
	t.Helper()
	goModPath := filepath.Join(target, "go.mod")
	if _, err := os.Stat(goModPath); err != nil {
		t.Fatalf("expected go.mod to exist: %v", err)
	}
	goMod, err := os.ReadFile(goModPath)
	if err != nil {
		t.Fatalf("read go.mod: %v", err)
	}
	if !strings.Contains(string(goMod), "go 1.26.4") {
		t.Error("expected generated go.mod to require Go 1.26.4")
	}
	data, err := os.ReadFile(filepath.Join(target, "AGENTS.md"))
	if err != nil {
		t.Fatalf("read AGENTS.md: %v", err)
	}
	if strings.Contains(string(data), "{{PROJECT_SLUG}}") {
		t.Error("expected placeholders to be replaced in AGENTS.md")
	}
	inventory, err := os.ReadFile(filepath.Join(target, "docs", "ai", "maturity-inventory.md"))
	if err != nil {
		t.Fatalf("read maturity inventory: %v", err)
	}
	if !strings.Contains(string(inventory), "## Baseline entregue") {
		t.Error("expected generated maturity inventory to include baseline section")
	}
	catalog, err := os.ReadFile(filepath.Join(target, "docs", "ai", "capabilities.md"))
	if err != nil {
		t.Fatalf("read capabilities catalog: %v", err)
	}
	if !strings.Contains(string(catalog), "## Capacidades") {
		t.Error("expected generated capabilities catalog to include capacities section")
	}
	if _, err := os.Stat(filepath.Join(target, "docs", "journeys", "README.md")); err != nil {
		t.Errorf("expected generated journeys hub: %v", err)
	}
}
