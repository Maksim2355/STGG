package vrblstorage

import (
	"errors"
	"stgg/utils"
)

//файл в формате csv в котором хранятся переменные
var variableStorageFilename = "./variables.csv"

// SaveVariable сохранение переменной
func SaveVariable(variableName, value string) error {
	return writeInCsv(variableStorageFilename, []string{variableName, value})
}

// AllVariable получение списка всех переменных
func AllVariable() ([][]string, error) {
	return readCsvFile(variableStorageFilename)
}

// RemoveVariable удаление переменной из списка
func RemoveVariable(variableName string) error {
	allVariables, err := AllVariable()
	if err != nil {
		return err
	}
	var indexRemovedItem = -1
	for i, item := range allVariables {
		if item[0] == variableName {
			indexRemovedItem = i
		}
	}

	if indexRemovedItem == -1 {
		return errors.New("переменной не существует")
	} else {
		variablesData := utils.RemoveFromStringMatrixWithSaveOrder(allVariables, indexRemovedItem)
		err := rewriteInCsv(variableStorageFilename, variablesData)

		if err != nil {
			return err
		}
	}

	return nil
}

// ReplaceVariable Замена выбранной переменной на новую
func ReplaceVariable(newVariableName, newVariableValue string) error {
	err := RemoveVariable(newVariableValue)
	if err != nil {
		return err
	}
	return SaveVariable(newVariableName, newVariableValue)
}

func GetValue(variable string) (string, error) {
	return findBy(variableStorageFilename, variable)
}

func FindAll(variables []string) (map[string]string, error) {
	return findByList(variableStorageFilename, variables)
}
