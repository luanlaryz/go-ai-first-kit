package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

type globalOptions struct {
	noColor bool
	verbose bool
}

func Execute() error {
	return newRootCmd().Execute()
}

func newRootCmd() *cobra.Command {
	opts := &globalOptions{}

	cmd := &cobra.Command{
		Use:   "gakit",
		Short: "Create AI-first Go applications from a preconfigured template",
		Long: "Go AI First Kit is a CLI for creating AI-first Go applications from a preconfigured template.\n\n" +
			"It can scaffold a project with repository conventions, docs, automation, skills,\n" +
			"examples, scripts, and agent-oriented configuration already in place.",
		Example: "  gakit create ./myapp --slug myapp --title \"My App\" --module github.com/acme/myapp\n\n" +
			"  gakit create ./myapp \\\n" +
			"    --slug myapp \\\n" +
			"    --title \"My App\" \\\n" +
			"    --module github.com/acme/myapp \\\n" +
			"    --description \"My AI-first Go app\" \\\n" +
			"    --author \"Acme\"\n\n" +
			"  gakit template list\n" +
			"  gakit template list --tree\n" +
			"  gakit template list --json",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, _ []string) error {
			out := cmd.OutOrStdout()
			fmt.Fprintln(out, "Go AI First Kit")
			fmt.Fprintln(out)
			fmt.Fprintln(out, "Tip: create a new project with:")
			fmt.Fprintln(out)
			fmt.Fprintln(out, "  gakit create ./myapp --slug myapp --title \"My App\" --module github.com/acme/myapp")
			fmt.Fprintln(out)
			fmt.Fprintln(out, "Run \"gakit help\" to see all commands.")
			return nil
		},
	}

	cmd.PersistentFlags().BoolVar(&opts.noColor, "no-color", false, "desabilita cores ANSI na saida")
	cmd.PersistentFlags().BoolVar(&opts.verbose, "verbose", false, "mostra detalhes adicionais durante a execucao")

	cmd.AddCommand(newCreateCmd(opts))
	cmd.AddCommand(newTemplateCmd(opts))
	cmd.AddCommand(newDiagnoseCmd(opts))
	cmd.AddCommand(newVersionCmd())

	return cmd
}
