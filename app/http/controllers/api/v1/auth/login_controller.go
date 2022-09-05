package auth

import (
	v1 "go-devops-admin/app/http/controllers/api/v1"
	"go-devops-admin/app/requests"
	"go-devops-admin/pkg/auth"
	"go-devops-admin/pkg/jwt"
	"go-devops-admin/pkg/response"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	v1.BaseAPIController
}

func (lc *LoginController) LoginByPhone(c *gin.Context) {

	request := requests.LoginByPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.LoginByPhone); !ok {
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
