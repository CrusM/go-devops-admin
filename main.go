package main

import (
	"flag"
	"fmt"
	"go-devops-admin/bootstrap"
	btsConfig "go-devops-admin/config"
	"go-devops-admin/pkg/config"

	"github.com/gin-gonic/gin"
)

func init() {
	btsConfig.Initialize()
}

func main() {

	// 配置初始化， 依赖命令行--env参数
	var env string
	flag.StringVar(&env, "env", "", "加载.env文件，如 --env=testing 加载的是 .env.testing 文件")
	flag.Parse()
	config.InitConfig(env)

	// 初始化日志配置
	bootstrap.SetupLogger()

	// 设置 gin 的运行模式，支持 debug, release, test
	// release 会屏蔽调试信息，官方建议生产环境中使用
	// 非 release 模式 gin 终端打印太多信息，干扰到我们程序中的 Log
	// 故此设置为 release，有特殊情况手动改为 debug 即可
	gin.SetMode(gin.ReleaseMode)

	// 初始化一个gin Engine实例
	r := gin.New()

	// 初始化DB
	bootstrap.SetupDB()

	// 初始化绑定路由
	bootstrap.SetupRoute(r)

	// 运行服务
	err := r.Run(":8080")
	if err != nil {
		// 错误处理，端口占用或者其他错误
		fmt.Println(err.Error())
	}
}
