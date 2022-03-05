package vrblstorage

import (
	"errors"
	"os"
	"stgg/cmd/printer"
)

func openFile() *os.File {
	var file *os.File

	if _, err := os.Stat(variableStorageFilename); errors.Is(err, os.ErrNotExist) {
		file, err = os.Create(variableStorageFilename)
		if err != nil {
			printer.PrintErrorAndExit(err.Error())
		}
	} else {
		file, err = os.Open(variableStorageFilename)
		if err != nil {
			printer.PrintErrorAndExit(err.Error())
		}
	}
	return file
}
