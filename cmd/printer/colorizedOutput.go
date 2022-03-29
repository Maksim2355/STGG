package printer

import (
	"fmt"
	"os"
)

var colorError = "\033[31m"
var colorWarning = "\033[33m"
var colorSuccess = "\033[92m"
var colorForListItems = "\033[96m"
var colorInfo = "\033[1m"
var colorReset = "\033[0m"

func PrintInfoMessage(message string) {
	fmt.Println(string(colorInfo) + message + colorReset)
}

func PrintSuccessMessage(message string) {
	fmt.Println(string(colorSuccess) + message + colorReset)
}

func PrintMessageForList(message string) {
	fmt.Println(string(colorForListItems) + message + colorReset)
}

func PrintWarning(warningText string) {
	fmt.Println(string(colorWarning) + warningText + colorReset)
}

func PrintError(errorText string) {
	fmt.Println(string(colorError) + errorText + colorReset)
}

func PrintErrorAndExit(errorText string) {
	PrintError(errorText)
	os.Exit(-1)
}
