package core

import (
	"fmt"
	"io/ioutil"
	"workspaceCmd/utils"
)

type FinalAnalysisConf struct {
	WorkSpaceGroup
	SpaceList []WorkSpace
}

type Check struct {
	workGroupConfigure []ListDirMeta       // 解析目录
	FinalAnalysis      []FinalAnalysisConf // 最终解析
}

func (receiver *Check) CheckWorkSpace() {
	//utils.GreenTips("checking workspace")
	for _, meta := range receiver.workGroupConfigure {
		iniHandler := utils.IniHelper{
			Path: meta.Path,
		}
		handler := iniHandler.NewIni()
		// 排除default section
		names := handler.GetSectionToStrings()
		var curFinalAnalysis FinalAnalysisConf
		for _, i := range names {
			if i == utils.WorkgroupConfigurationNameInIni {
				groupPath := handler.GetKey(i, utils.WorkGroupPath)
				// 判断工作组是否存在，不存在则删除文件，
				if !utils.DoesFolderExist(groupPath) {
					utils.GreenTips(fmt.Sprint("(", i, ")", " the workgroup no longer exists and has been removed from the configuration"))
					utils.RemoveFile(meta.Path)
					break
				}
				curFinalAnalysis.GroupName = utils.GetLastFileNameDirectoryNamePath(groupPath)
				curFinalAnalysis.GroupRemarks = handler.GetKey(i, utils.WorkGroupWithRemarks)
				curFinalAnalysis.Path = groupPath
				curFinalAnalysis.WorkGroupConfigPath = meta.Path
				curFinalAnalysis.LastOpenMethod = utils.StringDefault(handler.GetKey(i, utils.LastOpenMethodLabel), utils.LastOpenMethodLabelDefault)
				continue
			}
			spacePath := handler.GetKey(i, utils.WorkSpacePath)
			// 检测工作区路径是否存在
			if !utils.DoesFolderExist(spacePath) {
				// 不存在删除
				handler.DeleteSection(i)
				handler.Save()
				utils.GreenTips(fmt.Sprint("(", i, ")", " the workspace no longer exists, so it has been removed from the configuration"))
				continue
			}
			curFinalAnalysis.SpaceList = append(curFinalAnalysis.SpaceList,
				WorkSpace{
					Path:           spacePath,
					Remarks:        handler.GetKey(i, utils.WorkSpaceRemarks),
					Name:           i,
					GroupPath:      meta.Path,
					WithGroup:      curFinalAnalysis.GroupName,
					LastOpenMethod: utils.StringDefault(handler.GetKey(i, utils.LastOpenMethodLabel), utils.LastOpenMethodLabelDefault),
				},
			)
		}
		receiver.FinalAnalysis = append(receiver.FinalAnalysis, curFinalAnalysis)
	}
}

func (receiver *Check) CheckWorkGroup() {
	receiver.workGroupConfigure = ListDir(utils.IniConfigurationFolder, true)
}
func IsExcludeFolders(path string) bool {
	disList := []string{".git", ".idea"}
	for _, i := range disList {
		if i == path {
			return true
		}
	}
	return false
}

type ListDirMeta struct {
	Path string
	Name string
}

func ListDir(folder string, allowLogFiles bool) []ListDirMeta {
	var list []ListDirMeta
	files, errDir := ioutil.ReadDir(folder)
	utils.ColdKiller(errDir)
	for _, file := range files {
		if IsExcludeFolders(file.Name()) {
			continue
		}
		path := fmt.Sprint(folder + "/" + file.Name())
		//utils.ColdKiller(errPath)
		if allowLogFiles {
			list = append(list, ListDirMeta{Name: file.Name(), Path: path})
		} else if file.IsDir() {
			//// 输出绝对路径
			//strAbsPath, errPath := filepath.Abs(folder + "/" + file.Name())
			//utils.ColdKiller(errPath)
			list = append(list, ListDirMeta{Name: file.Name(), Path: path})
		}
	}
	return list
}

func IgnoreFolder(list []string, listDir []ListDirMeta) []ListDirMeta {
	if len(list) == 0 {
		return listDir
	}
	var result []ListDirMeta
	for _, meta := range listDir {
		if !utils.FromSliceFindString(list, meta.Name) {
			result = append(result, meta)
		}
	}
	return result
}

// CheckParameter
// 如果有组合参数 操作应该把组合参数操作放在最前面，以此类推，
// 这样方便执行完后返回1停止程序
func CheckParameter(rule bool, cb func()) {
	if rule {
		utils.BeforeStoppingProcess(
			cb,
		)
	}
}
