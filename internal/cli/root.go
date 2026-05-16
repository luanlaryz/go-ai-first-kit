package cli

import (
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
		Short: "CLI do go-ai-first-kit",
		Long: "gakit cria projetos Go AI-first a partir do template do kit e diagnostica " +
			"repositorios existentes com um report ponderado por pilar.",
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	cmd.PersistentFlags().BoolVar(&opts.noColor, "no-color", false, "desabilita cores ANSI na saida")
	cmd.PersistentFlags().BoolVar(&opts.verbose, "verbose", false, "mostra detalhes adicionais durante a execucao")

	cmd.AddCommand(newCreateCmd(opts))
	cmd.AddCommand(newDiagnoseCmd(opts))
	cmd.AddCommand(newVersionCmd())

	return cmd
}
