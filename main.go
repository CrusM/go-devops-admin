package main

import (
	"fmt"
	"go-devops-admin/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化一个gin Engine实例
	r := gin.New()

	// 初始化绑定路由
	bootstrap.SetupRoute(r)

	// 运行服务
	err := r.Run(":8080")
	if err != nil {
		// 错误处理，端口占用或者其他错误
		fmt.Println(err.Error())
	}
}
