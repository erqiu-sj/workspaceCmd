package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/ini.v1"
	"workspaceCmd/core"
	"workspaceCmd/utils"
)

var (
	viewAllWorkgroups bool
	deleteWorkSpace   bool
	editWorkSpace     string
	editWorkGroup     bool // 修改工作组参数 目前工作组参数只能修改备注
	openProject       bool // 打开工作区或者工作组
	alwaysSelect      bool // 打开工作区时 始终选择打开
	coreConfig        = core.Check{}
)

const (
	EditMode_DESC = "desc"
	EditMode_Path = "path"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "view workgroups and workspaces",
	Long:  "view workgroups and workspaces(查看工作组和工作区)",
	PreRun: func(cmd *cobra.Command, args []string) {
		coreConfig.CheckWorkGroup()
		coreConfig.CheckWorkSpace()
	},
	Run: func(cmd *cobra.Command, args []string) {
		// 打开工作组
		core.CheckParameter(openProject && editWorkGroup, func() {
			previewWorkspaceGroupHandler()
		})
		// 打开工作区
		core.CheckParameter(openProject, func() {
			openWorkspaceHandler()
		})
		// 只读工作组和工作区
		// 如删除 修改工作区 打开工作区操作 则不执行
		core.CheckParameter(viewAllWorkgroups, func() {
			previewWorkspaceGroupHandler()
		})
		// 删除工作组
		core.CheckParameter(deleteWorkSpace && editWorkGroup, func() {
			removeWorkGroup()
		})
		// 删除工作区
		// 如有工作组的参数则不执行
		core.CheckParameter(deleteWorkSpace, func() {
			deleteWorkSpaceHandler()
		})
		// 修改工作组
		core.CheckParameter(utils.NotEmpty(editWorkSpace) && editWorkSpace == EditMode_DESC && editWorkGroup, func() {
			editWorkGroupHandler()
		})
		// 修改工作区
		core.CheckParameter(utils.NotEmpty(editWorkSpace) || editWorkSpace == EditMode_DESC || editWorkSpace == EditMode_Path, func() {
			editWorkSpaceHandler()
		})
	},
}

func init() {
	lsCmd.Flags().BoolVarP(&viewAllWorkgroups, "workspaceGroup", "l", true, `view workgroups and workspaces(查看工作区和工作区)`)
	lsCmd.Flags().BoolVarP(&deleteWorkSpace, "deleteWorkspace", "d", false, "delete a workspace , if - D is added, the working group will be deleted(删除某个工作区,如带上-g 则是删除工作组)")
	lsCmd.Flags().StringVarP(&editWorkSpace, "editWorkspace", "e", "", "edit a workspace[desc,[path]](修改某个工作区)")
	lsCmd.Flags().BoolVarP(&editWorkGroup, "editWorkGroup", "g", false, `bring - E = [desc] to modify the remarks of the working group. 
If you bring - D, the working group will be deleted, 
and the deletion operation takes precedence over the modification operation
( 带上-e=[desc] 能修改工作组备注，如带上-d 则会删除工作组，且删除操作优先级高于修改操作)`)
	lsCmd.Flags().BoolVarP(&openProject, "open", "o", false, "open a workgroup or workspace (打开工作组或者工作区)")
	lsCmd.Flags().BoolVarP(&alwaysSelect, "alwaysSelect", "a", false, "always select open workspace or work group(始终选择打开工作区或工作组)")
	//alwaysSelect
	rootCmd.AddCommand(lsCmd)
}

func StartSelect() ([]core.WorkSpace, int) {
	result, i := core.SelectWorkSpaceGroup(coreConfig.FinalAnalysis)
	workList, workIndex := core.SelectWorkSpace(coreConfig.FinalAnalysis, result[i].GroupName)
	return workList, workIndex
}

// previewWorkspaceGroupHandler 预览工作区和工作组
func previewWorkspaceGroupHandler() {
	utils.GreenTips("browsing workgroups and workspaces(正在浏览工作组和工作区)")
	StartSelect()
}

// deleteWorkSpaceHandler 删除某个工作区
func deleteWorkSpaceHandler() {
	utils.RedTips("please be careful while deleting the workspace (正在执行删除工作区操作，请谨慎)")
	workList, workIndex := StartSelect()
	rightWrongList := []utils.ConciseSelector{
		{Label: "yes", Desc: "this operation is irreversible"},
		{Label: "no", Desc: "no desc"},
	}
	modelTips := utils.WarningModelTips("are you sure to delete the workspace?", rightWrongList)
	runIndex, _, err := modelTips.Run()
	utils.ColdKiller(err)
	fmt.Println(rightWrongList[runIndex].Label)
	if rightWrongList[runIndex].Label == "yes" {
		iniFile, iniErr := ini.Load(workList[workIndex].GroupPath)
		utils.ColdKiller(iniErr)
		utils.DeleteKey(iniFile, workList[workIndex].Name, workList[workIndex].GroupPath)
		utils.GreenTips("delete workspace succeeded")
	}
}

