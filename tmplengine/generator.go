package tmplengine

import "errors"

// GenerateTemplateWithLocalVariables сгенерировать шаблон, взяв переменные из командной строки
func GenerateTemplateWithLocalVariables(storage TemplateStorage, data *LocalVariables) error {
	//TODO генерация шаблона
	return errors.New("")
}

// GenerateTemplateWithJson Сгенерировать выбранный шаблон, взяв переменные из json
func GenerateTemplateWithJson(storage TemplateStorage, jsonPath string) error {
	//TODO генерация шаблона
	return errors.New("")
}
