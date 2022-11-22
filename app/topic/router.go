package topic

import (
	controller "go-devops-admin/app/topic/controllers/v1"
	"go-devops-admin/middleware"

	"github.com/gin-gonic/gin"
)

func TopicRouterRegistryV1(v1 *gin.RouterGroup) {
	tpc := new(controller.TopicsController)
	tpcGroup := v1.Group("/topics")
	{
		tpcGroup.GET("", tpc.List)
		tpcGroup.GET("/:id", tpc.Show)
		tpcGroup.POST("", middleware.AuthJWT(), tpc.Create)
		tpcGroup.PUT("/:id", middleware.AuthJWT(), tpc.Update)
		tpcGroup.DELETE("/:id", middleware.AuthJWT(), tpc.Delete)
	}
}
