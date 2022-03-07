package cmd

import (
	"github.com/spf13/cobra"
	"stgg/cmd/printer"
	"stgg/foundation"
	"stgg/tmplengine"
)

// allTemplatesCmd represents the allTemplates command
var allTemplatesCmd = &cobra.Command{
	Use:   "allTemplates",
	Short: "Показать все доступные шаблоны",
	Long:  `Показывает все доступные шаблоны ранее сохраненные командой saveTemplate`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			printer.PrintWarning("дополнительные аргументы не требуются")
		}
		storage := tmplengine.NewStorage()
		foundation.PrintDataStorage(storage)
	},
}

func init() {
	rootCmd.AddCommand(allTemplatesCmd)
}
