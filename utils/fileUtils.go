package utils

import (
	"os"
)

func OpenFileForRead(filepath string) (*os.File, error) {
	return os.OpenFile(filepath, os.O_CREATE|os.O_RDONLY, 0644)
}

func OpenFileForAppend(filepath string) (*os.File, error) {
	return os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
}

func OpenFileForWrite(filepath string) (*os.File, error) {
	return os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
}
