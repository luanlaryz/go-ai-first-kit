package categories

import (
	"go/parser"
	"go/token"
	"path/filepath"
	"strings"

	"github.com/inventa-co/go-ai-first-kit/internal/diagnose"
)

type hexagonalChecker struct{}

func (hexagonalChecker) Check(inv diagnose.Inventory) diagnose.CategoryResult {
	const pillar = "Hexagonal Architecture"
	requiredDirs := []string{"pkg", "internal"}
	requiredFiles := []string{"doc.go", "skills/01-hexagonal-architecture/SKILL.md", "skills/17-solid-go-ports/SKILL.md"}
	present := countDirs(inv, requiredDirs) + countFiles(inv, requiredFiles)
	total := len(requiredDirs) + len(requiredFiles)

	qualityChecks := 0
	qualityTotal := 4
	if inv.DirExists("pkg") && inv.DirExists("internal") {
		qualityChecks++
	}
	if inv.Contains("AGENTS.md", "pkg/app", "internal/runtime") {
		qualityChecks++
	}
	if missing := missingPkgDocs(inv); len(missing) == 0 {
		qualityChecks++
	}
	illegalImports := internalImportsFromPkg(inv)
	if len(illegalImports) == 0 {
		qualityChecks++
	}

	findings := make([]diagnose.Finding, 0)
	if !inv.DirExists("pkg") || !inv.DirExists("internal") {
		findings = append(findings, finding(pillar, diagnose.SeverityCritical, "Fronteiras pkg/internal ausentes", "Separar contratos publicos em pkg/ e detalhes operacionais em internal/."))
	}
	if len(illegalImports) > 0 {
		findings = append(findings, finding(pillar, diagnose.SeverityCritical, "Import indevido de internal em pkg", "Remover imports de internal/* dentro de pkg/*, exceto ponte publica aprovada.", diagnose.Snippet{Path: illegalImports[0], Content: "Arquivo em pkg importa internal/*."}))
	}
	if missing := missingPkgDocs(inv); len(missing) > 0 {
		findings = append(findings, finding(pillar, diagnose.SeverityWarn, "Pacotes publicos sem doc.go", "Adicionar doc.go para cada pacote publico em pkg/.", diagnose.Snippet{Path: missing[0], Content: "Diretorio publico sem doc.go."}))
	}

	return diagnose.CategoryResult{
		Name:          pillar,
		Weight:        0.15,
		CoverageScore: diagnose.CoverageScore(present, total),
		QualityScore:  diagnose.CoverageScore(qualityChecks, qualityTotal),
		Summary:       "Fronteiras pkg/internal, doc.go e imports publicos",
		Findings:      findings,
	}
}

func missingPkgDocs(inv diagnose.Inventory) []string {
	seen := make(map[string]struct{})
	for _, file := range inv.FilesWithPrefix("pkg") {
		if strings.HasSuffix(file, ".go") && !strings.HasSuffix(file, "_test.go") {
			seen[filepath.ToSlash(filepath.Dir(file))] = struct{}{}
		}
	}
	var missing []string
	for dir := range seen {
		if !inv.FileExists(filepath.ToSlash(filepath.Join(dir, "doc.go"))) {
			missing = append(missing, dir)
		}
	}
	return missing
}

func internalImportsFromPkg(inv diagnose.Inventory) []string {
	var offenders []string
	fset := token.NewFileSet()
	for _, file := range inv.FilesWithPrefix("pkg") {
		if !strings.HasSuffix(file, ".go") || strings.HasSuffix(file, "_test.go") {
			continue
		}
		parsed, err := parser.ParseFile(fset, filepath.Join(inv.Root, filepath.FromSlash(file)), nil, parser.ImportsOnly)
		if err != nil {
			continue
		}
		for _, imp := range parsed.Imports {
			if strings.Contains(strings.Trim(imp.Path.Value, `"`), "/internal/") {
				offenders = append(offenders, file)
				break
			}
		}
	}
	return offenders
}
