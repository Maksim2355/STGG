package res

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"stgg/cmd/printer"
	"stgg/crossplatform"
)

const templatesDirConfigKey = "templatesPath"
const templateDirName = "templates"

const globalVariablesDirConfigKey = "globalVariables"
const globalVariablesFileName = "globalVariables.yaml"

const GeneratedFilesDirPath = "."

// GetGlobalVariablesPath Получение пути до файла с глобальными переменными
func GetGlobalVariablesPath() string {
	globalVariablesPath := viper.GetString(globalVariablesDirConfigKey)
	if len(globalVariablesPath) == 0 {
		return HomeDir() + crossplatform.PATH_SEPARATOR + globalVariablesFileName
	}
	return globalVariablesPath
}

// GetTemplatesDirPath Возвращает путь до директории, где должны сохраняться примеры шаблонов
func GetTemplatesDirPath() string {
	templatesDirPath := viper.GetString(templatesDirConfigKey)
	if len(templatesDirPath) == 0 {
		return HomeDir() + crossplatform.PATH_SEPARATOR + templateDirName
	}
	return templatesDirPath
}

func HomeDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		printer.PrintError("ошибка получения домашней директории")
		return "."
	}
	return home
}

// PrepareDefaultConfig иннициализация дефолтного конфига
func PrepareDefaultConfig() string {
	homeDirWithSeparated := HomeDir() + crossplatform.PATH_SEPARATOR
	return fmt.Sprintf(
		"%s: %s \n%s: %s",
		templatesDirConfigKey,
		homeDirWithSeparated+templateDirName,
		globalVariablesDirConfigKey,
		homeDirWithSeparated+globalVariablesFileName,
	)
}
