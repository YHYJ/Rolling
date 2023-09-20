/*
File: version.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-02-17 16:08:20

Description: 程序子命令'version'时执行
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yhyj/rolling/function"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print program version",
	Long:  `Print program version and exit.`,
	Run: func(cmd *cobra.Command, args []string) {
		// 解析参数
		onlyFlag, _ := cmd.Flags().GetBool("only")

		programInfo := function.ProgramInfo(onlyFlag)
		fmt.Printf(programInfo)
	},
}

func init() {
	versionCmd.Flags().BoolP("only", "", false, "Only print the version number, like 'v0.1.1'")

	versionCmd.Flags().BoolP("help", "h", false, "help for version")
	rootCmd.AddCommand(versionCmd)
}
