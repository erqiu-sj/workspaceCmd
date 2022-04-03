package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"workspaceCmd/core"
	"workspaceCmd/utils"
)

//  新增工作区or工作组 (finish)
var (
	monitoringWorkingGroup  string // 监听工作组 工作组下的所有文件或文件夹生成对应的工作区和工作组
	excludeFolders          string // 监听工作组时，排除哪些文件夹
	addWorkspace            string // 新增一个工作区
	coreConfigWithWorkSpace = core.Check{}
)

var add = &cobra.Command{
	Use:   "add",
	Short: "add workspace",
	Long:  "add workspace(新增工作区)",
	PreRun: func(cmd *cobra.Command, args []string) {
		coreConfigWithWorkSpace.CheckWorkGroup()
		coreConfigWithWorkSpace.CheckWorkSpace()
	},
	Run: func(cmd *cobra.Command, args []string) {
		areYouReadyAddSpace(func() {
			addWorkspaceHandler()
		})
		areYouReadyMonitorWorkingGroup(func() {
			monitorAWorkingGroup()
		})
	},
}

func init() {
	add.Flags().StringVarP(&monitoringWorkingGroup, "listen", "l", "", fmt.Sprintln("turn the path folder into a workgroup, and the subfolders under the folder into a workspace\n(将路径文件夹转为工作组，文件夹下的子文件夹转为工作区)"))
	add.Flags().StringVarP(&excludeFolders, "exclude", "e", "", fmt.Sprintln("which folders are excluded when listening to workgroups\n(监听工作组时,排除哪些文件夹)"))
	add.Flags().StringVarP(&addWorkspace, "addSpace", "a", "", "add a workspace(新增一个工作区)")
	//monitoringWorkingGroup
	rootCmd.AddCommand(add)
}

func monitorAWorkingGroup() {
	if !utils.IsDir(monitoringWorkingGroup) {
		utils.BeforeStoppingProcess(func() {
			utils.RedTips("you listen to a folder")
		})
	}

	path := utils.JoinPwd(monitoringWorkingGroup)
	core.CreateWork(&core.WorkSpaceGroup{
		GroupRemarks: utils.NoComments,
		Path:         monitoringWorkingGroup,
	})

	ignoreFolderList := core.IgnoreFolder(
		utils.StringsToSlice(excludeFolders, ","),
		core.ListDir(path, false),
	)
	for _, meta := range ignoreFolderList {
		core.CreateWork(&core.WorkSpace{
			Path:      meta.Name,
			WithGroup: utils.GetLastFileNameDirectoryNamePath(path),
			Remarks:   utils.NoComments,
		})
	}
}

func addWorkspaceHandler() {
	if !utils.IsDir(addWorkspace) {
		utils.BeforeStoppingProcess(func() {
			utils.RedTips("you listen to a folder")
		})
	}

	path := utils.JoinPwd(addWorkspace)
	fmt.Println(path)
	tips := utils.InquiryTips("what are remarks?")
	result, queryError := tips.Run()
	utils.ColdKiller(queryError)
	workgroup, workgroupIndex := core.SelectWorkSpaceGroup(coreConfigWithWorkSpace.FinalAnalysis)
	curWorkGroup := workgroup[workgroupIndex]
	iniFileHandler := utils.IniHelper{
		Path: curWorkGroup.WorkGroupConfigPath,
	}
	workspaceName := utils.GetLastFileNameDirectoryNamePath(path)
	iniFileHandler.NewIni().
		NewSection(workspaceName).
		NewKey(workspaceName, utils.WorkSpaceRemarks, result).
		NewKey(workspaceName, utils.WorkSpacePath, path).
		NewKey(workspaceName, utils.WorkSpaceWithGroupPath, curWorkGroup.WorkGroupConfigPath).
		Save()
	utils.GreenTips("added workspace successfully")
}

func areYouReadyMonitorWorkingGroup(cb func()) {
	if utils.NotEmpty(addWorkspace) {
		return
	}
	if utils.NotEmpty(monitoringWorkingGroup) {
		cb()
	}
}

func areYouReadyAddSpace(cb func()) {
	if utils.NotEmpty(addWorkspace) {
		cb()
	}
}
