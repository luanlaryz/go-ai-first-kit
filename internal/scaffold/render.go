package scaffold

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"unicode/utf8"

	kit "github.com/inventa-co/go-ai-first-kit"
)

type Options struct {
	Force   bool
	InitGit bool
}

type Result struct {
	TargetDir string
	Files     []string
	Warnings  []string
}

func Render(ctx context.Context, targetDir string, params Params, opts Options) (Result, error) {
	params.ApplyDefaults()
	if err := params.Validate(); err != nil {
		return Result{}, err
	}
	if strings.TrimSpace(targetDir) == "" {
		return Result{}, fmt.Errorf("target dir is required")
	}

	absTarget, err := filepath.Abs(targetDir)
	if err != nil {
		return Result{}, err
	}
	if err := ensureWritableTarget(absTarget, opts.Force); err != nil {
		return Result{}, err
	}

	repls := replacements(params)
	result := Result{TargetDir: absTarget}
	err = fs.WalkDir(kit.TemplateFS, kit.TemplateRoot, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		rel, err := filepath.Rel(kit.TemplateRoot, path)
		if err != nil {
			return err
		}
		if rel == "." {
			return nil
		}

		outRel := renderPath(filepath.FromSlash(rel), repls)
		outPath := filepath.Join(absTarget, outRel)
		if entry.IsDir() {
			return os.MkdirAll(outPath, 0o755)
		}

		data, err := kit.TemplateFS.ReadFile(path)
		if err != nil {
			return err
		}
		if err := os.MkdirAll(filepath.Dir(outPath), 0o755); err != nil {
			return err
		}
		if utf8.Valid(data) {
			text := renderString(string(data), repls)
			data = []byte(text)
		}
		if err := os.WriteFile(outPath, data, 0o644); err != nil {
			return err
		}
		result.Files = append(result.Files, outRel)
		return nil
	})
	if err != nil {
		return Result{}, err
	}

	warnings, err := postInit(absTarget, opts.InitGit)
	if err != nil {
		return Result{}, err
	}
	result.Warnings = append(result.Warnings, warnings...)
	sort.Strings(result.Files)
	return result, nil
}

func ensureWritableTarget(target string, force bool) error {
	entries, err := os.ReadDir(target)
	if err != nil {
		if os.IsNotExist(err) {
			return os.MkdirAll(target, 0o755)
		}
		return err
	}
	if len(entries) > 0 && !force {
		return fmt.Errorf("target dir %s nao esta vazio; use --force para sobrescrever arquivos", target)
	}
	return nil
}

func replacements(params Params) map[string]string {
	upstreamLower := strings.ToLower(params.UpstreamName)
	return map[string]string{
		"PROJECT_SLUG":          params.ProjectSlug,
		"PROJECT_TITLE":         params.ProjectTitle,
		"MODULE_PATH":           params.ModulePath,
		"PROJECT_DESCRIPTION":   params.ProjectDescription,
		"AUTHOR_NAME":           params.AuthorName,
		"LICENSE_NAME":          params.LicenseName,
		"UPSTREAM_NAME":         params.UpstreamName,
		"UPSTREAM_NAME_LOWER":   upstreamLower,
		"UPSTREAM_OPS_NAME":     params.UpstreamName + "Ops",
		"DOMAIN_ACTOR":          "domain actor",
		"DOMAIN_ACTOR_TITLE":    "Domain Actor",
		"EXTERNAL_SYSTEM":       "External System",
		"EXTERNAL_SYSTEM_LOWER": "external system",
		"DOMAIN_ENTITY_SET":     "domain entities",
	}
}

func renderPath(path string, repls map[string]string) string {
	parts := strings.Split(path, string(filepath.Separator))
	for i, part := range parts {
		parts[i] = renderString(part, repls)
	}
	return filepath.Join(parts...)
}

func renderString(value string, repls map[string]string) string {
	for key, replacement := range repls {
		value = strings.ReplaceAll(value, "{{"+key+"}}", replacement)
	}
	return value
}
