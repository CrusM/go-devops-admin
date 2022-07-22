/*
 * @Author: 蒋宏飞 jianghongfei@cnstrong.cn
 * @Date: 2022-07-22 16:34:15
 * @LastEditors: 蒋宏飞 jianghongfei@cnstrong.cn
 * @LastEditTime: 2022-07-22 17:00:39
 * @FilePath: \go-devops-admin\main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"net/http"

	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化gin实例
	r := gin.Default()

	// 注册一个路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Hello": "word",
		})
	})

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

	// 运行服务
	r.Run()
}
