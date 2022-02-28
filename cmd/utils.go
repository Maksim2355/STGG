package cmd

import (
	"fmt"
	"os"
)

//Завершение программы и вывод ошибки в консоль
func printErrorAndStopProgramm(errorText string) {
	fmt.Errorf(errorText)
	os.Exit(-1)
}
