package middleware

import (
	"errors"
	"go-devops-admin/pkg/response"

	"github.com/gin-gonic/gin"
)

// 强制请求 headers 中必须带 User-Agent
func ForceUA() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 User-Agent
		if len(c.Request.Header["User-Agent"]) == 0 {
			response.BadRequest(c, errors.New("User-Agent 标头未找到"), "请求中必须附带 User-Agent 标头")
		}
	}
}
