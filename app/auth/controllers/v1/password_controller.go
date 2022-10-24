package v1

import (
	"go-devops-admin/app"
	authRequest "go-devops-admin/app/auth/requests"
	user "go-devops-admin/app/user/models"
	"go-devops-admin/pkg/response"

	"github.com/gin-gonic/gin"
)

type PasswordController struct {
	app.BaseAPIController
}

func (pc *PasswordController) ResetByPhone(c *gin.Context) {
	request := authRequest.ResetPasswordByPhoneRequest{}
	if ok := app.Validate(c, &request, authRequest.ResetPasswordByPhone); !ok {
		return
	}

	// 更新密码
	userModel := user.GetByPhone(request.Phone)
	if userModel.ID == 0 {
		response.Abort404(c)
	} else {
		userModel.Password = request.Password
		userModel.Save()
		response.SUCCESS(c)
	}

}

func (pc *PasswordController) ResetByEmail(c *gin.Context) {
	request := authRequest.ResetPasswordByEmailRequest{}
	if ok := app.Validate(c, &request, authRequest.ResetPasswordByEmail); !ok {
		return
	}

	// 更新密码
	userModel := user.GetByEmail(request.Email)
	if userModel.ID == 0 {
		response.Abort404(c)
	} else {
		userModel.Password = request.Password
		userModel.Save()
		response.SUCCESS(c)
	}

}
