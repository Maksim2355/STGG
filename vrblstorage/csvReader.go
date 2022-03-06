package vrblstorage

import (
	"encoding/csv"
	"errors"
	"io"
	"stgg/cmd/printer"
	"stgg/utils"
	"unicode/utf8"
)

//Считывает файл с csv
func readCsvFile(filename string) ([][]string, error) {
	var file, err = utils.OpenFileForRead(filename)
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, file.Close()
}

//Добавляет новые данные в файл
func writeInCsv(filename string, data []string) error {
	if len(data) != 2 {
		return errors.New("число данных в строке должно быть равным двум")
	}
	lines, err := readCsvFile(filename)
	if err != nil {
		return err
	}

	lines = append(lines, data)

	for i := 0; i < len(lines); i++ {
		if lines[i][0] == data[0] {
			return errors.New("Переменная уже существует")
		}
	}

	// write the file
	file, err := utils.OpenFileForAppend(filename)

	w := csv.NewWriter(file)

	if err = w.WriteAll(lines); err != nil {
		file.Close()
		return err
	}
	w.Flush()

	if err := w.Error(); err != nil {
		printer.PrintError(err.Error())
	}
	return file.Close()
}

//Перезапизывает данные в файл
func rewriteInCsv(filename string, data [][]string) error {
	file, err := utils.OpenFileForWrite(filename)
	if err != nil {
		return err
	}
	w := csv.NewWriter(file)
	if err := w.WriteAll(data); err != nil {
		file.Close()
		return err
	}
	w.Flush()
	return file.Close()
}

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

//Нахождение значений переменных по ее имени
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
