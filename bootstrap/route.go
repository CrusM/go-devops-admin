package bootstrap

import (
	"go-devops-admin/routes"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 初始化路由
func SetupRoute(r *gin.Engine) {
	// 注册全局中间件
	registerGlobalMiddleWare(r)

	// 注册api路由
	routes.RegisterAPIRouters(r)

	// 注册404路由
	setup404Handler(r)
}

func registerGlobalMiddleWare(r *gin.Engine) {
	r.Use(
		gin.Logger(),
		gin.Recovery(),
	)
}

func setup404Handler(r *gin.Engine) {
	// 处理404请求
	r.NoRoute(func(c *gin.Context) {
		// 获取标头信息的 Accept 信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// 如果是HTML的话, 返回404页面
			c.String(http.StatusNotFound, "页面返回404")
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"code": 404,
				"msg":  "路由未定义, 请确认url和请求方法是否正确",
			})
		}
	})
}
