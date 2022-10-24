package routes

import (
	auth "go-devops-admin/app/auth/controllers/v1"
	category "go-devops-admin/app/category/controllers/v1"
	link "go-devops-admin/app/link/controllers/v1"
	topic "go-devops-admin/app/topic/controllers/v1"
	user "go-devops-admin/app/user/controllers/v1"
	"go-devops-admin/middleware"
	"go-devops-admin/pkg/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRouters(r *gin.Engine) {
	// 测试一个 v1 的路由组，所有v1版本的路由都存在这里
	var v1 *gin.RouterGroup
	if len(config.Get("api.api_domain")) == 0 {
		v1 = r.Group("/api/v1")
	} else {
		v1 = r.Group("/v1")
	}
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
			// 注册接口
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
		// user 接口
		uc := new(user.UsersController)
		// 获取当前用户
		v1.GET("/user", uc.CurrentUser)
		userGroup := v1.Group("/users")
		{
			userGroup.GET("", middleware.AuthJWT(), uc.List)
			userGroup.PUT("/update-profile", middleware.AuthJWT(), uc.UpdateProfile)
			userGroup.PUT("/update-email", middleware.AuthJWT(), uc.UpdateUserEmail)
			userGroup.PUT("/update-phone", middleware.AuthJWT(), uc.UpdateUserPhone)
			userGroup.PUT("/update-password", middleware.AuthJWT(), uc.UpdateUserPassword)
			userGroup.PUT("/upload-avatar", middleware.AuthJWT(), uc.UpdateUserAvatar)
		}

		// category
		cgc := new(category.CategoriesController)
		cgcGroup := v1.Group("/categories")
		{
			cgcGroup.GET("", cgc.List)
			cgcGroup.POST("", middleware.AuthJWT(), cgc.Create)
			cgcGroup.PUT("/:id", middleware.AuthJWT(), cgc.Update)
			cgcGroup.DELETE("/:id", middleware.AuthJWT(), cgc.Delete)
		}

		// topic
		tpc := new(topic.TopicsController)
		tpcGroup := v1.Group("/topics")
		{
			tpcGroup.GET("", tpc.List)
			tpcGroup.GET("/:id", tpc.Show)
			tpcGroup.POST("", middleware.AuthJWT(), tpc.Create)
			tpcGroup.PUT("/:id", middleware.AuthJWT(), tpc.Update)
			tpcGroup.DELETE("/:id", middleware.AuthJWT(), tpc.Delete)
		}

		// topic
		lsc := new(link.LinksController)
		lscGroup := v1.Group("/links")
		{
			lscGroup.GET("", lsc.List)
			lscGroup.GET("/:id", lsc.Show)
			lscGroup.POST("", middleware.AuthJWT(), lsc.Create)
			lscGroup.PUT("/:id", middleware.AuthJWT(), lsc.Update)
			lscGroup.DELETE("/:id", middleware.AuthJWT(), lsc.Delete)
		}
	}
}
