package main

import (
	"fmt"
	"go-devops-admin/app/cmd"
	"go-devops-admin/bootstrap"
	btsConfig "go-devops-admin/config"
	"go-devops-admin/pkg/config"
	"go-devops-admin/pkg/console"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	btsConfig.Initialize()
}

func main() {

	var rootCmd = &cobra.Command{
		Use:   "GoDev",
		Short: "A simple from project",
		Long:  `Default will run "server" command, you can use "-h" flag to see all command`,

		// rootCmd 所有子命令都会执行以下代码
		PersistentPreRun: func(command *cobra.Command, args []string) {
			// 配置初始化, 依赖命令行 --env 参数
			config.InitConfig(cmd.Env)

			// 初始化日志配置
			bootstrap.SetupLogger()
			// 初始化 DB
			bootstrap.SetupDB()

			// 初始化 Redis
			bootstrap.SetupRedis()
		},
	}

	// 注册子命令
	rootCmd.AddCommand(
		cmd.CmdServer,
	)

	// 配置默认运行 Web 服务
	cmd.RegisterDefaultCommand(rootCmd, cmd.CmdServer)

	// 执行主命令
	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}
}
