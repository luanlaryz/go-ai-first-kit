package templatecatalog

import (
	"context"
	"sort"
	"testing"
)

func TestListDefaultContainsKeyEntries(t *testing.T) {
	inv, err := List(context.Background(), "default")
	if err != nil {
		t.Fatalf("List failed: %v", err)
	}
	if inv.Name != "default" {
		t.Fatalf("expected name %q, got %q", "default", inv.Name)
	}
	if len(inv.Entries) == 0 {
		t.Fatal("expected entries")
	}

	byPath := make(map[string]Entry, len(inv.Entries))
	for _, e := range inv.Entries {
		byPath[e.Path] = e
	}
	for _, want := range []string{"AGENTS.md", "README.md", "go.mod.tmpl", "docs", "skills"} {
		if _, ok := byPath[want]; !ok {
			t.Errorf("expected entry %q in inventory", want)
		}
	}

	if mod, ok := byPath["go.mod.tmpl"]; ok {
		if mod.Type != EntryTypeFile {
			t.Errorf("expected go.mod.tmpl to be a file, got %q", mod.Type)
		}
		if mod.Size <= 0 {
			t.Errorf("expected go.mod.tmpl to have a positive size, got %d", mod.Size)
		}
	}
	if docs, ok := byPath["docs"]; ok && docs.Type != EntryTypeDir {
		t.Errorf("expected docs to be a dir, got %q", docs.Type)
	}
}

func TestListEntriesSortedByPath(t *testing.T) {
	inv, err := List(context.Background(), "default")
	if err != nil {
		t.Fatalf("List failed: %v", err)
	}
	if !sort.SliceIsSorted(inv.Entries, func(i, j int) bool {
		return inv.Entries[i].Path < inv.Entries[j].Path
	}) {
		t.Error("expected entries to be sorted by path")
	}
}

func TestListEmptyNameUsesDefault(t *testing.T) {
	inv, err := List(context.Background(), "")
	if err != nil {
		t.Fatalf("List failed: %v", err)
	}
	if inv.Name != "default" {
		t.Fatalf("expected default template, got %q", inv.Name)
	}
}

func TestListUnknownTemplateReturnsError(t *testing.T) {
	_, err := List(context.Background(), "api")
	if err == nil {
		t.Fatal("expected error for unknown template")
	}
	want := `template "api" not found; available templates: default`
	if err.Error() != want {
		t.Fatalf("expected error %q, got %q", want, err.Error())
	}
}
