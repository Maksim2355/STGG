package vrblstorage

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"unicode/utf8"
)

//считывает файл с csv
func readCsvFile(filename string) ([][]string, error) {
	var file = openFile()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, file.Close()
}

//записывает данные в csvFile
func writeInCsv(filename string, data []string) error {
	if len(data) != 2 {
		return errors.New("число данных в строке должно быть равным двум")
	}
	// read the file
	lines, err := readCsvFile(filename)
	if err != nil {
		return err
	}

	// add column
	l := len(lines)
	if len(data) < l {
		l = len(data)
	}
	for i := 0; i < l; i++ {
		lines[i] = append(lines[i], data[i])
	}

	// write the file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	w := csv.NewWriter(file)
	if err = w.WriteAll(lines); err != nil {
		file.Close()
		return err
	}
	return file.Close()
}

//записывает данные в csvFile
func rewriteInCsv(filename string, data [][]string) error {
	// write the file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	w := csv.NewWriter(file)
	if err = w.WriteAll(data); err != nil {
		file.Close()
		return err
	}
	return file.Close()
}

//Нахождение значение переменной по ее имени
func findBy(variable string) (string, error) {
	var file = openFile()
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
