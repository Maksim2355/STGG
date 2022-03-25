package cmd

import (
	"github.com/spf13/cobra"
	"stgg/cmd/printer"
	"stgg/tmplengine"
)

var saveTemplateCmd = &cobra.Command{
	Use:   "saveTemplate",
	Short: "Сохранение шаблона",
	Long: `Сохранение шаблона в хранилище программы
	Принимает на вход аргумент:
	1. Краткий литерал для шаблона TEMPLATE_NAME
    2. путь до шаблона NEW_TEMPLATE_PATH
	
	Пример вызова:
	stgg saveTemplate TEMPLATE_NAME NEW_TEMPLATE_PATH
	
	В случае если аргумент не один-выкидывает ошибку`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			printer.PrintErrorAndExit("Необходимо 2 аргумента: имя шаблона и путь до папки с шаблонами")
		}
		var templateName, templatePath = args[0], args[1]
		err := tmplengine.SaveTemplate(templateName, templatePath)
		if err != nil {
			printer.PrintErrorAndExit(err.Error())
		}
		printer.PrintMessage("Шаблон успешно сохранен")
	},
}

func init() {
	rootCmd.AddCommand(saveTemplateCmd)
}
