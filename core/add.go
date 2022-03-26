package core

import (
	"fmt"
	"gopkg.in/ini.v1"
	"path/filepath"
	"workspaceCmd/utils"
)

// WorkSpaceGroup  工作组配置
type WorkSpaceGroup struct {
	GroupName    string
	GroupRemarks string
	Path         string
}

func (receiver *WorkSpaceGroup) NewConfig() {
	defer func() {
		utils.GreenTips("new workgroup succeeded")
	}()

	receiver.GroupName = utils.GetLastFileNameDirectoryNamePath(receiver.Path)

	path := utils.CreateIniFile(receiver.GroupName)
	// 读ini文件
	iniFile, parseError := ini.Load(path)
	utils.ColdKiller(parseError)
	// 新建工作组分区(section)
	_, newSectionErr := iniFile.NewSection(utils.WorkgroupConfigurationNameInIni)
	utils.ColdKiller(newSectionErr)
	// 新建工作组分区字段
	_, newNameKeyErr := iniFile.Section(utils.WorkgroupConfigurationNameInIni).NewKey(utils.WorkGroupName, receiver.GroupName)
	utils.ColdKiller(newNameKeyErr)
	_, newRemarksErr := iniFile.Section(utils.WorkgroupConfigurationNameInIni).NewKey(utils.WorkGroupWithRemarks, receiver.GroupRemarks)
	utils.ColdKiller(newRemarksErr)
	_, newPathErr := iniFile.Section(utils.WorkgroupConfigurationNameInIni).NewKey(utils.WorkGroupPath, utils.JoinPwd(receiver.Path))
	utils.ColdKiller(newPathErr)
	saveErr := iniFile.SaveTo(path)
	utils.ColdKiller(saveErr)
}

// WorkSpace 工作区配置
type WorkSpace struct {
	Path      string // 工作组路径
	WithGroup string // 与哪个工作组绑定
	Remarks   string // 备注
	Name      string // 工作区名 非必填，默认路径最后路径（file name or dir name）
}

func (receiver *WorkSpace) NewConfig() {
	defer func() {
		utils.GreenTips("new workspace succeeded")
	}()
	path := filepath.Join(utils.GetPwd(), utils.IniConfigurationFolder, fmt.Sprint(receiver.WithGroup, ".ini"))
	iniFile, parseErr := ini.Load(path)
	utils.InterceptErrorsAndKillProcessImmediately(parseErr, func(msg string) {
		utils.RedTips(utils.CheckWorkSpaceOrWorkGroup)
	})
	receiver.Name = utils.GetLastFileNameDirectoryNamePath(receiver.Path)
	// 新建工作区 section
	_, newSectionErr := iniFile.NewSection(receiver.Name)
	utils.ColdKiller(newSectionErr)
	// 配置工作区
	_, newPathKeyErr := iniFile.Section(receiver.Name).NewKey(utils.WorkSpacePath, utils.JoinPwd(receiver.Path))
	utils.ColdKiller(newPathKeyErr)
	_, newRemarksErr := iniFile.Section(receiver.Name).NewKey(utils.WorkSpaceRemarks, receiver.Remarks)
	utils.ColdKiller(newRemarksErr)
	saveErr := iniFile.SaveTo(path)
	utils.ColdKiller(saveErr)
}

func CreateWork(conf Create) {
	conf.NewConfig()
}
