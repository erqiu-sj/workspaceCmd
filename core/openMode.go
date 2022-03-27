package core

import (
	"errors"
	"workspaceCmd/utils"
)

type OpenMode struct {
	Label    string //  显示label
	CmdAlias string //  打开项目的电脑全局变量别名
}

func (receiver *OpenMode) NewConfig() {
	defer func() {
		utils.GreenTips("successfully added open mode")
	}()
	iniHandler := utils.IniHelper{
		Path: utils.OpenProjectModeConfigFile,
	}
	iniHandler.NewIni().
		NewSection(receiver.Label).
		NewKey(receiver.Label, utils.ModeCmd, receiver.CmdAlias).
		NewKey(receiver.Label, utils.Label, receiver.Label).
		Save()
}

func (receiver OpenMode) GetItem() (OpenMode, error) {
	var config OpenMode
	for _, item := range receiver.GetOpenModeConfig() {
		if item.Label == receiver.Label {
			config = item
			return config, nil
		}
	}
	return config, errors.New("no corresponding configuration found")
}
func (receiver OpenMode) GetOpenModeConfig() []OpenMode {
	iniFile := utils.IniHelper{
		Path: utils.OpenProjectModeConfigFile,
	}
	handler := iniFile.NewIni()
	allSections := handler.GetSectionToStrings()
	var list []OpenMode
	for _, item := range allSections {
		list = append(list,
			OpenMode{
				Label:    handler.GetKey(item, utils.Label),
				CmdAlias: handler.GetKey(item, utils.ModeCmd),
			},
		)
	}
	return list
}
