// 处理用户认证相关逻辑
package auth

import (
	"fmt"
	v1 "go-devops-admin/app/http/controllers/api/v1"
	"go-devops-admin/app/models/user"
	"go-devops-admin/app/requests"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 注册控制器
type SignUpController struct {
	v1.BaseAPIController
}

// 检查手机号是否被注册
func (sc *SignUpController) IsPhoneExist(c *gin.Context) {

	// 初始化请求对象
	request := requests.SignUpPhoneExistRequest{}

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

	// 表单验证
	errs := requests.ValidateSignUpPhoneExist(&request, c)
	// 返回的错误长度为0，则表示通过
	if len(errs) > 0 {
		// 验证失败，返回422状态码
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errs,
		})
		return

	}
	// 检查数据库并返回响应
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}

// 检查邮箱是否被注册
func (sc *SignUpController) IsEmailExist(c *gin.Context) {
	request := requests.SignUpEmailExistRequest{}

	// 格式化请求数据
	if err := c.ShouldBindJSON(&request); err != nil {
		// 解析失败，返回422状态码和错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		// 打印错误信息
		fmt.Println(err.Error())
		return
	}

	// 验证表单
	errs := requests.ValidateSignUpEmailExist(&request, c)
	if len(errs) > 0 {
		// 验证失败，返回422状态码
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errs,
		})
		return
	}

	// 检查数据库并返回响应
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsEmailExist(request.Email),
	})
}
