package cmd

import (
	"go-devops-admin/pkg/helpers"
	"os"

	"github.com/spf13/cobra"
)

var Env string

// 注册全局 flag
func RegisterGlobalFlags(root *cobra.Command) {
	root.PersistentFlags().StringVarP(&Env, "env", "e", "", "load .env file")
}


// 注册默认命令
func RegisterDefaultCommand(rootCmd *cobra.Command, subCmd *cobra.Command) {
	cmd, _, err := rootCmd.Find(os.Args[1:])
	firstArg := helpers.FirstElement(os.Args[1:])
	if err == nil && cmd.Use == rootCmd.Use && firstArg != "-h" && firstArg != "--help" {
		args := append([]string{subCmd.Use}, os.Args[1:]...)
		rootCmd.SetArgs(args)
	}
}
