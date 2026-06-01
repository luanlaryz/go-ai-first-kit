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

func TestGovernanceRewardsADRAndReleasePolicy(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, ".github/PULL_REQUEST_TEMPLATE.md", "## Objetivo\n## Specs lidas\n## Checklist\n## Gaps restantes\n")

	without := governanceChecker{}.Check(newTestInventory(t, root))

	writeTestFile(t, root, "docs/decisions/README.md", "# Architecture Decision Records (ADRs)\n")
	writeTestFile(t, root, "docs/release-versioning-policy.md", "# Release Versioning Policy\n")

	with := governanceChecker{}.Check(newTestInventory(t, root))

	if with.QualityScore <= without.QualityScore {
		t.Fatalf("expected ADR/release governance to improve quality: without=%d with=%d", without.QualityScore, with.QualityScore)
	}
	if hasFinding(with.Findings, "Sistema de ADR ausente") {
		t.Fatal("did not expect ADR finding when docs/decisions present")
	}
	if hasFinding(with.Findings, "Politica de release ausente") {
		t.Fatal("did not expect release policy finding when policy present")
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
