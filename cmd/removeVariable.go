package cmd

import (
	"github.com/spf13/cobra"
	"stgg/cmd/output"
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
			output.PrintErrorAndExit("Необходим один аргумент-имя шаблона")
		}
		//TODO удаление шаблона
	},
}

func init() {
	rootCmd.AddCommand(removeVariableCmd)
}
