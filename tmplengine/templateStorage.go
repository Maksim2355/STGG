package tmplengine

import (
	"errors"
	"io/ioutil"
	"os"
	"stgg/utils"
)

const TemplatesDir = "./templates/"

// AllTemplates Отдает список сохраненных шаблонов
func AllTemplates() ([]string, error) {
	dir, err := ioutil.ReadDir(TemplatesDir)
	if err != nil && os.IsNotExist(err) {
		if os.IsNotExist(err) {
			return make([]string, 0), nil
		}
		return nil, errors.New("ошибка при чтенции директории с шаблонами")
	}
	templatesInfo := make([]string, 0)

	for i, f := range dir {
		if f.IsDir() {
			templatesInfo[i] = f.Name()
		}
	}

	return templatesInfo, nil
}

// SaveData Сохраняет выбранный шаблон под именем key. Значение value является путем к каталогу с шаблонами
func SaveData(templateName, srcTemplatePath string) error {
	var allTemplates, err = AllTemplates()

	if err != nil {
		return err
	}

	for _, tmpl := range allTemplates {
		if tmpl == templateName {
			return errors.New("шаблон уже существует")
		}
	}

	var newTemplateDir = TemplatesDir + templateName

	err = utils.CopyDir(srcTemplatePath, newTemplateDir)
	if err != nil {
		return errors.New("ошибка при сохранении шаблона " + err.Error())
	}
	return nil
}

// RemoveTemplate RemoveByKey удаляет выбранный шаблон из хранилища. ключом является названием шаблона
func RemoveTemplate(templateName string) error {
	var allTemplates, err = AllTemplates()
	if err != nil {
		return err
	}

	var hasBeenRemoved bool

	for _, tmpl := range allTemplates {
		if tmpl == templateName {
			err = utils.RemoveContents(TemplatesDir + templateName)
			if err != nil {
				return errors.New("ошибка при удалении шаблона")
			}
			hasBeenRemoved = true
		}
	}
	if !hasBeenRemoved {
		return errors.New("шаблон не найден")
	}
	return nil
}
