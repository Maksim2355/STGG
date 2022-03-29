package cmd

import (
	"github.com/spf13/cobra"
	"stgg/cmd/printer"
	"stgg/res"
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
			printer.PrintInfoMessage("Нет доступных шаблонов")
		} else {
			printer.PrintInfoMessage("Список шаблонов, находящихся в " + res.GetTemplatesDirPath())
			for _, tmpl := range allTemplates {
				printer.PrintMessageForList(tmpl)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(allTemplatesCmd)
}
