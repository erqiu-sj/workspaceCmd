package core

import (
	"fmt"
	"path/filepath"
	"workspaceCmd/utils"
)

// WorkSpaceGroup  工作组配置
type WorkSpaceGroup struct {
	GroupName           string
	GroupRemarks        string
	Path                string // 工作组路径
	WorkGroupConfigPath string // 工作组配置路径
	LastOpenMethod      string // 上次打开方式
}

func (receiver *WorkSpaceGroup) NewConfig() {
	defer func() {
		utils.GreenTips("new workgroup succeeded")
	}()
	receiver.GroupName = utils.GetLastFileNameDirectoryNamePath(receiver.Path)

	path := utils.CreateIniFile(receiver.GroupName)
	initHandler := utils.IniHelper{
		Path: path,
	}
	// 读ini文件
	initHandler.NewIni().
		// 新建工作组分区(section)
		NewSection(utils.WorkgroupConfigurationNameInIni).
		// 新建工作组分区字段
		NewKey(utils.WorkgroupConfigurationNameInIni, utils.WorkGroupName, receiver.GroupName).
		NewKey(utils.WorkgroupConfigurationNameInIni, utils.WorkGroupWithRemarks, receiver.GroupRemarks).
		NewKey(utils.WorkgroupConfigurationNameInIni, utils.WorkGroupPath, utils.JoinPwd(receiver.Path)).
		NewKey(utils.WorkgroupConfigurationNameInIni, utils.WorkGroupConfigPath, path).
		NewKey(utils.WorkgroupConfigurationNameInIni, utils.LastOpenMethodLabel, utils.LastOpenMethodLabelDefault).
		Save()
}

// WorkSpace 工作区配置
type WorkSpace struct {
	Path           string // 工作组路径
	WithGroup      string // 与哪个工作组绑定
	Remarks        string // 备注
	Name           string // 工作区名 非必填，默认路径最后路径（file name or dir name）
	GroupPath      string // 工作组路径 用于ls -d 指令删除时不用循环配置文件夹去找工作组文件
	LastOpenMethod string // 上次打开方式
}

func (receiver *WorkSpace) NewConfig() {
	defer func() {
		utils.GreenTips("new workspace succeeded")
	}()
	path := filepath.Join(utils.GetPwd(), utils.IniConfigurationFolder, fmt.Sprint(receiver.WithGroup, ".ini"))
	iniHandler := utils.IniHelper{
		Path: path,
	}

	receiver.Name = utils.GetLastFileNameDirectoryNamePath(receiver.Path)
	iniHandler.NewIni().
		NewSection(receiver.Name).
		NewKey(receiver.Name, utils.WorkSpacePath, utils.JoinPwd(receiver.Path)).
		NewKey(receiver.Name, utils.WorkSpaceRemarks, receiver.Remarks).
		NewKey(receiver.Name, utils.WorkSpaceWithGroupPath, path).
		NewKey(receiver.Name, utils.LastOpenMethodLabel, utils.LastOpenMethodLabelDefault).
		Save()
}

func CreateWork(conf Create) {
	conf.NewConfig()
}
