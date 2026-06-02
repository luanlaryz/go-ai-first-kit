package templatecatalog

import (
	"fmt"
	"strings"
)

// EntryType identifies whether an Entry is a file or a directory.
type EntryType string

const (
	EntryTypeFile EntryType = "file"
	EntryTypeDir  EntryType = "dir"
)

// DefaultTemplate is the name of the only template currently embedded in the kit.
const DefaultTemplate = "default"

// Entry is a single file or directory inside a template.
type Entry struct {
	Path string    `json:"path"`
	Type EntryType `json:"type"`
	Size int64     `json:"size,omitempty"`
}

// Inventory is the full listing of a template's contents.
type Inventory struct {
	Name    string  `json:"name"`
	Root    string  `json:"root"`
	Entries []Entry `json:"entries"`
}

// Available returns the known template names.
func Available() []string {
	return []string{DefaultTemplate}
}

// Resolve normalizes an empty name to the default template and validates it.
func Resolve(name string) (string, error) {
	if name == "" {
		return DefaultTemplate, nil
	}
	if name != DefaultTemplate {
		return "", fmt.Errorf("template %q not found; available templates: %s", name, strings.Join(Available(), ", "))
	}
	return name, nil
}
