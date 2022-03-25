package tmplengine

import (
	"errors"
	"os"
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
	var _ = utils.MergeMaps(globalVariables, variableData)

	return errors.New("")
}
