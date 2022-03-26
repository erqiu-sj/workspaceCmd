package core

import (
	"fmt"
	"gopkg.in/ini.v1"
	"io/ioutil"
	"path/filepath"
	"workspaceCmd/utils"
)

type FinalAnalysisConf struct {
	WorkSpaceGroup
	SpaceList []WorkSpace
}

type Check struct {
	workGroupConfigure []ListDirMeta     // 解析目录
	FinalAnalysis      FinalAnalysisConf // 最终解析
}

func (receiver *Check) CheckWorkSpace() {
	for _, meta := range receiver.workGroupConfigure {
		conf, parseErr := ini.Load(meta.Path)
		utils.ColdKiller(parseErr)
		// 排除default section
		names := conf.SectionStrings()[1:]
		for _, i := range names {
			if i == utils.WorkgroupConfigurationNameInIni {
				curSection := conf.Section(i)
				// 判断工作组是否存在，不存在则删除文件，
				if !utils.DoesFolderExist(utils.WorkGroupPath) {
					utils.GreenTips(fmt.Sprint("(", i, ")", " the workgroup no longer exists and has been removed from the configuration"))
					utils.RemoveFile(meta.Path)
					break
				}
				receiver.FinalAnalysis.GroupName = curSection.Key(utils.WorkGroupName).Value()
				receiver.FinalAnalysis.GroupRemarks = curSection.Key(utils.WorkGroupWithRemarks).Value()
				receiver.FinalAnalysis.Path = curSection.Key(utils.WorkGroupPath).Value()
				continue
			}
			space := conf.Section(i)
			spacePath := space.Key(utils.WorkSpacePath).Value()
			// 检测工作区路径是否存在
			if !utils.DoesFolderExist(spacePath) {
				// 不存在删除
				conf.DeleteSection(i)
				saveConf := conf.SaveTo(meta.Path)
				utils.GreenTips(fmt.Sprint("(", i, ")", " the workspace no longer exists, so it has been removed from the configuration"))
				utils.ColdKiller(saveConf)
				continue
			}
			receiver.FinalAnalysis.SpaceList = append(receiver.FinalAnalysis.SpaceList,
				WorkSpace{
					Path:    spacePath,
					Remarks: space.Key(utils.WorkSpaceRemarks).Value(),
					Name:    i,
				},
			)
		}
	}
}

func (receiver *Check) CheckWorkGroup() {
	receiver.workGroupConfigure = ListDir(utils.IniConfigurationFolder)
}

type ListDirMeta struct {
	Path string
	Name string
}

func ListDir(folder string) []ListDirMeta {
	var list []ListDirMeta
	files, errDir := ioutil.ReadDir(folder)
	utils.ColdKiller(errDir)
	for _, file := range files {
		// 输出绝对路径
		strAbsPath, errPath := filepath.Abs(folder + "/" + file.Name())
		utils.ColdKiller(errPath)
		list = append(list, ListDirMeta{Name: file.Name(), Path: strAbsPath})
	}
	return list
}
