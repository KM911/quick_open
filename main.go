package main

import (
	"github.com/KM911/oslib/adt"
	"github.com/spf13/viper"
	"o/command"
	"o/config"
	"os"
)

var (
	ArgsLens int
)

//  init could read the argv ? Answer is no
func init() {
	config.LoadConfig()
	config.ShortcutMap = viper.GetStringMapStringSlice("shortcut")
	ArgsLens = len(os.Args)
	println(ArgsLens)
}

func main() {
	//defer pprof.Profile(1).Stop()
	defer adt.TimerStart().End()
	if ArgsLens < 2 {
		println("need a valid command")
		// TODO help
		config.CheckShortcut()
		return
	}
	ParseUserInput()
}
func ParseUserInput() {
	for i := 1; i < ArgsLens; i++ {
		command.LogCommand(os.Args[i])
		command.Start(os.Args[i])

	}
}
