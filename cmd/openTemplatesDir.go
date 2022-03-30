package cmd

import (
	"stgg/cmd/printer"
	"stgg/crossplatform"
	"stgg/res"

	"github.com/spf13/cobra"
)

var openTemplatesDirCmd = &cobra.Command{
	Use:   "openTemplatesDir",
	Short: "Открытие директори с сохраненными шаблонами",
	Run: func(cmd *cobra.Command, args []string) {
		tmplDirPath := res.GetTemplatesDirPath()
		err := crossplatform.OpenFileInExplorer(tmplDirPath)
		if err != nil {
			printer.PrintError("Ошибка открытия директории с шаблонами по пути: " + tmplDirPath + "\n" + err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(openTemplatesDirCmd)
}
