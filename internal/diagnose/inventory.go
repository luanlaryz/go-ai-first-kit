package diagnose

import (
	"bufio"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
)

type Inventory struct {
	Root     string
	Files    map[string]struct{}
	Dirs     map[string]struct{}
	contents map[string]string
}

func NewInventory(root string) (Inventory, error) {
	absRoot, err := filepath.Abs(root)
	if err != nil {
		return Inventory{}, err
	}
	inv := Inventory{
		Root:     absRoot,
		Files:    make(map[string]struct{}),
		Dirs:     make(map[string]struct{}),
		contents: make(map[string]string),
	}
	err = filepath.WalkDir(absRoot, func(filePath string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() && shouldSkipDir(entry.Name()) {
			if filePath == absRoot {
				return nil
			}
			return filepath.SkipDir
		}
		rel, err := filepath.Rel(absRoot, filePath)
		if err != nil {
			return err
		}
		if rel == "." {
			return nil
		}
		rel = filepath.ToSlash(rel)
		if entry.IsDir() {
			inv.Dirs[rel] = struct{}{}
		} else {
			inv.Files[rel] = struct{}{}
		}
		return nil
	})
	return inv, err
}

func shouldSkipDir(name string) bool {
	switch name {
	case ".git", "vendor", "node_modules", ".next", "dist", "build", "coverage":
		return true
	default:
		return false
	}
}

func (i Inventory) FileCount() int {
	return len(i.Files)
}

func (i Inventory) FileExists(rel string) bool {
	_, ok := i.Files[path.Clean(filepath.ToSlash(rel))]
	return ok
}

func (i Inventory) DirExists(rel string) bool {
	_, ok := i.Dirs[path.Clean(filepath.ToSlash(rel))]
	return ok
}

func (i Inventory) FilesWithPrefix(prefix string) []string {
	prefix = strings.TrimSuffix(filepath.ToSlash(prefix), "/") + "/"
	var out []string
	for file := range i.Files {
		if strings.HasPrefix(file, prefix) {
			out = append(out, file)
		}
	}
	sort.Strings(out)
	return out
}

func (i Inventory) FilesWithSuffix(suffix string) []string {
	var out []string
	for file := range i.Files {
		if strings.HasSuffix(file, suffix) {
			out = append(out, file)
		}
	}
	sort.Strings(out)
	return out
}

func (i Inventory) Read(rel string) (string, error) {
	rel = path.Clean(filepath.ToSlash(rel))
	if text, ok := i.contents[rel]; ok {
		return text, nil
	}
	data, err := os.ReadFile(filepath.Join(i.Root, filepath.FromSlash(rel)))
	if err != nil {
		return "", err
	}
	text := string(data)
	i.contents[rel] = text
	return text, nil
}

func (i Inventory) Contains(rel string, needles ...string) bool {
	text, err := i.Read(rel)
	if err != nil {
		return false
	}
	for _, needle := range needles {
		if !strings.Contains(text, needle) {
			return false
		}
	}
	return true
}

func (i Inventory) CountMarkdownFiles(prefix string) int {
	count := 0
	for _, file := range i.FilesWithPrefix(prefix) {
		if strings.HasSuffix(file, ".md") {
			count++
		}
	}
	return count
}

func (i Inventory) SnippetForPath(rel string, fallback string) Snippet {
	text, err := i.Read(rel)
	if err != nil {
		return Snippet{Path: rel, StartLine: 0, EndLine: 0, Content: fallback}
	}
	scanner := bufio.NewScanner(strings.NewReader(text))
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		if len(lines) >= 5 {
			break
		}
	}
	if len(lines) == 0 {
		return Snippet{Path: rel, StartLine: 0, EndLine: 0, Content: fallback}
	}
	return Snippet{Path: rel, StartLine: 1, EndLine: len(lines), Content: strings.Join(lines, "\n")}
}
