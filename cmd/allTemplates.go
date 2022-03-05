package cmd

import (
	"stgg/cmd/printer"

	"github.com/spf13/cobra"
)

// allTemplatesCmd represents the allTemplates command
var allTemplatesCmd = &cobra.Command{
	Use:   "allTemplates",
	Short: "Показать все доступные шаблоны",
	Long:  `Показывает все доступные шаблоны ранее сохраненные командой saveTemplate`,
	Run: func(cmd *cobra.Command, args []string) {
		printer.PrintErrorAndExit("Доступные шаблоны для генерации файлов:")
	},
}

func init() {
	rootCmd.AddCommand(allTemplatesCmd)
}
