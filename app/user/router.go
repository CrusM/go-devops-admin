package user

import (
	controller "go-devops-admin/app/user/controllers/v1"
	"go-devops-admin/middleware"

	"github.com/gin-gonic/gin"
)

func UserRouterRegistryV1(v1 *gin.RouterGroup) {
	uc := new(controller.UsersController)
	// v1.GET("/user", uc.CurrentUser)
	userGroup := v1.Group("user")
	{
		userGroup.GET("", middleware.AuthJWT(), uc.List)
		userGroup.PUT("/update-profile", middleware.AuthJWT(), uc.UpdateProfile)
		userGroup.PUT("/update-email", middleware.AuthJWT(), uc.UpdateUserEmail)
		userGroup.PUT("/update-phone", middleware.AuthJWT(), uc.UpdateUserPhone)
		userGroup.PUT("/update-password", middleware.AuthJWT(), uc.UpdateUserPassword)
		userGroup.PUT("/upload-avatar", middleware.AuthJWT(), uc.UpdateUserAvatar)
	}
}
