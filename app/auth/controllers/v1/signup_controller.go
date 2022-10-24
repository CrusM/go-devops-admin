// 处理用户认证相关逻辑
package v1

import (
	"go-devops-admin/app"
	authRequest "go-devops-admin/app/auth/requests"
	user "go-devops-admin/app/user/models"
	"go-devops-admin/pkg/jwt"
	"go-devops-admin/pkg/response"

	"github.com/gin-gonic/gin"
)

// 注册控制器
type SignUpController struct {
	app.BaseAPIController
}

// 检查手机号是否被注册
func (sc *SignUpController) IsPhoneExist(c *gin.Context) {

	// 初始化请求对象
	request := authRequest.SignUpPhoneExistRequest{}

	// 校验请求数据
	if ok := app.Validate(c, &request, authRequest.ValidateSignUpPhoneExist); !ok {
		return
	}
	// 检查数据库并返回响应
	response.JSON(c, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}

// 检查邮箱是否被注册
func (sc *SignUpController) IsEmailExist(c *gin.Context) {
	request := authRequest.SignUpEmailExistRequest{}

	if ok := app.Validate(c, &request, authRequest.ValidateSignUpEmailExist); !ok {
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
	request := authRequest.SignUpUsingPhoneRequest{}
	if ok := app.Validate(c, &request, authRequest.SignUpUsingPhone); !ok {
		return
	}

	_user := user.User{
		Name:     request.Name,
		Phone:    request.Phone,
		Password: request.Password,
	}

	_user.Create()

	if _user.ID > 0 {
		token := jwt.NewJWT().IssueToken(_user.GetStringID(), _user.Name)
		response.CreatedJSON(c, gin.H{
			"data":  _user,
			"token": token,
		})
	} else {
		response.Abort500(c, "创建用户失败, 请稍后再试~")
	}
}

// 手机号注册请求数据结构体
// {
//     "name":"summer",
//     "password":"secret",
//     "password_confirm":"secret",
//     "verify_code": "{{verify_code_phone}}",
//     "email": "test@testing.com"
// }
func (sc *SignUpController) SignUpUsingEmail(c *gin.Context) {
	// 验证表单
	request := authRequest.SignUpUsingEmailRequest{}
	if ok := app.Validate(c, &request, authRequest.SignUpUsingEmail); !ok {
		return
	}

	_user := user.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	_user.Create()

	if _user.ID > 0 {
		token := jwt.NewJWT().IssueToken(_user.GetStringID(), _user.Name)
		response.CreatedJSON(c, gin.H{
			"data":  _user,
			"token": token,
		})
	} else {
		response.Abort500(c, "创建用户失败, 请稍后再试~")
	}
}
