package cmd

import (
	"github.com/spf13/cobra"
	"stgg/cmd/printer"
	"stgg/crossplatform"
)

// editStggConfig represents the editTemplate command
var editStggConfig = &cobra.Command{
	Use:   "editStggConfig",
	Short: "Открытие файла с конфигурацией приложения",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			printer.PrintWarning("Аргументы не требуются см документацию")
		}
		err := crossplatform.OpenFileInExplorer(СfgFile)
		if err != nil {
			printer.PrintErrorAndExit("Ошибка открытия конфига по пути " + СfgFile)
		}
	},
}

func init() {
	rootCmd.AddCommand(editStggConfig)
}
