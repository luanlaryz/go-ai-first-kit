package cli

import (
	"context"
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/inventa-co/go-ai-first-kit/internal/scaffold"
	"github.com/inventa-co/go-ai-first-kit/internal/templatecatalog"
	"github.com/spf13/cobra"
)

type createOptions struct {
	target         string
	params         scaffold.Params
	templateName   string
	nonInteractive bool
	force          bool
}

func newCreateCmd(global *globalOptions) *cobra.Command {
	opts := &createOptions{}
	opts.params.LicenseName = "MIT"
	opts.params.UpstreamName = "none"

	cmd := &cobra.Command{
		Use:   "create [target]",
		Short: "Create a new project from the embedded Go AI First Kit template",
		Long: "Create a new project from the embedded Go AI First Kit template.\n\n" +
			"The target argument defines where the project will be created.\n" +
			"The --module flag defines the Go module path written to go.mod.",
		Example: "  gakit create ./myapp --slug myapp --title \"My App\" --module github.com/acme/myapp\n\n" +
			"  gakit create ./services/orders \\\n" +
			"    --slug orders \\\n" +
			"    --title \"Orders Service\" \\\n" +
			"    --module github.com/acme/platform/services/orders \\\n" +
			"    --description \"Orders service\" \\\n" +
			"    --author \"Acme\"",
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 1 && opts.target == "" {
				opts.target = args[0]
			}
			if _, err := templatecatalog.Resolve(opts.templateName); err != nil {
				return err
			}
			if err := opts.collect(cmd.Context()); err != nil {
				return err
			}

			result, err := scaffold.Render(cmd.Context(), opts.target, opts.params, scaffold.Options{
				Force:        opts.force,
				InitGit:      true,
				TemplateName: opts.templateName,
			})
			if err != nil {
				return err
			}

			out := cmd.OutOrStdout()
			fmt.Fprintf(out, "Projeto AI-first criado em %s\n", result.TargetDir)
			fmt.Fprintf(out, "Arquivos renderizados: %d\n", len(result.Files))
			for _, warning := range result.Warnings {
				fmt.Fprintf(out, "Aviso: %s\n", warning)
			}
			if global.verbose {
				fmt.Fprintln(out, "Proximos comandos sugeridos:")
				fmt.Fprintln(out, "  make setup")
				fmt.Fprintln(out, "  make check-compliance")
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&opts.target, "target", "", "diretorio de destino do projeto")
	cmd.Flags().StringVar(&opts.templateName, "template", "default", "template embarcado a usar")
	cmd.Flags().StringVar(&opts.params.ProjectSlug, "slug", "", "slug Go-safe do projeto")
	cmd.Flags().StringVar(&opts.params.ProjectTitle, "title", "", "nome humano do projeto")
	cmd.Flags().StringVar(&opts.params.ModulePath, "module", "", "modulo Go, por exemplo github.com/acme/myapp")
	cmd.Flags().StringVar(&opts.params.ProjectDescription, "description", "", "descricao objetiva do projeto")
	cmd.Flags().StringVar(&opts.params.AuthorName, "author", "", "autor ou mantenedor inicial")
	cmd.Flags().StringVar(&opts.params.LicenseName, "license", "MIT", "licenca do projeto")
	cmd.Flags().StringVar(&opts.params.UpstreamName, "upstream", "none", "referencia ou upstream conceitual")
	cmd.Flags().BoolVar(&opts.nonInteractive, "non-interactive", false, "falha se algum parametro obrigatorio estiver ausente")
	cmd.Flags().BoolVar(&opts.force, "force", false, "permite renderizar em diretorio nao vazio")
	return cmd
}

func (o *createOptions) collect(ctx context.Context) error {
	o.params.ApplyDefaults()
	if o.nonInteractive {
		if o.target == "" {
			return fmt.Errorf("--target ou argumento posicional e obrigatorio em modo non-interactive")
		}
		return o.params.Validate()
	}

	if o.target == "" {
		if err := huh.NewInput().
			Title("Diretorio de destino").
			Placeholder("./myapp").
			Value(&o.target).
			Run(); err != nil {
			return err
		}
		if ctx.Err() != nil {
			return ctx.Err()
		}
	}

	inputs := []struct {
		title       string
		placeholder string
		value       *string
	}{
		{"Project slug", "myapp", &o.params.ProjectSlug},
		{"Project title", "My App", &o.params.ProjectTitle},
		{"Go module path", "github.com/acme/myapp", &o.params.ModulePath},
		{"Project description", "AI-first Go project", &o.params.ProjectDescription},
		{"Author name", "Acme Team", &o.params.AuthorName},
		{"License name", "MIT", &o.params.LicenseName},
		{"Upstream/reference name", "none", &o.params.UpstreamName},
	}
	for _, input := range inputs {
		if *input.value != "" {
			continue
		}
		if err := huh.NewInput().
			Title(input.title).
			Placeholder(input.placeholder).
			Value(input.value).
			Run(); err != nil {
			return err
		}
		if ctx.Err() != nil {
			return ctx.Err()
		}
	}

	o.params.ApplyDefaults()
	return o.params.Validate()
}
