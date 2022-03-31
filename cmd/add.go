package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"workspaceCmd/core"
	"workspaceCmd/utils"
)

//  新增工作区or工作组 (finish)
var (
	monitoringWorkingGroup string // 监听工作组 工作组下的所有文件或文件夹生成对应的工作区和工作组
	excludeFolders         string // 监听工作组时，排除哪些文件夹
)

var add = &cobra.Command{
	Use:   "add",
	Short: "add workspace",
	Long:  "add workspace(新增工作区)",
	Run: func(cmd *cobra.Command, args []string) {
		areYouReadyMonitorWorkingGroup(func() {
			monitorAWorkingGroup()
		})
	},
}

func init() {
	add.Flags().StringVarP(&monitoringWorkingGroup, "listen", "l", "", fmt.Sprintln("turn the path folder into a workgroup, and the subfolders under the folder into a workspace\n(将路径文件夹转为工作组，文件夹下的子文件夹转为工作区)"))
	add.Flags().StringVarP(&excludeFolders, "exclude", "e", "", fmt.Sprintln("which folders are excluded when listening to workgroups\n(监听工作组时,排除哪些文件夹)"))
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

func areYouReadyMonitorWorkingGroup(cb func()) {
	if utils.NotEmpty(monitoringWorkingGroup) {
		cb()
	}
}
