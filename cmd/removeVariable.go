package cmd

import (
	"github.com/spf13/cobra"
	"stgg/cmd/printer"
	"stgg/vrblstorage"
)

// removeVariableCmd represents the removeVariable command
var removeVariableCmd = &cobra.Command{
	Use:   "removeVariable",
	Short: "Удаление переменных из списка ранее сохраненных",
	Long: `Удаляет переменную VARIABLE_NAME из списка ранее сохраненных.

	Пример вызова:
	stgg removeVariable VARIABLE_NAME
	
	Если такой переменной нет, то программа завершится с ошибкой
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			printer.PrintErrorAndExit("Необходим один аргумент-имя шаблона")
		} else {
			storage := vrblstorage.NewStorage()
			variableName := args[0]
			err := storage.RemoveTemplate(variableName)
			if err != nil {
				printer.PrintErrorAndExit(err.Error())
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(removeVariableCmd)
}
