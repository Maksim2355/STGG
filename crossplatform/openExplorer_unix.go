//go:build !windows
// +build !windows

package crossplatform

import "github.com/skratchdot/open-golang/open"

func OpenFileInExplorer(filePath string) error {
	return open.Run(filePath)
}
