package cmd

import (
	"github.com/spf13/cobra"
	"stgg/cmd/printer"
	"stgg/foundation"
	"stgg/vrblstorage"
)

// allVariableCmd represents the allVariable command
var allVariableCmd = &cobra.Command{
	Use:   "allVariable",
	Short: "Показывает все доступные переменные",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			printer.PrintWarning("дополнительные аргументы не требуются")
		}
		storage := vrblstorage.VariableStorage{}
		foundation.PrintDataStorage(storage)
	},
}

func init() {
	rootCmd.AddCommand(allVariableCmd)
}
