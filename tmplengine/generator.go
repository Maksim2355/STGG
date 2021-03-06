package tmplengine

import (
	"io/fs"
	"os"
	"path/filepath"
	"stgg/cmd/printer"
	"stgg/crossplatform"
	"stgg/res"
	"stgg/utils"
	"stgg/yamlstgg"
	"strings"
	"text/template"
)

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
	var globalVariables, err = yamlstgg.ReadYamlFile(res.GetGlobalVariablesPath())
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	var variables map[interface{}]interface{}
	if os.IsNotExist(err) {
		variables = variableData
	} else {
		variables = utils.MergeMaps(globalVariables, variableData)
	}

	var templatesDirPath = res.GetTemplatesDirPath()
	var templatePath = templatesDirPath + crossplatform.PATH_SEPARATOR + templateName

	err = filepath.Walk(templatePath, func(path string, info fs.FileInfo, err error) error {
		if path == templatePath {
			return nil
		}
		if err != nil {
			return err
		}
		if info.IsDir() {
			var newPathDir = strings.Replace(
				path,
				templatesDirPath+crossplatform.PATH_SEPARATOR+templateName,
				res.GeneratedFilesDirPath,
				-1,
			)
			err = os.Mkdir(newPathDir, 0777)
			if err != nil && !os.IsNotExist(err) {
				printer.PrintInfoMessage(err.Error())
			}
		} else {
			newPathFile := strings.Replace(
				path,
				res.GetTemplatesDirPath()+crossplatform.PATH_SEPARATOR+templateName,
				res.GeneratedFilesDirPath,
				-1,
			)
			printer.PrintInfoMessage("Генерация нового файла с путем " + newPathFile)
			err = generateTemplateFile(path, newPathFile, variables)
			if err != nil {
				printer.PrintError("ошибка генерации " + err.Error())
			}
		}
		return nil
	})
	return err
}

func generateTemplateFile(path, newPath string, variables map[interface{}]interface{}) error {
	fileBytes, err := utils.ReadFile(path)
	if err != nil {
		return err
	}
	tmpl, err := template.New(path).Parse(string(fileBytes))
	if err != nil {
		return err
	}
	fileForWrite, err := utils.OpenFileForWrite(newPath)
	if err != nil {
		return err
	}
	return tmpl.Execute(fileForWrite, variables)
}
