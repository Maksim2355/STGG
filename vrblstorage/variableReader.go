package vrblstorage

import (
	"encoding/csv"
	"errors"
	"io"
	"stgg/utils"
	"unicode/utf8"
)

//Нахождение значение переменной по ее имени
func findBy(filename, variable string) (string, error) {
	var file, err = utils.OpenFileForRead(filename)
	if err != nil {
		return "", err
	}
	reader := csv.NewReader(file)

	var value string
	for {
		record, err := reader.Read()

		if record[0] == variable {
			value = record[1]
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
	}
	if utf8.RuneCountInString(value) == 0 {
		return "", errors.New("найдена пустая строка. Возможно файл быз изменен")
	}

	return value, nil
}

func findByList(filename string, variablesList []string) (map[string]string, error) {
	var file, err = utils.OpenFileForRead(filename)
	if err != nil {
		return nil, err
	}
	countVariables := len(variablesList)
	variablesMap := make(map[string]string, countVariables)

	reader := csv.NewReader(file)

	for {
		if len(variablesMap) >= countVariables {
			break
		}

		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		currentVariable := record[0]
		currentValue := record[1]

		for i := 0; i < countVariables; i++ {
			variables := variablesList[i]
			if currentVariable == variables {
				variablesMap[currentVariable] = currentValue
				break
			}
		}
	}

	return variablesMap, nil
}
