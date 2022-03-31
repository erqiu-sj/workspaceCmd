package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// DoesFolderExist 文件夹或文件不存在
func DoesFolderExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
func Mkdir(path string) {
	err := os.MkdirAll(path, 0766)
	ColdKiller(err)
}

func RemoveFile(file string) {
	Shell(fmt.Sprint("sudo rm -rf ", file), func() {
		GreenTips("successfully deleted")
	}, func(msg string) {
		RedTips(msg)
	})
	//err := os.Remove(file)
	//ColdKiller(err)
}

func GetPwd() string {
	dir, _ := os.Getwd()
	return dir
}

func JoinPwd(path string) string {
	return filepath.Join(GetPwd(), path)
}

func CreateFile(path string) {
	Shell(fmt.Sprint("sudo touch ", path), func() {
	}, func(msg string) {
		RedTips(CreateErrorFile)
	})
	//f, createErr := os.Create(path)
	//defer func(f *os.File) {
	//	_ = f.Close()
	//}(f)
	//InterceptErrorsAndKillProcessImmediately(createErr, func(msg string) {
	//	RedTips(CreateErrorFile)
	//})
}

func IsDir(path string) bool {
	f, err := os.Stat(JoinPwd(path))
	ColdKiller(err)
	return f.IsDir()

}

func CreateIniFile(name string) string {
	path := filepath.Join(IniConfigurationFolder, fmt.Sprint(name, ".ini"))
	CreateFile(path)
	return path
}

// GetLastFileNameDirectoryNamePath 获取路径最后的文件夹或文件 /src/pack  => pack
func GetLastFileNameDirectoryNamePath(path string) string {
	if path == "/" {
		return GetLastFileNameDirectoryNamePath(GetPwd())
	}
	sliceList := strings.FieldsFunc(path, func(r rune) bool {
		if r == '/' {
			return true
		}
		return false
	})
	return sliceList[len(sliceList)-1]
}
