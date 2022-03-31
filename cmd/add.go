package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"path/filepath"
	"workspaceCmd/core"
	"workspaceCmd/utils"
)

//  新增工作区or工作组 (finish)
var (
	remark                 string // 工作区备注
	addWorkspacePath       string // 新增工作区路径
	workingGroupNotes      string // 工作组备注
	addAWorkingGroup       string // 新增工作组 路径
	workSpaceWithGroup     string // 工作区属于什么组
	monitoringWorkingGroup string // 监听工作组 工作组下的所有文件或文件夹生成对应的工作区和工作组
	excludeFolders         string // 监听工作组时，排除哪些文件夹
)

var add = &cobra.Command{
	Use:   "add",
	Short: "add workspace",
	Long:  "add workspace(新增工作区)",
	PreRun: func(cmd *cobra.Command, args []string) {
		// 存在新增工作区的意向
		areYouReadyAddNewWorkspace(func() {
			addWorkSpaceVerifyHandler()
		})
		areYouReadyAddNewWorkGroup(func() {
			addWorkGroupVerifyHandler()
		})
	},
	Run: func(cmd *cobra.Command, args []string) {
		areYouReadyMonitorWorkingGroup(func() {
			monitorAWorkingGroup()
		})
		areYouReadyAddNewWorkGroup(func() {
			core.CreateWork(&core.WorkSpaceGroup{
				GroupRemarks: workingGroupNotes,
				Path:         addAWorkingGroup,
			})
		})
		areYouReadyAddNewWorkspace(func() {
			core.CreateWork(&core.WorkSpace{
				Path:      addWorkspacePath,
				WithGroup: workSpaceWithGroup,
				Remarks:   remark,
			})
		})
	},
}

func init() {
	add.Flags().StringVarP(&addWorkspacePath, "add", "a", "", fmt.Sprintln("new workspace, path is required(新增工作区，路径为必填)"))
	add.Flags().StringVarP(&remark, "remark", "r", "", fmt.Sprintln("workspace remarks, required when adding a workspace(工作区备注,新增工作区时必填)"))
	add.Flags().StringVarP(&workingGroupNotes, "workingGroupNotes", "g", "", fmt.Sprintln("working group notes, required when adding a working group(工作组备注，新增工作组时必填)"))
	add.Flags().StringVarP(&addAWorkingGroup, "addAWorkingGroup", "m", "", fmt.Sprintln("add a new workgroup, and the path is required(新增一个工作组，路径为必填项)"))
	add.Flags().StringVarP(&workSpaceWithGroup, "with", "w", "", fmt.Sprintln("bind the workspace to the workgroup, which is required when adding a workspace(将工作区与工作组绑定，新增工作区时必填)"))
	add.Flags().StringVarP(&monitoringWorkingGroup, "listen", "l", "", fmt.Sprintln("turn the path folder into a workgroup, and the subfolders under the folder into a workspace(将路径文件夹转为工作组，文件夹下的子文件夹转为工作区)"))
	add.Flags().StringVarP(&excludeFolders, "exclude", "e", "", fmt.Sprintln("which folders are excluded when listening to workgroups(监听工作组时,排除哪些文件夹)"))
	//monitoringWorkingGroup
	rootCmd.AddCommand(add)
}

type validationGroup struct {
}

// verifyPath 验证路径
func (that *validationGroup) verifyPath(verifyPath string) *validationGroup {
	path := utils.JoinPwd(verifyPath)
	if !utils.DoesFolderExist(path) {
		utils.BeforeStoppingProcess(func() {
			utils.RedTips(
				utils.NotAFile,
			)
		})
	}
	return that
}

// verificationRemarks 验证备注
func (that *validationGroup) verificationString(remarks string, tips string) *validationGroup {
	if !utils.NotEmpty(remarks) {
		utils.BeforeStoppingProcess(func() {
			utils.RedTips(fmt.Sprint(tips))
		})
	}
	return that
}

func (that *validationGroup) validationWorkingGroup(path, tips string) *validationGroup {
	parsePath := filepath.Join(utils.GetPwd(), utils.IniConfigurationFolder, fmt.Sprint(path, ".ini"))
	println(parsePath)
	if !utils.DoesFolderExist(parsePath) {
		utils.BeforeStoppingProcess(func() {
			utils.RedTips(fmt.Sprint(tips))
		})
	}
	return that
}

// addWorkSpaceHandler 新增工作区 add workspace
// cannot be empty field (addWorkspacePath, remark, workSpaceWithGroup)
func addWorkSpaceVerifyHandler() {
	verifyGroup := validationGroup{}
	// 验证工作区 verify workspace
	verifyGroup.verifyPath(addWorkspacePath)
	// 验证备注 verify remarks
	verifyGroup.verificationString(remark, fmt.Sprint("workspace ", utils.RemarksNotEmpty))
	// 验证工作组
	verifyGroup.validationWorkingGroup(workSpaceWithGroup, fmt.Sprint(utils.NotExitsWorkGroup))
}

func addWorkGroupVerifyHandler() {
	verifyGroup := validationGroup{}
	// 工作组备注不能为空
	verifyGroup.verificationString(workingGroupNotes, fmt.Sprint("workgroup ", utils.RemarksNotEmpty))
	// 工作组路径不能为空
	verifyGroup.verificationString(addAWorkingGroup, fmt.Sprint("workGroup path", utils.CannotBeEmpty))
	// 验证工作组路径
	verifyGroup.verifyPath(addAWorkingGroup)
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

func areYouReadyAddNewWorkspace(cb func()) {
	if utils.NotEmpty(monitoringWorkingGroup) {
		return
	}
	if utils.NotEmpty(addWorkspacePath) {
		cb()
	}
}
func areYouReadyAddNewWorkGroup(cb func()) {
	if utils.NotEmpty(monitoringWorkingGroup) {
		return
	}
	if utils.NotEmpty(addAWorkingGroup) {
		cb()
	}
}
func areYouReadyMonitorWorkingGroup(cb func()) {
	if utils.NotEmpty(monitoringWorkingGroup) {
		cb()
	}
}
