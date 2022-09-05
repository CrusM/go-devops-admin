package middleware

import (
	"fmt"
	"go-devops-admin/app/models/user"
	"go-devops-admin/pkg/config"
	"go-devops-admin/pkg/jwt"
	"go-devops-admin/pkg/response"

	"github.com/gin-gonic/gin"
)

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从header Authorization:Bearer 中获取 token 信息, 并验证准确性
		claims, err := jwt.NewJWT().ParserToken(c)

		// jwt 解析失败，有错误发生
		if err != nil {
			response.Unauthorized(c, fmt.Sprintf("请查看 %v 相关的接口认证文档", config.GetString("app.name")))
		}

		// 解析成功, 设置用户信息
		userModel := user.Get(claims.UserID)
		if userModel.ID == 0 {
			response.Unauthorized(c, "用户不存在")
		}

		c.Set("current_user_id", userModel.GetStringID())
		c.Set("current_user_name", userModel.Name)
		c.Set("current_user", userModel)

		c.Next()
	}
}
