package yamlstgg

import (
	"errors"
	"gopkg.in/yaml.v2"
	"stgg/utils"
)

// ReadYamlFile читает yaml файл и возвращает данные прочитанные из него
func ReadYamlFile(filePath string) (map[interface{}]interface{}, error) {
	var fileData, err = utils.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var data = make(map[interface{}]interface{})
	err = yaml.Unmarshal(fileData, data)
	if err != nil {
		return nil, errors.New("ошибка чтения yaml, возможно файл " + filePath + " невалидный")
	}

	return data, nil
}
