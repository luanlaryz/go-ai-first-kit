package gakit

import (
	"os/exec"
	"testing"
)

func TestBootstrapPromptIsCurrent(t *testing.T) {
	t.Helper()

	cmd := exec.Command("bash", "scripts/build-master-prompt.sh", "--check")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("bootstrap prompt is outdated: %v\n%s", err, output)
	}
}
