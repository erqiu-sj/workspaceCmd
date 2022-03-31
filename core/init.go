package core

import (
	"fmt"
	"workspaceCmd/utils"
)

func CheckInitFile() {
	if !utils.DoesFolderExist(utils.WorkSpacePathWithLocal) {
		utils.Shell(fmt.Sprint("sudo mkdir ", utils.WorkSpacePathWithLocal), func() {
		}, func(msg string) {
			utils.RedTips(msg)
		})
	}
	if !utils.DoesFolderExist(utils.IniConfigurationFolder) {
		utils.Shell(fmt.Sprint("sudo mkdir ", utils.IniConfigurationFolder), func() {

		}, func(msg string) {
			utils.RedTips(msg)
		})
	}
	if !utils.DoesFolderExist(utils.ConfigDir) {
		utils.Shell(fmt.Sprint("sudo mkdir ", utils.ConfigDir), func() {

		}, func(msg string) {
			utils.RedTips(msg)
		})
	}
	if !utils.DoesFolderExist(utils.OpenProjectModeConfigFile) {
		utils.Shell(fmt.Sprint("sudo touch ", utils.OpenProjectModeConfigFile), func() {
		}, func(msg string) {
			utils.RedTips(msg)
		})
	}
}
