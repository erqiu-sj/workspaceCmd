package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"workspaceCmd/utils"
)

var (
	version bool
)

//var coreConfig = core.Check{}

var rootCmd = &cobra.Command{
	Use:   "work",
	Short: "without parameters, the specified workspace will be checked by default",
	Long:  "without parameters, the specified workspace will be checked by default(不带参数的情况下，默认会检查指定工作区)",
	Run: func(cmd *cobra.Command, args []string) {
		printVersion()
	},
}

func Execute() {
	rootCmd.PersistentFlags().BoolVarP(&version, "version", "v", false, fmt.Sprintln("print version for", utils.CmdName))
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func printVersion() {
	if version {
		utils.BeforeStoppingProcess(func() {
			utils.GreenTips(
				fmt.Sprint(utils.CmdName, " version is ", utils.Version),
			)
		})
	}
}
