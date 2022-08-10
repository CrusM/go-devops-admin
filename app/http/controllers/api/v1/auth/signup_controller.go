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

// 手机号注册请求数据结构体
// {
//     "name":"summer",
//     "password":"secret",
//     "password_confirm":"secret",
//     "verify_code": "{{verify_code_phone}}",
//     "phone": "00000000000"
// }
func (sc *SignUpController) SignUpUsingPhone(c *gin.Context) {
	// 验证表单
	request := requests.SignUpUsingPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.SignUpUsingPhone); !ok {
		return
	}

	_user := user.User{
		Name:     request.Name,
		Phone:    request.Phone,
		Password: request.Password,
	}

	_user.Create()

	if _user.ID > 0 {
		response.CreatedJSON(c, gin.H{
			"data": _user,
		})
	} else {
		response.Abort500(c, "创建用户失败, 请稍后再试~")
	}
}
