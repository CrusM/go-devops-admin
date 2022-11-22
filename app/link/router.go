package link

import (
	controller "go-devops-admin/app/link/controllers/v1"
	"go-devops-admin/middleware"

	"github.com/gin-gonic/gin"
)

func LinkRouterRegistryV1(v1 *gin.RouterGroup) {
	lsc := new(controller.LinksController)
	lscGroup := v1.Group("/links")
	{
		lscGroup.GET("", lsc.List)
		lscGroup.GET("/:id", lsc.Show)
		lscGroup.POST("", middleware.AuthJWT(), lsc.Create)
		lscGroup.PUT("/:id", middleware.AuthJWT(), lsc.Update)
		lscGroup.DELETE("/:id", middleware.AuthJWT(), lsc.Delete)
	}
}
