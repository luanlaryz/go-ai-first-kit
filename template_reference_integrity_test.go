package gakit

import (
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

var (
	templateSpecPathPattern = regexp.MustCompile(`specs/[A-Za-z0-9_.{}-]+\.md`)
	templateSpecNumPattern  = regexp.MustCompile(`\bSpec (\d{3})\b`)
	markdownLinkPattern     = regexp.MustCompile(`\]\(([^)]+)\)`)
)

// templateMarkdownFiles walks template/ and returns every markdown file path
// relative to the template root.
func templateMarkdownFiles(t *testing.T) []string {
	t.Helper()

	var files []string
	err := filepath.WalkDir("template", func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || !strings.HasSuffix(path, ".md") {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		t.Fatalf("walk template: %v", err)
	}
	if len(files) == 0 {
		t.Fatal("expected markdown files in template")
	}
	return files
}

// TestTemplateNormativeSpecReferencesExist garante que todo path specs/*.md
// citado em qualquer documento do template aponta para um arquivo real.
// Specs futuras devem ser descritas como candidatas, sem path concreto.
func TestTemplateNormativeSpecReferencesExist(t *testing.T) {
	for _, source := range templateMarkdownFiles(t) {
		data, err := os.ReadFile(source)
		if err != nil {
			t.Fatalf("read %s: %v", source, err)
		}
		for _, ref := range templateSpecPathPattern.FindAllString(string(data), -1) {
			if _, err := os.Stat(filepath.Join("template", filepath.FromSlash(ref))); err != nil {
				t.Errorf("%s requires missing spec %s: %v", source, ref, err)
			}
		}
	}
}

// TestTemplateSpecNumberReferencesExist garante que referencias textuais no
// formato "Spec NNN" correspondem a uma spec real em template/specs. Isso
// impede que o template volte a citar specs historicas de outro repositorio
// (por exemplo Spec 348/350/700/701/726) como se existissem aqui.
func TestTemplateSpecNumberReferencesExist(t *testing.T) {
	entries, err := os.ReadDir(filepath.Join("template", "specs"))
	if err != nil {
		t.Fatalf("read template/specs: %v", err)
	}
	known := make(map[string]bool, len(entries))
	for _, entry := range entries {
		name := entry.Name()
		if len(name) >= 3 && strings.HasSuffix(name, ".md") {
			known[name[:3]] = true
		}
	}

	for _, source := range templateMarkdownFiles(t) {
		data, err := os.ReadFile(source)
		if err != nil {
			t.Fatalf("read %s: %v", source, err)
		}
		for _, match := range templateSpecNumPattern.FindAllStringSubmatch(string(data), -1) {
			if !known[match[1]] {
				t.Errorf("%s references %q but template/specs has no spec %s", source, match[0], match[1])
			}
		}
	}
}

// TestTemplateMarkdownLinksResolve garante que links Markdown locais do
// template resolvem para arquivos ou diretorios reais. URLs externas e
// ancoras sao ignoradas. Como os placeholders sao substituidos de forma
// identica em paths e conteudos, a validacao na arvore do template tambem
// vale para o projeto renderizado.
func TestTemplateMarkdownLinksResolve(t *testing.T) {
	for _, source := range templateMarkdownFiles(t) {
		data, err := os.ReadFile(source)
		if err != nil {
			t.Fatalf("read %s: %v", source, err)
		}
		for _, match := range markdownLinkPattern.FindAllStringSubmatch(string(data), -1) {
			target := strings.TrimSpace(match[1])
			if target == "" ||
				strings.HasPrefix(target, "http://") ||
				strings.HasPrefix(target, "https://") ||
				strings.HasPrefix(target, "mailto:") ||
				strings.HasPrefix(target, "#") {
				continue
			}
			if idx := strings.Index(target, "#"); idx >= 0 {
				target = target[:idx]
			}
			if target == "" {
				continue
			}
			resolved := filepath.Join(filepath.Dir(source), filepath.FromSlash(target))
			if _, err := os.Stat(resolved); err != nil {
				t.Errorf("%s links to missing target %s: %v", source, match[1], err)
			}
		}
	}
}
