package cli

import (
	"bytes"
	"strings"
	"testing"
)

// executeRoot runs the root command with the given args, capturing stdout and
// stderr into a single buffer for assertions.
func executeRoot(t *testing.T, args ...string) (string, error) {
	t.Helper()
	if args == nil {
		args = []string{}
	}
	cmd := newRootCmd()
	var buf bytes.Buffer
	cmd.SetOut(&buf)
	cmd.SetErr(&buf)
	cmd.SetArgs(args)
	err := cmd.Execute()
	return buf.String(), err
}

func TestRootNoArgsShowsTip(t *testing.T) {
	out, err := executeRoot(t)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !strings.Contains(out, "gakit create ./myapp") {
		t.Errorf("expected tip to contain create example, got:\n%s", out)
	}
	if !strings.Contains(out, "gakit help") {
		t.Errorf("expected tip to mention gakit help, got:\n%s", out)
	}
}

func TestRootHelpShowsExamples(t *testing.T) {
	out, err := executeRoot(t, "--help")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !strings.Contains(out, "gakit create ./myapp") {
		t.Errorf("expected help to contain create example, got:\n%s", out)
	}
	if !strings.Contains(out, "template list") {
		t.Errorf("expected help to mention template list, got:\n%s", out)
	}
}
