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
	if _, err := os.Stat(filepath.Join(target, "go.mod")); err != nil {
		t.Fatalf("expected go.mod to exist: %v", err)
	}
	data, err := os.ReadFile(filepath.Join(target, "AGENTS.md"))
	if err != nil {
		t.Fatalf("read AGENTS.md: %v", err)
	}
	if strings.Contains(string(data), "{{PROJECT_SLUG}}") {
		t.Error("expected placeholders to be replaced in AGENTS.md")
	}
}
