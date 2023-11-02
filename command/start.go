package command

import (
	"github.com/KM911/oslib"
	"o/config"
	"path/filepath"
)

func Start(cmd string) {
	target := ""
	for key, value := range config.ShortcutMap {
		if key == cmd {
			target = key
		}
		for _, v := range value {
			if v == cmd {
				target = key
			}
		}
	}
	if target == "" {
		println("command not found")
		//	 TODO help user to find the right command
		//  I need to parse the command
		guessTheCommand(cmd)
	} else {
		oslib.Run("start " + filepath.Join(config.ShortcutDir, target))
	}
}

func guessTheCommand(cmd string) {
	// TODO
	if len(cmd) >= 2 {
		// full 	name search from the map

	} else {
		// short name
		// beacuse the short name will have too many matched words
		// so we need to find it in the history

		// load history

		// parse the history

		// find the most matched command
	}
}
