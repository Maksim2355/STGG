package cmd

import (
	"stgg/cmd/printer"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "STGG",
	Short: "CLI для простой шаблонизации",
	Long: `Шаблонизатор с использованием стандартного языка шаблонов golang
	
	Более подробную документацию можно найти в репозитории проекта https://github.com/Maksim2355/STGG`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		printer.PrintErrorAndExit(err.Error())
	}
}

func init() {
	//nothing
}
