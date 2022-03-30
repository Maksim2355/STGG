package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
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

// ReadFile Читает содержимое файла или возвращает ошибку
func ReadFile(path string) ([]byte, error) {
	var data, err = ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// RemoveContents Удаление выбранной директории и всех файлов внутри
func RemoveContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return os.Remove(dir)
}

func CountDirsAndFiles(dirPath string) (int, error) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return -1, err
	}
	return len(files), nil
}
