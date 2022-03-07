package tmplengine

import (
	"errors"
	"stgg/utils"
)

type TemplateStorage struct{}

var templatesDir = "./templates/"

var fileWithTemplatesInfo = "./templateInfo.csv"

// AllCsvData Отдает список сохраненных шаблонов и путь к ним
func (storage TemplateStorage) AllCsvData() ([][]string, error) {
	return utils.ReadCsvFile(fileWithTemplatesInfo)
}

// SaveData Сохраняет выбранный шаблон под именем key. Значение value является путем к каталогу с шаблонами
func (storage TemplateStorage) SaveData(key, value string) error {
	//TODO доработать
	return errors.New("")
}

// RemoveByKey удаляет выбранный шаблон из хранилища. ключом является названием шаблона
func (storage TemplateStorage) RemoveByKey(key string) error {
	return errors.New("")
}

// ReplaceData заменяет выбранный шаблон новым шаблном
func (storage TemplateStorage) ReplaceData(key, value string) error {
	return errors.New("")
}

// EditTemplate Открывает конфиг для выбраного шаблна
func (storage TemplateStorage) EditTemplate(templateName string) error {

	return errors.New("")
}

func (storage TemplateStorage) getConfig(templateName string) error {

	return errors.New("")
}
