package auth

import (
	"fmt"
	controller "go-devops-admin/app/auth/controllers/v1"
	"go-devops-admin/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRouterRegistryV1(v1 *gin.RouterGroup) {
	fmt.Printf("%v\n", v1)
	authGroup := v1.Group("auth")
	authGroup.Use(middleware.LimitIP("400-H"))
	{
		suc := new(controller.SignUpController)
		// 注册接口
		authGroup.POST("/signUp/phone/exist", middleware.GuestJWT(), middleware.LimitIP("60-H"), suc.IsPhoneExist)
		authGroup.POST("/signUp/email/exist", middleware.GuestJWT(), middleware.LimitIP("60-H"), suc.IsEmailExist)
		authGroup.POST("/signUp/using-phone", middleware.GuestJWT(), suc.SignUpUsingPhone)
		authGroup.POST("/signUp/using-email", middleware.GuestJWT(), suc.SignUpUsingEmail)

		// 验证码接口
		vcc := new(controller.VerifyCodeController)
		authGroup.POST("/verify-codes/captcha", vcc.ShowCaptcha)
		authGroup.POST("/verify-codes/phone", vcc.SendUsingPhone)
		authGroup.POST("/verify-codes/email", vcc.SendUsingEmail)

		// 登录接口
		lc := new(controller.LoginController)
		authGroup.POST("/login/using-phone", middleware.GuestJWT(), lc.LoginByPhone)
		authGroup.POST("/login/using-password", middleware.GuestJWT(), lc.LoginByPassword)
		authGroup.POST("/login/refresh-token", middleware.AuthJWT(), lc.RefreshToken)

		// 重置密码
		pwd := new(controller.PasswordController)
		authGroup.POST("reset-password/using-phone", middleware.GuestJWT(), pwd.ResetByPhone)
		authGroup.POST("reset-password/using-email", middleware.GuestJWT(), pwd.ResetByEmail)

	}
}
