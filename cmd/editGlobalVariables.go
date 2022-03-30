package cmd

import (
	"stgg/cmd/printer"
	"stgg/crossplatform"
	"stgg/res"

	"github.com/spf13/cobra"
)

var editGlobalVariablesCmd = &cobra.Command{
	Use:   "editGlobalVariables",
	Short: "Открытия файла для редактирования глобальных переменных",
	Run: func(cmd *cobra.Command, args []string) {
		globalVariablesPath := res.GetGlobalVariablesPath()
		err := crossplatform.OpenFileInExplorer(globalVariablesPath)
		if err != nil {
			printer.PrintError("Ошибка открытия файла с глобальными переменными по пути " + globalVariablesPath)
		}
	},
}

func init() {
	rootCmd.AddCommand(editGlobalVariablesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editGlobalVariablesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editGlobalVariablesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
