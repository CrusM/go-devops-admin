package v1

import (
	"go-devops-admin/app"
	authRequest "go-devops-admin/app/auth/requests"
	"go-devops-admin/pkg/auth"
	"go-devops-admin/pkg/jwt"
	"go-devops-admin/pkg/response"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	app.BaseAPIController
}

func (lc *LoginController) LoginByPhone(c *gin.Context) {

	request := authRequest.LoginByPhoneRequest{}
	if ok := app.Validate(c, &request, authRequest.LoginByPhone); !ok {
		return
	}

	// 登录
	user, err := auth.LoginByPhone(request.Phone)
	if err != nil {
		response.ERROR(c, err, "账号不存在")
	} else {
		// 登录成功
		token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Name)
		response.JSON(c, gin.H{
			"token": token,
		})
	}
}

func (lc *LoginController) LoginByPassword(c *gin.Context) {
	request := authRequest.LoginByPasswordRequest{}

	if ok := app.Validate(c, &request, authRequest.LoginByPassword); !ok {
		return
	}

	user, err := auth.Attempt(request.LoginID, request.Password)
	if err != nil {
		response.Unauthorized(c, "账号不存在或密码错误")
	} else {
		token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Name)
		response.JSON(c, gin.H{
			"token": token,
		})
	}

}

func (lc *LoginController) RefreshToken(c *gin.Context) {
	token, err := jwt.NewJWT().RefreshToken(c)

	if err != nil {
		response.ERROR(c, err, "token 刷新失败")
	} else {
		response.JSON(c, gin.H{
			"token": token,
		})
	}
}
