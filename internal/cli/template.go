package cli

import (
	"github.com/spf13/cobra"
)

func newTemplateCmd(global *globalOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "template",
		Short: "Inspect the embedded project template",
		Long:  "Inspect the embedded Go AI First Kit template, including listing its files and directories.",
	}
	cmd.AddCommand(newTemplateListCmd(global))
	return cmd
}
