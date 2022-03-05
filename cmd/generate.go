package cmd

import (
	"stgg/cmd/output"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Сгенерировать шаблон, который был ранее сохранен",
	Long: `Генерация шаблона из списка ранее сохраненных.
	На вход принимает:
	TEMPLATE_NAME-название шаблона
	Остальные аргументы зависят от флага

	Первый способ вызова:
	-lv: local variables в этом случае мы прописываем переменные для генерации в данной команде через двоеточие
	variableName1 valueName1 где первое значение-названием переменной, прописанной в шаблоне, а второе - значение переменной
	
	Нечетные аргументы означают название переменной, а четные ее значение.
	Если общее число аргументов нечетное-команда завершится с ошибкой
	Пример команды с использованием флага -lv:
	stgg generate TEMPLATE_NAME variable1 value1 variable2 value2 -lv

	Второй способ вызова: 
	-jsn получение переменных из json файла. В этом случае мы передаем один аргумент-путь до json файла с переменными
	Читается только первый аргумент. Остальные игнорируются.

	Пример команды с использованием -jsn
	stgg generate TEMPLATE_NAME JSON_CONFIG_PATH -jsn

	В случае если переменная содержащаяся в аргументах уже сохранена ранее с помошью командой saveVariable, то она будет переопределена
	`,
	Run: func(cmd *cobra.Command, args []string) {
		output.PrintErrorAndExit("Старт генерации")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
