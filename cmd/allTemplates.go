package cmd

import (
	"github.com/spf13/cobra"
	"stgg/cmd/printer"
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

		var allTemplates, err = tmplengine.AllTemplates()
		if err != nil {
			printer.PrintErrorAndExit(err.Error())
		}
		if len(allTemplates) == 0 {
			printer.PrintMessage("Нет доступных шаблонов")
		} else {
			printer.PrintMessage("Список шаблонов, находящихся в " + tmplengine.TemplatesDir)
			for _, tmpl := range allTemplates {
				printer.PrintMessage(tmpl)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(allTemplatesCmd)
}
