package command

import (
	"o/config"
	"os"
	"path/filepath"
	"strings"
)

func LogCommand(cmd string) {
	file, err := os.Open("history.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString(cmd)
	file.WriteString("\n")
}

// 为什么我们需要历史记录？
// becaus the cmd is too short and will match too many commands
// but if we have a history, we can find the most matched command

// cmd will be look like only o or two letters
func SearchHistory(cmd string) {
	// TODO
	file, err := os.Open(filepath.Join(config.PWD, "history.txt"))
	if err != nil {

	}
	defer file.Close()
	fileInfo, erro := file.Stat()
	if erro != nil {

	}
	cache := make([]byte, fileInfo.Size())
	file.Read(cache)
	str := string(cache)
	str_slice := strings.Split(str, "\n")
	match_silce := make([]byte, len(str_slice))
	for _, char := range cmd {
		for index, v := range str_slice {
			if strings.Contains(v, string(char)) {
				match_silce[index]++
			}
		}
	}
}
