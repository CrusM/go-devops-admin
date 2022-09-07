package routes

import (
	"go-devops-admin/app/http/controllers/api/v1/auth"
	"go-devops-admin/app/http/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRouters(r *gin.Engine) {
	// 测试一个 v1 的路由组，所有v1版本的路由都存在这里
	v1 := r.Group("/v1")
	// 添加全局限流中间件: 每小时限速. 这里是所有 API (根据 IP) 请求加起来.
	// 参考 github api 每小时最多 60 个请求
	v1.Use(middleware.LimitIP("200-H"))
	{
		// 注册一个路由
		v1.GET("/", func(ctx *gin.Context) {
			// 返回JSON格式数据
			ctx.JSON(http.StatusOK, gin.H{
				"Hello": "World!",
			})
		})
		authGroup := v1.Group("auth")
		authGroup.Use(middleware.LimitIP("400-H"))
		{
			suc := new(auth.SignUpController)
			// 判断手机是否注册
			authGroup.POST("/signUp/phone/exist", middleware.GuestJWT(), middleware.LimitIP("60-H"), suc.IsPhoneExist)
			authGroup.POST("/signUp/email/exist", middleware.GuestJWT(), middleware.LimitIP("60-H"), suc.IsEmailExist)
			authGroup.POST("/signUp/using-phone", middleware.GuestJWT(), suc.SignUpUsingPhone)
			authGroup.POST("/signUp/using-email", middleware.GuestJWT(), suc.SignUpUsingEmail)

			// 验证码接口
			vcc := new(auth.VerifyCodeController)
			authGroup.POST("/verify-codes/captcha", vcc.ShowCaptcha)
			authGroup.POST("/verify-codes/phone", vcc.SendUsingPhone)
			authGroup.POST("/verify-codes/email", vcc.SendUsingEmail)

			// 登录接口
			lc := new(auth.LoginController)
			authGroup.POST("/login/using-phone", middleware.GuestJWT(), lc.LoginByPhone)
			authGroup.POST("/login/using-password", middleware.GuestJWT(), lc.LoginByPassword)
			authGroup.POST("/login/refresh-token", middleware.AuthJWT(), lc.RefreshToken)

			// 重置密码
			pwd := new(auth.PasswordController)
			authGroup.POST("reset-password/using-phone", middleware.GuestJWT(), pwd.ResetByPhone)
			authGroup.POST("reset-password/using-email", middleware.GuestJWT(), pwd.ResetByEmail)
		}
	}
}
