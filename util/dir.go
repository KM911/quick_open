package util

import "os"

func Dir(path string) (subfiles []string) {
	files, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		// 我们不希望就是显示文件夹 而是只显示文件
		if !file.IsDir() {
			subfiles = append(subfiles, file.Name())
		}
	}
	return
}
