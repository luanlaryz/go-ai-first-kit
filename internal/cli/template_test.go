package cli

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestTemplateListDefault(t *testing.T) {
	out, err := executeRoot(t, "template", "list")
	if err != nil {
		t.Fatalf("template list failed: %v", err)
	}
	if !strings.Contains(out, "Template: default") {
		t.Errorf("expected header 'Template: default', got:\n%s", out)
	}
	if !strings.Contains(out, "Directories:") {
		t.Errorf("expected Directories section, got:\n%s", out)
	}
	if !strings.Contains(out, "Files:") {
		t.Errorf("expected Files section, got:\n%s", out)
	}
	if !strings.Contains(out, "AGENTS.md") {
		t.Errorf("expected AGENTS.md in output, got:\n%s", out)
	}
}

func TestTemplateListJSON(t *testing.T) {
	out, err := executeRoot(t, "template", "list", "--json")
	if err != nil {
		t.Fatalf("template list --json failed: %v", err)
	}
	if !json.Valid([]byte(out)) {
		t.Fatalf("expected valid JSON, got:\n%s", out)
	}
	var inv struct {
		Name    string `json:"name"`
		Entries []struct {
			Path string `json:"path"`
			Type string `json:"type"`
		} `json:"entries"`
	}
	if err := json.Unmarshal([]byte(out), &inv); err != nil {
		t.Fatalf("unmarshal JSON: %v", err)
	}
	if inv.Name != "default" {
		t.Errorf("expected name default, got %q", inv.Name)
	}
	if len(inv.Entries) == 0 {
		t.Error("expected entries in JSON output")
	}
}

func TestTemplateListTree(t *testing.T) {
	out, err := executeRoot(t, "template", "list", "--tree")
	if err != nil {
		t.Fatalf("template list --tree failed: %v", err)
	}
	if !strings.Contains(out, "default") {
		t.Errorf("expected tree root 'default', got:\n%s", out)
	}
	if !strings.Contains(out, "AGENTS.md") {
		t.Errorf("expected AGENTS.md in tree, got:\n%s", out)
	}
}

func TestTemplateListFilesOnly(t *testing.T) {
	out, err := executeRoot(t, "template", "list", "--files-only")
	if err != nil {
		t.Fatalf("template list --files-only failed: %v", err)
	}
	if strings.Contains(out, "Directories:") {
		t.Errorf("did not expect Directories section, got:\n%s", out)
	}
	if !strings.Contains(out, "Files:") {
		t.Errorf("expected Files section, got:\n%s", out)
	}
}

func TestTemplateListDirsOnly(t *testing.T) {
	out, err := executeRoot(t, "template", "list", "--dirs-only")
	if err != nil {
		t.Fatalf("template list --dirs-only failed: %v", err)
	}
	if strings.Contains(out, "Files:") {
		t.Errorf("did not expect Files section, got:\n%s", out)
	}
	if !strings.Contains(out, "Directories:") {
		t.Errorf("expected Directories section, got:\n%s", out)
	}
}

func TestTemplateListFilesAndDirsOnlyConflict(t *testing.T) {
	_, err := executeRoot(t, "template", "list", "--files-only", "--dirs-only")
	if err == nil {
		t.Fatal("expected error when combining --files-only and --dirs-only")
	}
}
