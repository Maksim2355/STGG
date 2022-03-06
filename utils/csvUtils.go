package utils

import (
	"encoding/csv"
	"errors"
	"stgg/cmd/printer"
)

// ReadCsvFile Считывает файл с csv
func ReadCsvFile(filename string) ([][]string, error) {
	var file, err = OpenFileForRead(filename)
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
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

	lines = append(lines, data)

	for i := 0; i < len(lines); i++ {
		if lines[i][0] == data[0] {
			return errors.New("значение уже существует")
		}
	}

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
