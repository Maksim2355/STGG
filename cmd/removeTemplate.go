package cmd

import (
	"stgg/cmd/printer"
	"stgg/tmplengine"

	"github.com/spf13/cobra"
)

// removeTemplateCmd represents the removeTemplate command
var removeTemplateCmd = &cobra.Command{
	Use:     "removeTemplate [TEMPLATE_NAME]",
	Short:   "Удаление шаблона",
	Long:    `Удаляет шаблон из списка ранее сохраненных`,
	Example: "stgg removeTemplate my_template",
	Args:    cobra.ExactArgs(1),
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
