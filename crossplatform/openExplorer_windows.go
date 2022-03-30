//go:build windows
// +build windows

package crossplatform

import (
	"os/exec"
)

func OpenFileInExplorer(filePath string) {
	exec.Command(`explorer`, `/select,`, filePath)
}
