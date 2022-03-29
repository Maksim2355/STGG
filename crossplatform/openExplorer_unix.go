//go:build !windows
// +build !windows

package crossplatform

func OpenFileInExplorer(filePath string) {
	exec.Command(`explorer`, `/select,`, filePath)
}
