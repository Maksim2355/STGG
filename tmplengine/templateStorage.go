package tmplengine

import (
	"errors"
	"stgg/utils"
)

type TemplateStorage struct{}

var templatesDir = "./templates/"

var fileWithTemplatesInfo = "./templateInfo.csv"

func (storage TemplateStorage) AllCsvData() ([][]string, error) {
	return utils.ReadCsvFile(fileWithTemplatesInfo)
}

func (storage TemplateStorage) SaveData(key, value string) error {
	//TODO доработать
	return errors.New("")
}

func (storage TemplateStorage) RemoveByKey(key string) error {
	return errors.New("")
}

func (storage TemplateStorage) ReplaceData(key, value string) error {
	return errors.New("")
}
