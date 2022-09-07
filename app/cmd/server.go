package cmd

import (
	"go-devops-admin/bootstrap"
	"go-devops-admin/pkg/config"
	"go-devops-admin/pkg/console"
	"go-devops-admin/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var CmdServer = &cobra.Command{
	Use:   "server",
	Short: "Start web server",
	Run:   runWeb,
}

func runWeb(cmd *cobra.Command, args []string) {
	// 设置 gin 的运行模式，支持 debug, release, test
	// release 会屏蔽调试信息，官方建议生产环境中使用
	// 非 release 模式 gin 终端打印太多信息，干扰到我们程序中的 Log
	// 故此设置为 release，有特殊情况手动改为 debug 即可
	gin.SetMode(gin.ReleaseMode)

	// gin 实例
	r := gin.New()

	// 初始化 DB
	// bootstrap.SetupDB()

	// 初始化 Redis
	// bootstrap.SetupRedis()

	// 初始化绑定路由
	bootstrap.SetupRoute(r)

	// 运行服务
	err := r.Run(":" + config.Get("app.port"))
	if err != nil {
		// 错误处理，端口占用或者其他错误
		logger.ErrorString("CMD", "server", err.Error())
		console.Exit("Unable to start server, error:" + err.Error())
	}
}
