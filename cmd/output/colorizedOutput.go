package output

import (
	"fmt"
	"os"
)

var colorError = "\033[31m"
var colorWarning = "\033[33m"
var colorReset = "\033[0m"

func PrintWarning(warningText string) {
	fmt.Println(colorWarning + warningText + colorReset)
}

func PrintError(errorText string) {
	fmt.Println(string(colorError) + errorText + colorReset)
}

func PrintErrorAndExit(errorText string) {
	PrintError(errorText)
	os.Exit(-1)
}

func PrintMessage(text string) {
	fmt.Println(text)
}
