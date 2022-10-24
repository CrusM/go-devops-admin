package auth

import (
	"errors"
	user "go-devops-admin/app/user/models"
	"go-devops-admin/pkg/logger"

	"github.com/gin-gonic/gin"
)

// 尝试登录
func Attempt(email string, password string) (user.User, error) {
	userModel := user.GetByMulti(email)
	if userModel.ID == 0 {
		return user.User{}, errors.New("账号不存在")
	}

	if !userModel.ComparePassword(password) {
		return user.User{}, errors.New("密码错误")
	}

	return userModel, nil
}

// 手机号登录
func LoginByPhone(phone string) (user.User, error) {
	userModel := user.GetByPhone(phone)
	if userModel.ID == 0 {
		return user.User{}, errors.New("账号不存在")
	}
	return userModel, nil
}

// 从 gin.Context 中获取当前登录用户
func CurrentUser(c *gin.Context) user.User {
	userModel, ok := c.MustGet("current_user").(user.User)
	if !ok {
		logger.LogIf(errors.New("无法获取用户"))
		return user.User{}
	}
	return userModel
}

// 从 gin.Context 中获取当前登录用户 ID
func CurrentUID(c *gin.Context) string {
	return c.GetString("current_user_id")
}
