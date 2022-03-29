package cmd

import (
	"stgg/cmd/printer"
	"stgg/tmplengine"

	"github.com/spf13/cobra"
)

// removeTemplateCmd represents the removeTemplate command
var removeTemplateCmd = &cobra.Command{
	Use:   "removeTemplate",
	Short: "Удаление шаблона",
	Long: `Удаляет шаблон из списка ранее сохраненных

	Принимает аргумент- TEMPLATE_NAME имя ранее сохраненого шаблона
	
	Пример вызова:
	stgg removeTemplate TEMPLATE_NAME
	
	В случае если не указали имя шаблона-выкидывает ошибку`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			printer.PrintErrorAndExit("Необходим один аргумент-название ранее сохраненного шаблона")
		}
		var tmplName = args[0]
		err := tmplengine.RemoveTemplate(tmplName)
		if err != nil {
			printer.PrintErrorAndExit(err.Error())
		}
		printer.PrintSuccessMessage("Шаблон " + tmplName + " успешно удален")
	},
}

func init() {
	rootCmd.AddCommand(removeTemplateCmd)
}
