package utils

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"stgg/cmd/printer"
	"unicode/utf8"
)

// ReadCsvFile Считывает файл с csv
func ReadCsvFile(filename string) ([][]string, error) {
	var file, err = OpenFileForRead(filename)
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, errors.New("Ошибка при чтении csv файла: " + err.Error())
	}

	return records, file.Close()
}

// WriteInCsv Добавляет новые данные в файл
//где filename- путь к файлу
//data массив с двумя элементами, где первый элемент ключ, а второй-значение
func WriteInCsv(filename string, data []string) error {
	if len(data) != 2 {
		return errors.New("число данных в строке должно быть равным двум")
	}
	lines, err := ReadCsvFile(filename)
	if err != nil {
		return err
	}

	for i := 0; i < len(lines); i++ {
		fmt.Println(lines[i])
		if lines[i][0] == data[0] {
			return errors.New("значение уже существует")
		}
	}

	lines = append(lines, data)

	// write the file
	file, err := OpenFileForAppend(filename)

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

// RewriteInCsv Перезапизывает данные в файл
func RewriteInCsv(filename string, data [][]string) error {
	file, err := OpenFileForWrite(filename)
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

// FindByKey Получение значения по определенному ключу
func FindByKey(filename, key string) (string, error) {
	var file, err = OpenFileForRead(filename)
	if err != nil {
		return "", err
	}
	reader := csv.NewReader(file)

	var value string
	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}

		if len(record) != 2 {
			return "", errors.New("ошибка чтения данных из конфиг файла. Возможно данные были изменены")
		}

		if record[0] == key {
			value = record[1]
		}

	}
	if utf8.RuneCountInString(value) == 0 {
		return "", errors.New("найдена пустая строка. Возможно файл быз изменен")
	}

	return value, nil
}

// FindByListKeys Получения карты ключей и значений из файла [filename]
func FindByListKeys(filename string, keysList []string) (map[string]string, error) {
	var file, err = OpenFileForRead(filename)
	if err != nil {
		return nil, err
	}
	countVariables := len(keysList)
	data := make(map[string]string, countVariables)

	reader := csv.NewReader(file)

	for {
		if len(data) >= countVariables {
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
			variables := keysList[i]
			if currentVariable == variables {
				data[currentVariable] = currentValue
				break
			}
		}
	}

	return data, nil
}
