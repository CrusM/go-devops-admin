// 处理用户认证相关逻辑
package auth

import (
	v1 "go-devops-admin/app/http/controllers/api/v1"
	"go-devops-admin/app/models/user"
	"go-devops-admin/app/requests"
	"go-devops-admin/pkg/response"

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

	// 校验请求数据
	if ok := requests.Validate(c, request, requests.ValidateSignUpPhoneExist); !ok {
		return
	}
	// 检查数据库并返回响应
	response.JSON(c, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}

// 检查邮箱是否被注册
func (sc *SignUpController) IsEmailExist(c *gin.Context) {
	request := requests.SignUpEmailExistRequest{}

	if ok := requests.Validate(c, &request, requests.ValidateSignUpEmailExist); !ok {
		return
	}

	// 检查数据库并返回响应
	response.JSON(c, gin.H{
		"exist": user.IsEmailExist(request.Email),
	})
}
