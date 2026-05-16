package categories

import (
	"os"
	"regexp"
	"strings"
	"testing"
)

func TestRequiredComplianceFilesMatchShellScript(t *testing.T) {
	data, err := os.ReadFile("../../../template/scripts/check-compliance.sh")
	if err != nil {
		t.Fatalf("read check-compliance.sh: %v", err)
	}
	got := requiredFilesFromShell(string(data))
	want := RequiredComplianceFiles
	if strings.Join(got, "\n") != strings.Join(want, "\n") {
		t.Fatalf("required files diverged\nshell:\n%s\n\ngo:\n%s", strings.Join(got, "\n"), strings.Join(want, "\n"))
	}
}

func requiredFilesFromShell(script string) []string {
	re := regexp.MustCompile(`(?s)required_files=\(\n(.*?)\n\)`)
	match := re.FindStringSubmatch(script)
	if len(match) != 2 {
		return nil
	}
	var out []string
	for _, line := range strings.Split(match[1], "\n") {
		line = strings.TrimSpace(line)
		line = strings.Trim(line, `"`)
		if line != "" {
			out = append(out, line)
		}
	}
	return out
}
