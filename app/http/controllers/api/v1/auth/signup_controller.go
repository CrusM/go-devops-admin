// 处理用户认证相关逻辑
package auth

import (
	"fmt"
	v1 "go-devops-admin/app/http/controllers/api/v1"
	"go-devops-admin/app/models/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 注册控制器
type SignUpController struct {
	v1.BaseAPIController
}

// 检查收集是否被注册
func (sc *SignUpController) IsPhoneExist(c *gin.Context) {
	// 请求对象
	type PhoneExistRequest struct {
		Phone string `json:"-"`
	}

	request := PhoneExistRequest{}

	// 解析JSON请求
	if err := c.ShouldBindJSON(&request); err != nil {
		// 解析失败，返回422状态码和错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		// 打印错误信息
		fmt.Println(err.Error())
		return
	}

	// 检查数据库并返回响应
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}
