package cmd

import (
	"github.com/spf13/cobra"
	"stgg/cmd/printer"
	"stgg/tmplengine"
)

var saveTemplateCmd = &cobra.Command{
	Use:   "saveTemplate [TEMPLATE_NAME] [TEMPLATE_PATH]",
	Short: "Сохранение шаблона",
	Long: `Сохранение шаблона в хранилище программы
	Принимает на вход аргумент:
	1. Краткий литерал для шаблона TEMPLATE_NAME
    2. путь до папки с шаблоном TEMPLATE_PATH`,
	Args:    cobra.ExactArgs(2),
	Example: "stgg saveTemplate my_template C:/yourTemplateDir",
	Run: func(cmd *cobra.Command, args []string) {
		var templateName, templatePath = args[0], args[1]
		err := tmplengine.SaveTemplate(templateName, templatePath)
		if err != nil {
			printer.PrintErrorAndExit(err.Error())
		}
		printer.PrintSuccessMessage("Шаблон успешно сохранен")
	},
}

func init() {
	rootCmd.AddCommand(saveTemplateCmd)
}
