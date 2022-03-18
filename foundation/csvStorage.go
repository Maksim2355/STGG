package foundation

import (
	"stgg/cmd/printer"
)

type BaseCsvStorage interface {
	AllCsvData() ([][]string, error)
	RemoveTemplate(key string) error
	SaveData(key, value string) error
	ReplaceData(key, value string) error
}

func PrintDataStorage(storage BaseCsvStorage) {
	data, err := storage.AllCsvData()
	if err != nil {
		printer.PrintErrorAndExit(err.Error())
	}
	lenVariables := len(data)
	if lenVariables > 0 {
		printer.PrintMessage("Список достпных данных")
		for i := 0; i < len(data); i++ {
			variable := data[i][0]
			value := data[i][1]
			printer.PrintMessage(variable + "," + value)
		}
	} else {
		printer.PrintError("Нет данных")
	}
}
