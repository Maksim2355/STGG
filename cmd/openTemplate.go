package cmd

import (
	"github.com/spf13/cobra"
	"stgg/cmd/printer"
	"stgg/crossplatform"
	"stgg/res"
	"stgg/tmplengine"
)

var openTemplateCmd = &cobra.Command{
	Use:   "openTemplate [TEMPLATE_NAME]",
	Short: "Открытие директории с ранее сохраненным шаблоном",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var templateForOpened = args[0]
		var allTmpl, err = tmplengine.AllTemplates()
		if err != nil {
			printer.PrintWarning("Внимание. Нельзя проверить присутствие шаблона в списке всех шаблонов")
			openTemplate(templateForOpened)
		} else {
			var hasTemplateInStorage bool
			for _, tmplInStorage := range allTmpl {
				if templateForOpened == tmplInStorage {
					hasTemplateInStorage = true
					break
				}
			}
			if hasTemplateInStorage {
				openTemplate(templateForOpened)
			} else {
				printer.PrintErrorAndExit("Шаблон " + templateForOpened + " не найден")
			}
		}
	},
}

func openTemplate(templateName string) {
	var templatePath = res.GetTemplatesDirPath() + crossplatform.PATH_SEPARATOR + templateName
	err := crossplatform.OpenFileInExplorer(templatePath)
	if err != nil {
		printer.PrintErrorAndExit("Ошибка при открытии шаблона по пути " + templatePath + "\n" + err.Error())
	}
}

func init() {
	rootCmd.AddCommand(openTemplateCmd)
}
