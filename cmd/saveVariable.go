package cmd

import (
	"github.com/spf13/cobra"
	"stgg/cmd/output"
)

var saveVariableCmd = &cobra.Command{
	Use:   "saveVariable",
	Short: "Сохранение переменной для последующей подстановки в шаблоны",
	Long: `Следует сохранять переменные, которые не меняются от шаблона шаблону, а являются одинаковыми, такие как packgae_name и тд
	Принимает на вход два аргумент:
	VARIABLE_NAME-название переменной, которая будет определяться в шаблоне
	VALUE-значение переменной

	Пример вызова:
	stgg saveVariable VARIABLE_NAME VALUE

	В случае если число аргументов не равно двум-выкидываестся ошибка
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			output.PrintErrorAndExit("Число аргументов должно быть равным двум")
		}
	},
}

func init() {
	rootCmd.AddCommand(saveVariableCmd)
}
