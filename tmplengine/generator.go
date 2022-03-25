package tmplengine

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"stgg/cmd/printer"
	"stgg/utils"
	"stgg/yamlstgg"
)

const PathGenerated = "./"

// GenerateTemplateWithLocalVariables сгенерировать шаблон, взяв переменные из командной строки или переменных из конфига
func GenerateTemplateWithLocalVariables(templateName string, localVariables []string) error {
	var localVariableMap = make(map[interface{}]interface{})
	for i := 0; i < len(localVariables); i += 2 {
		var variableName = localVariables[0]
		var value = localVariables[1]
		localVariableMap[variableName] = value
	}
	return GenerateTemplate(templateName, localVariableMap)
}

// GenerateTemplateWithYaml Сгенерировать выбранный шаблон, взяв переменные из yaml файла
func GenerateTemplateWithYaml(templateName string, yamlPath string) error {
	var ymlVariables, err = yamlstgg.ReadYamlFile(yamlPath)
	if err != nil {
		return err
	}
	return GenerateTemplate(templateName, ymlVariables)
}

func GenerateTemplate(templateName string, variableData map[interface{}]interface{}) error {
	var globalVariables, err = yamlstgg.ReadYamlFile(GlobalVariablesPath)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	var variables = utils.MergeMaps(globalVariables, variableData)

	var templatePath = TemplatesDir + templateName
	return filepath.Walk(templatePath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			var newPathDir, err = utils.ReplaceBefore(path, "", "/")
			if err != nil {
				return errors.New("указан неккоректный путь до файла")
			}
			err = os.Mkdir(newPathDir, 0777)
			if err != nil {
				return err
			}
		}
		go func() {
			err := generateTemplateFile(path, info, err, variables)
			if err != nil {
				printer.PrintError("ошибка генерации " + path)
			}
		}()

		return nil
	})
}

func generateTemplateFile(path string, info os.FileInfo, err error, variables map[interface{}]interface{}) error {

	return nil
}
