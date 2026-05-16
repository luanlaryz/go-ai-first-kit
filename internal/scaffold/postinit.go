package scaffold

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func postInit(target string, initGit bool) ([]string, error) {
	var warnings []string

	tmplMod := filepath.Join(target, "go.mod.tmpl")
	goMod := filepath.Join(target, "go.mod")
	if _, err := os.Stat(tmplMod); err == nil {
		if _, goModErr := os.Stat(goMod); os.IsNotExist(goModErr) {
			if err := os.Rename(tmplMod, goMod); err != nil {
				return warnings, err
			}
		}
	}

	if err := chmodGeneratedScripts(target); err != nil {
		return warnings, err
	}

	if initGit {
		if _, err := os.Stat(filepath.Join(target, ".git")); os.IsNotExist(err) {
			cmd := exec.Command("git", "-C", target, "init", "-b", "main")
			if out, err := cmd.CombinedOutput(); err != nil {
				warnings = append(warnings, "git init falhou: "+strings.TrimSpace(string(out)))
			}
		}
	}

	return warnings, nil
}

func chmodGeneratedScripts(target string) error {
	return filepath.WalkDir(target, func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".sh") {
			return nil
		}
		rel, err := filepath.Rel(target, path)
		if err != nil {
			return err
		}
		rel = filepath.ToSlash(rel)
		if strings.HasPrefix(rel, "scripts/") || strings.HasPrefix(rel, ".cursor/hooks/") {
			return os.Chmod(path, 0o755)
		}
		return nil
	})
}