//  editWorkSpaceHandler 修改工作区回调
// 修改工作区会有两种模式
// desc 修改备注
// path 修改路径
func editWorkSpaceHandler() {
	if editWorkSpace == EditMode_DESC {
		descWorkList, descWorkIndex := StartSelect()
		curWorks := descWorkList[descWorkIndex]
		inquiryTips := utils.InquiryTips("what is a new remarks?")
		newRemarks, remarksErr := inquiryTips.Run()
		utils.ColdKiller(remarksErr)
		utils.EditWorkSpaceOrGroupField(utils.EditWorkSpaceOrGroupFieldConfig{
			SectionName: curWorks.Name,
			Key:         utils.WorkSpaceRemarks,
			Val:         newRemarks,
			Path:        curWorks.GroupPath,
		})
	}
	if editWorkSpace == EditMode_Path {

		pathWorkList, pathWorkIndex := StartSelect()
		inquiryTips := utils.InquiryTips("what is a new path?")
		newPath, editNewPathErr := inquiryTips.Run()
		utils.ColdKiller(editNewPathErr)
		if !utils.DoesFolderExist(utils.JoinPwd(newPath)) {
			// 文件路径不存在，请检查
			utils.RedTips("file path does not exist, please check")
			return
		}
		curPathWork := pathWorkList[pathWorkIndex]
		utils.EditWorkSpaceOrGroupField(utils.EditWorkSpaceOrGroupFieldConfig{
			SectionName: curPathWork.Name,
			Key:         utils.WorkSpacePath,
			Val:         utils.JoinPwd(newPath),
			Path:        curPathWork.GroupPath,
		})
	}
}

// editWorkGroupHandler 修改工作组备注
func editWorkGroupHandler() {
	result, i := core.SelectWorkSpaceGroup(coreConfig.FinalAnalysis)
	tips := utils.InquiryTips("what is a new remarks?")
	newRemarks, err := tips.Run()
	utils.ColdKiller(err)
	curResult := result[i]
	utils.EditWorkSpaceOrGroupField(utils.EditWorkSpaceOrGroupFieldConfig{
		SectionName: utils.WorkgroupConfigurationNameInIni,
		Key:         utils.WorkGroupWithRemarks,
		Val:         newRemarks,
		Path:        curResult.WorkGroupConfigPath,
	})
}

// removeWorkGroup 删除工作组
func removeWorkGroup() {
	utils.RedTips("please be careful while deleting the workgroup (目前正在执行删除工作组操作，请谨慎)")
	result, i := core.SelectWorkSpaceGroup(coreConfig.FinalAnalysis)
	warningOptions := []utils.ConciseSelector{
		{Label: "yes, I have to delete it(是的，我必须删除它)", Desc: "this operation is irreversible. Please be careful again(此操作不可逆，请谨慎再谨慎)"},
		{
			Label: "no, sorry, I'm wrong(不好意思，点错了)",
			Desc:  "my God, thank God(我的天呐，谢天谢地)",
		},
	}
	tips := utils.WarningModelTips("delete this working group(是否删除改工作组)", warningOptions)
	runIndex, _, err := tips.Run()
	utils.ColdKiller(err)
	if runIndex == 0 {
		utils.RemoveFile(result[i].WorkGroupConfigPath)
	}
}

func openWorkspaceHandler() {
	openWork, openIndex := StartSelect()
	curWorkSpace := openWork[openIndex]
	selectOpen := func() {
		openMode := core.CreateOpenModeSelect()
		iniHandler := utils.IniHelper{Path: curWorkSpace.GroupPath}
		utils.Shell(
			fmt.Sprint(openMode.CmdAlias, " ", curWorkSpace.Path),
			func() {
				iniHandler.NewIni().EditKey(curWorkSpace.Name, utils.LastOpenMethodLabel, openMode.Label).Save()
			},
			func(msg string) {
				utils.RedTips(msg)
			},
		)
	}
	if alwaysSelect {
		selectOpen()
		return
	} else if curWorkSpace.LastOpenMethod == utils.LastOpenMethodLabelDefault {
		// 第一次打开工作区或工作区,选择打开方式
		selectOpen()
		return
	}
	// 不是第一次打开
	openHandler := core.OpenMode{
		Label: curWorkSpace.LastOpenMethod,
	}
	conf, _ := openHandler.GetItem()
	utils.Shell(
		fmt.Sprint(conf.CmdAlias, " ", curWorkSpace.Path),
		func() {

		},
		func(msg string) {
			utils.RedTips(msg)
		},
	)
}
