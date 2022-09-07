package auth

import (
	v1 "go-devops-admin/app/http/controllers/api/v1"
	"go-devops-admin/app/models/user"
	"go-devops-admin/app/requests"
	authRequest "go-devops-admin/app/requests/auth"
	"go-devops-admin/pkg/response"

	"github.com/gin-gonic/gin"
)

type PasswordController struct {
	v1.BaseAPIController
}

func (pc *PasswordController) ResetByPhone(c *gin.Context) {
	request := authRequest.ResetPasswordByPhoneRequest{}
	if ok := requests.Validate(c, &request, authRequest.ResetPasswordByPhone); !ok {
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
