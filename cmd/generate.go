package cmd

import (
	"github.com/spf13/cobra"
	"stgg/cmd/printer"
	"stgg/tmplengine"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate [TEMPLATE_NAME] [ARGS..]",
	Short: "Сгенерировать шаблон, который был ранее сохранен",
	Long: `Генерация шаблона из списка ранее сохраненных.
	На вход принимает:
	TEMPLATE_NAME-название шаблона
	Остальные аргументы зависят от флага

	По умолчанию принимает список с переменными. Т.е массив с ключом-значением. Внимание, число аргументов должно быть
	четное, каждому ключу должно соответствовать значение

	При добавление флага --path или -p список переменных для генерации берется из yaml файла, к которому мы предоставим
	путь т.е вторым аргументом должен быть путь до yaml-файла c вашими переменными
	
	Переменные, переданные с помощью данной команды имеют выше приоритет чем глобальные переменные
	В случае если есть глобальная переменная с таким имене, то будет переопределена
	`,
	Example: "stgg generate my_template variable1 value1 variable2 value2\n" +
		"stgg generate my_template C:/yourPath/variables.yaml --path\n",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			printer.PrintErrorAndExit("Не передано название шаблона для генерации")
		}
		printer.PrintSuccessMessage("Старт генерации")

		shouldGetVariablesFromPath, _ := cmd.Flags().GetBool("path")

		var templateName = args[0]
		if !shouldGetVariablesFromPath {
			printer.PrintInfoMessage("Генерация шаблона с локальными переменными")
			var localVariables = args[1:]
			if len(localVariables)%2 != 0 {
				printer.PrintErrorAndExit("каждой переменной необходимо проставить значение")
			}
			err := tmplengine.GenerateTemplateWithLocalVariables(templateName, localVariables)
			if err != nil {
				printer.PrintErrorAndExit(err.Error())
			}
		} else {
			var ymlPath = args[1]
			printer.PrintInfoMessage("Генерация шаблона с переменными из конфиг-файла: " + ymlPath)
			err := tmplengine.GenerateTemplateWithYaml(templateName, ymlPath)
			if err != nil {
				printer.PrintErrorAndExit(err.Error())
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().BoolP("path", "p", false, "Взять переменные с конфиг файла yaml")
}
