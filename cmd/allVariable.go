package cmd

import (
	"stgg/cmd/printer"
	"stgg/vrblstorage"

	"github.com/spf13/cobra"
)

// allVariableCmd represents the allVariable command
var allVariableCmd = &cobra.Command{
	Use:   "allVariable",
	Short: "Показывает все доступные переменные",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			printer.PrintWarning("дополнительные аргументы не требуются")
		}
		data, err := vrblstorage.AllVariable()
		if err != nil {
			printer.PrintErrorAndExit(err.Error())
		}
		lenVariables := len(data)
		if lenVariables > 0 {
			printer.PrintMessage("Список всех доступных переменных:")
			for i := 0; i < len(data); i++ {
				variable := data[i][0]
				value := data[i][1]
				printer.PrintMessage(variable + "," + value)
			}
		} else {
			printer.PrintMessage("Нет доступных переменных")
		}
	},
}

func init() {
	rootCmd.AddCommand(allVariableCmd)
}
