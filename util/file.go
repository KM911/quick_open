package util

import (
	"os"
	"path/filepath"
)

func ReadAll(src string) string {
	data, err := os.ReadFile(src)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func CreatFile(src string, data string) {
	err := os.WriteFile(src, []byte(data), 0666)
	if err != nil {
		panic(err)
	}
}

func DeleteFile(src string) {
	err := os.Remove(src)
	if err != nil {
		panic(err)
	}
}

func CopyFile(src string, dst string) {
	data := ReadAll(src)
	CreatFile(dst, data)
}

func FileName(file_ string) string {
	ext := filepath.Ext(file_)
	file_ = filepath.Base(file_)
	length := len(file_)
	return file_[:length-len(ext)]
}
