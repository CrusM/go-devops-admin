package routes

import (
	"go-devops-admin/app/http/controllers/api/v1/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRouters(r *gin.Engine) {
	// 测试一个 v1 的路由组，所有v1版本的路由都存在这里
	v1 := r.Group("/v1")
	{
		// 注册一个路由
		v1.GET("/", func(ctx *gin.Context) {
			// 返回JSON格式数据
			ctx.JSON(http.StatusOK, gin.H{
				"Hello": "World!",
			})
		})
		authGroup := v1.Group("auth")
		{
			suc := new(auth.SignUpController)
			// 判断手机是否注册
			authGroup.POST("/signUp/phone/exist", suc.IsPhoneExist)
			authGroup.POST("/signUp/email/exist", suc.IsEmailExist)
			authGroup.POST("/signUp/using-phone", suc.SignUpUsingPhone)
			authGroup.POST("/signUp/using-email", suc.SignUpUsingEmail)

			// 验证码接口
			vcc := new(auth.VerifyCodeController)
			authGroup.POST("/verify-codes/captcha", vcc.ShowCaptcha)
			authGroup.POST("/verify-codes/phone", vcc.SendUsingPhone)
			authGroup.POST("/verify-codes/email", vcc.SendUsingEmail)

			// 登录接口
			lc := new(auth.LoginController)
			authGroup.POST("/login/using-phone", lc.LoginByPhone)
			authGroup.POST("/login/using-password", lc.LoginByPassword)
			authGroup.POST("/login/refresh-token", lc.RefreshToken)

			// 重置密码
			pwd := new(auth.PasswordController)
			authGroup.POST("reset-password/using-phone", pwd.ResetByPhone)
			authGroup.POST("reset-password/using-email", pwd.ResetByEmail)
		}
	}
}
