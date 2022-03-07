package cmd

import (
	"stgg/cmd/printer"
	"stgg/tmplengine"

	"github.com/spf13/cobra"
)

// editTemplateCmd represents the editTemplate command
var editTemplateCmd = &cobra.Command{
	Use:   "editTemplate",
	Short: "Редактирования конфигурации шаблонов",
	Long: `Открывает файл конфигурации шаблона. Если его нет, то создает его и позже настройки применятся к шаблону
	На вход принимает один аргумент TEMPLATE_NAME-название шаблона, который мы хотим редактировать

	Пример вызова:
	stgg editTemplate TEMPLATE_NAME
	
	В случае если количество аргументов больше или меньше одного, то программа завершается с ошибкой`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			printer.PrintErrorAndExit("Аргумент должен быть один-имя шаблона")
		}
		var storage = tmplengine.NewStorage()
		var templateName = args[0]
		err := storage.EditTemplate(templateName)
		if err != nil {
			printer.PrintErrorAndExit(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(editTemplateCmd)
}
