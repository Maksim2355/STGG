package vrblstorage

import (
	"errors"
	"stgg/utils"
)

type VariableStorage struct{}

//файл в формате csv в котором хранятся переменные
var variableStorageFilename = "./variables.csv"

// SaveData сохранение переменной
func (storage VariableStorage) SaveData(key, value string) error {
	return utils.WriteInCsv(variableStorageFilename, []string{key, value})
}

// AllCsvData получение списка всех переменных
func (storage VariableStorage) AllCsvData() ([][]string, error) {
	return utils.ReadCsvFile(variableStorageFilename)
}

// RemoveByKey удаление переменной из списка
func (storage VariableStorage) RemoveByKey(key string) error {
	allVariables, err := storage.AllCsvData()
	if err != nil {
		return err
	}
	var indexRemovedItem = -1
	for i, item := range allVariables {
		if item[0] == key {
			indexRemovedItem = i
		}
	}

	if indexRemovedItem == -1 {
		return errors.New("переменной не существует")
	} else {
		variablesData := utils.RemoveFromStringMatrixWithSaveOrder(allVariables, indexRemovedItem)
		err := utils.RewriteInCsv(variableStorageFilename, variablesData)

		if err != nil {
			return err
		}
	}

	return nil
}

// ReplaceData Замена выбранной переменной на новую
func (storage VariableStorage) ReplaceData(key, value string) error {
	err := storage.RemoveByKey(key)
	if err != nil {
		return err
	}
	return storage.SaveData(key, value)
}

func GetValue(variable string) (string, error) {
	return findBy(variableStorageFilename, variable)
}

func FindAll(variables []string) (map[string]string, error) {
	return findByList(variableStorageFilename, variables)
}

func NewStorage() *VariableStorage {
	return &VariableStorage{}
}
