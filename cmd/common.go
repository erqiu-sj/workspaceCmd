package cmd

import (
	"github.com/spf13/cobra"
	"workspaceCmd/core"
	"workspaceCmd/utils"
)

var (
	addProjectMode bool
)

var common = &cobra.Command{
	Use:   "common",
	Short: "some configuration modifications",
	Long:  "some configuration modifications(通用配置)",
	Run: func(cmd *cobra.Command, args []string) {
		onAddProjectMode(func() {
			addProjectModeHandler()
		})
	},
}

func init() {
	common.Flags().BoolVarP(&addProjectMode, "addMode", "a", false, "add open project mode(新增打开项目方式)")
	rootCmd.AddCommand(common)
}

func onAddProjectMode(cb func()) {
	if addProjectMode {
		cb()
	}
}

func addProjectModeHandler() {
	labelTips := utils.InquiryTips("give the open mode a name!")
	label, labelErr := labelTips.Run()
	utils.ColdKiller(labelErr)
	cmdTips := utils.InquiryTips(`what global variables are used to use this pattern? For example, vscode uses code open`)
	cmd, cmdErr := cmdTips.Run()
	utils.ColdKiller(cmdErr)
	core.CreateWork(
		&core.OpenMode{
			Label:    label,
			CmdAlias: cmd,
		},
	)
}
