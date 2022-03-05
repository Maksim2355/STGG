package cmd

import (
	"github.com/spf13/cobra"
	"stgg/cmd/output"
)

var saveTemplateCmd = &cobra.Command{
	Use:   "saveTemplate",
	Short: "Сохранение шаблона",
	Long: `Сохранение шаблона в хранилище программы
	Принимает на вход аргумент-путь до шаблона NEW_TEMPLATE_PATH
	
	Пример вызова:
	stgg saveTemplate NEW_TEMPLATE_PATH
	
	В случае если аргумент не один-выкидывает ошибку`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			output.PrintErrorAndExit("Необходим один единственный аргумент-путь до шаблона, который необходимо сохранить")
		}
	},
}

func init() {
	rootCmd.AddCommand(saveTemplateCmd)
}
