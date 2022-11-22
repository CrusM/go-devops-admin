package category

import (
	controller "go-devops-admin/app/category/controllers/v1"
	"go-devops-admin/middleware"

	"github.com/gin-gonic/gin"
)

func CategoryRouterRegistryV1(v1 *gin.RouterGroup) {
	cgc := new(controller.CategoriesController)
	cgcGroup := v1.Group("/categories")
	{
		cgcGroup.GET("", cgc.List)
		cgcGroup.POST("", middleware.AuthJWT(), cgc.Create)
		cgcGroup.PUT("/:id", middleware.AuthJWT(), cgc.Update)
		cgcGroup.DELETE("/:id", middleware.AuthJWT(), cgc.Delete)
	}
}
