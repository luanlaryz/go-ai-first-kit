package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/inventa-co/go-ai-first-kit/internal/diagnose"
	"github.com/inventa-co/go-ai-first-kit/internal/diagnose/categories"
	"github.com/inventa-co/go-ai-first-kit/internal/ui"
	"github.com/spf13/cobra"
)

type diagnoseOptions struct {
	path           string
	minScore       int
	outDir         string
	reportOnly     bool
	json           bool
	planPrompt     bool
	planPromptOnly bool
}

func newDiagnoseCmd(global *globalOptions) *cobra.Command {
	opts := &diagnoseOptions{}
	cmd := &cobra.Command{
		Use:   "diagnose --path <dir>",
		Short: "Diagnostica maturidade AI-first de um projeto",
		Example: "  gakit diagnose --path .\n" +
			"  gakit diagnose --path . --min-score 80 --report-only --json --out ./reports\n" +
			"  gakit diagnose --path . --plan-prompt --plan-prompt-only --out ./reports",
		RunE: func(cmd *cobra.Command, _ []string) error {
			if strings.TrimSpace(opts.path) == "" {
				return fmt.Errorf("--path e obrigatorio")
			}
			report, err := diagnose.Run(cmd.Context(), opts.path, categories.Default())
			if err != nil {
				return err
			}
			fmt.Fprint(cmd.OutOrStdout(), ui.RenderTerminal(report, global.noColor))

			shouldPersist := opts.reportOnly
			if !shouldPersist {
				confirmed, err := ui.ConfirmPersist(cmd.Context())
				if err != nil {
					return err
				}
				shouldPersist = confirmed
			}
			if shouldPersist {
				paths, err := persistReport(report, opts.outDir, opts.json)
				if err != nil {
					return err
				}
				for _, path := range paths {
					fmt.Fprintf(cmd.OutOrStdout(), "Relatorio salvo em %s\n", path)
				}
			}

			if opts.planPrompt {
				fmt.Fprintln(cmd.OutOrStdout())
				fmt.Fprint(cmd.OutOrStdout(), ui.RenderCorrectionPlanPrompt(report))
			}
			shouldPersistPrompt := opts.planPromptOnly
			if opts.planPrompt && !shouldPersistPrompt {
				confirmed, err := ui.ConfirmPersistPlanPrompt(cmd.Context())
				if err != nil {
					return err
				}
				shouldPersistPrompt = confirmed
			}
			if shouldPersistPrompt {
				path, err := persistPlanPrompt(report, opts.outDir)
				if err != nil {
					return err
				}
				fmt.Fprintf(cmd.OutOrStdout(), "Prompt de plano salvo em %s\n", path)
			}
			if report.GlobalScore < opts.minScore {
				return exitError{code: 1, message: fmt.Sprintf("score global %d abaixo do minimo %d", report.GlobalScore, opts.minScore)}
			}
			return nil
		},
	}
	cmd.Flags().StringVar(&opts.path, "path", "", "diretorio do projeto a diagnosticar")
	cmd.Flags().IntVar(&opts.minScore, "min-score", 0, "score minimo para exit code 0")
	cmd.Flags().StringVar(&opts.outDir, "out", "", "diretorio para salvar relatorios; default e o cwd")
	cmd.Flags().BoolVar(&opts.reportOnly, "report-only", false, "salva o relatorio sem perguntar ao final")
	cmd.Flags().BoolVar(&opts.json, "json", false, "salva tambem uma versao JSON do relatorio")
	cmd.Flags().BoolVar(&opts.planPrompt, "plan-prompt", false, "imprime um prompt para criar plano de correcao dos achados")
	cmd.Flags().BoolVar(&opts.planPromptOnly, "plan-prompt-only", false, "salva o prompt de plano em Markdown sem perguntar")
	return cmd
}

func persistReport(report diagnose.Report, outDir string, writeJSON bool) ([]string, error) {
	if strings.TrimSpace(outDir) == "" {
		var err error
		outDir, err = os.Getwd()
		if err != nil {
			return nil, err
		}
	}
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return nil, err
	}
	base := fmt.Sprintf("%s+%s", safeFilename(report.ProjectSlug), time.Now().Format("2006-01-02T15-04-05"))
	mdPath := filepath.Join(outDir, base+".md")
	if err := os.WriteFile(mdPath, []byte(ui.RenderMarkdown(report)), 0o644); err != nil {
		return nil, err
	}
	paths := []string{mdPath}
	if writeJSON {
		data, err := json.MarshalIndent(report, "", "  ")
		if err != nil {
			return nil, err
		}
		jsonPath := filepath.Join(outDir, base+".json")
		if err := os.WriteFile(jsonPath, data, 0o644); err != nil {
			return nil, err
		}
		paths = append(paths, jsonPath)
	}
	return paths, nil
}

func persistPlanPrompt(report diagnose.Report, outDir string) (string, error) {
	if strings.TrimSpace(outDir) == "" {
		var err error
		outDir, err = os.Getwd()
		if err != nil {
			return "", err
		}
	}
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return "", err
	}
	base := fmt.Sprintf("%s+%s-correction-plan-prompt", safeFilename(report.ProjectSlug), time.Now().Format("2006-01-02T15-04-05"))
	mdPath := filepath.Join(outDir, base+".md")
	if err := os.WriteFile(mdPath, []byte(ui.RenderCorrectionPlanPrompt(report)), 0o644); err != nil {
		return "", err
	}
	return mdPath, nil
}

func safeFilename(value string) string {
	value = strings.TrimSpace(strings.ToLower(value))
	if value == "" {
		return "project"
	}
	var b strings.Builder
	for _, r := range value {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' || r == '_' {
			b.WriteRune(r)
			continue
		}
		b.WriteRune('-')
	}
	return strings.Trim(b.String(), "-")
}

type exitError struct {
	code    int
	message string
}

func (e exitError) Error() string {
	return e.message
}

func (e exitError) ExitCode() int {
	return e.code
}
