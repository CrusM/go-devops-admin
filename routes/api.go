package routes

import (
	"go-devops-admin/app"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRouters(r *gin.Engine) {
	// 测试一个 v1 的路由组，所有v1版本的路由都存在这里
	// var v1 *gin.RouterGroup
	// v1.GET("/", func(ctx *gin.Context) {
	// 	// 返回JSON格式数据
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"Hello": "World!",
	// 	})
	// })
	// var g *gin.RouterGroup
	// d := g.Group("")
	// {
	// 	d.GET("/", func(ctx *gin.Context) {
	// 		ctx.JSON(http.StatusOK, gin.H{
	// 			"Hello": "World!",
	// 		})
	// 	})
	// }
	app.InitRegistryRouter(r)
}
