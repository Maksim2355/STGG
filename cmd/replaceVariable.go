package cmd

import (
	"github.com/spf13/cobra"
	"stgg/cmd/printer"
	"stgg/vrblstorage"
)

var replaceVariableCmd = &cobra.Command{
	Use:   "replaceVariable",
	Short: "Заменяет переменную из ранее сохраненных",
	Long: `Замена переменных из списка сохраненных переменных проекта
	Первый аргумент-название переменной, которую мы хотим заменить VARIABLE_NAME
	Второй аргумент-новое значение для переменной NEW_VALUE
	
	Пример вызова: stgg replaceVariable VARIABLE_NAME NEW_VALUE
	
	Если число аргумента больше или меньше двух или переменной нет в списке-выкидывается ошибка`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			printer.PrintErrorAndExit("Необходимо два аргумента. Название переменной и новое значение")
		} else {
			variableName := args[0]
			newValue := args[1]
			err := vrblstorage.ReplaceVariable(variableName, newValue)
			if err != nil {
				printer.PrintErrorAndExit(err.Error())
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(replaceVariableCmd)
}
