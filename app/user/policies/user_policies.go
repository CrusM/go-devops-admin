package user

import (
	user "go-devops-admin/app/user/models"
	"go-devops-admin/pkg/auth"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func CanModifyUser(c *gin.Context, _user user.User) bool {
	return auth.CurrentUID(c) == cast.ToString(_user.ID)
}
