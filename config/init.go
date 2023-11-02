package config

import (
	"fmt"
	"github.com/KM911/oslib/adt"
	"github.com/KM911/oslib/fs"
	"github.com/spf13/viper"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var (
	PWD         string
	ConfigFile  string
	ShortcutDir string

	ShortcutMap map[string][]string
)

func LoadConfig() {
	SetGlobal()
	CheckEnv()
	CheckDebug()
}

/*
CheckEnv checks if the environment is ready for the o to run
if not, it will try to fix it
*/
func SetGlobal() {
	PWD = fs.ExecutePath()
	ConfigFile = filepath.Join(PWD, "config", "config.json")
	ShortcutDir = filepath.Join(PWD, "config", "shortcut")
}
func CheckEnv() {
	viper.SetConfigFile(ConfigFile)
	err := viper.ReadInConfig()
	//viper.WatchConfig()
	if err != nil {
		fmt.Println("正在初始化环境")
		InitConfig()
		fmt.Println("初始化环境完成")
		fmt.Println("		可以使用命令: o [argv] 启动程序")
		fmt.Println("		可以修改config.json文件进行个性化配置")
		os.Exit(1)
	}
}

func Move(sourcePath, destPath string) error {
	inputFile, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("Couldn't open source file: %s", err)
	}
	outputFile, err := os.Create(destPath)
	if err != nil {
		inputFile.Close()
		return fmt.Errorf("Couldn't open dest file: %s", err)
	}
	defer outputFile.Close()
	_, err = io.Copy(outputFile, inputFile)
	inputFile.Close()
	if err != nil {
		return fmt.Errorf("Writing to output file failed: %s", err)
	}
	// The copy was successful, so now delete the original file
	err = os.Remove(sourcePath)
	if err != nil {
		return fmt.Errorf("Failed removing original file: %s", err)
	}
	return nil
}

func InitConfig() {
	os.Mkdir(filepath.Join(PWD, "config"), os.ModePerm)
	os.Mkdir(ShortcutDir, os.ModePerm)
	viper.Set("shortcut", map[string]string{})
	// 开始移动文件夹

	MoveDesktopShortcut()

	ShortcutList := fs.Dir(filepath.Join(PWD, "config", "shortcut"))
	for i := 0; i < len(ShortcutList); i++ {
		viper.Set("shortcut."+fs.FileName(ShortcutList[i]), []string{})
	}

	viper.Set("debug", false)
	viper.Set("template", "")
	//viper.Set()
	viper.WriteConfig()
}

func MoveDesktopShortcut() {
	PublicDesktop := filepath.Join("C:\\Users\\Public\\Desktop")
	PublicDesktopDir := fs.Dir(PublicDesktop)
	UserDesktop := filepath.Join("C:\\Users", os.Getenv("username"), "Desktop")
	UserDesktopDir := fs.Dir(UserDesktop)
	for _, i := range PublicDesktopDir {
		if path.Ext(i) == ".lnk" {
			println("move " + filepath.Join(PublicDesktop, i) + " to " + filepath.Join(PWD,
				"config", "shortcut", i))
			Move(filepath.Join(PublicDesktop, i), filepath.Join(PWD, "config", "shortcut", i))
		}
	}
	for _, i := range UserDesktopDir {
		if path.Ext(i) == ".lnk" {
			println("move " + filepath.Join(UserDesktop, i) + " to " + filepath.Join(PWD,
				"config", "shortcut", i))
			Move(filepath.Join(UserDesktop, i), filepath.Join(PWD, "config", "shortcut", i))
		}
	}

}

func CheckDebug() {
	if viper.GetBool("debug") {
		println("debug mode")
	}
}

func RemoveShortcutKey(key string) {
	newvalue := viper.GetStringMapStringSlice("shortcut")
	for i, v := range newvalue {
		if v == nil {
			newvalue[i] = []string{}
		}
	}
	delete(newvalue, key)
	viper.Set("shortcut", newvalue)
	viper.WriteConfig()
}

func CheckShortcut() {
	// 添加新的快捷方式
	MoveDesktopShortcut()
	shortcutList := fs.Dir(ShortcutDir)
	//fmt.Println(shortcutList)
	for i := range shortcutList {
		shortcutList[i] = strings.ToLower(fs.FileName(shortcutList[i]))
	}
	//fmt.Println(shortcutList)
	for _, i := range shortcutList {
		// 去掉后缀
		shortcut := strings.ToLower(fs.FileName(i))
		value := viper.Get("shortcut." + shortcut)
		if value == nil {
			fmt.Println("add new shortcut", shortcut)
			viper.Set("shortcut."+shortcut, []string{})
		}
	}

	// 删除不存在的快捷方式

	for key := range viper.GetStringMapStringSlice("shortcut") {
		if !adt.InArray(key, shortcutList) {
			fmt.Println("shortcut " + key + " not exist")
			RemoveShortcutKey(key)
		}
	}

	// 判断是否存在key 和 value 相同

	for key := range viper.GetStringMapStringSlice("shortcut") {
		for _, value := range viper.GetStringMapStringSlice("shortcut") {
			for _, v := range value {
				if key == v {
					fmt.Println("current key:"+key, "is repeate as value")
					break
				}
			}
		}
	}

	// 判断是否存在相同的value

	for key1, value1 := range viper.GetStringMapStringSlice("shortcut") {
		for key2, value2 := range viper.GetStringMapStringSlice("shortcut") {
			if key1 == key2 {
				continue
			}
			for _, v1 := range value1 {
				for _, v2 := range value2 {
					if v1 == v2 {
						fmt.Println("current value:" + v1 + " is repeat as value")
						break
					}
				}
			}
		}
	}
	viper.WriteConfig()
}
