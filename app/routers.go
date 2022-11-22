package app

import (
	auth "go-devops-admin/app/auth"
	category "go-devops-admin/app/category"
	link "go-devops-admin/app/link"
	topic "go-devops-admin/app/topic"
	user "go-devops-admin/app/user"
	"go-devops-admin/middleware"

	"go-devops-admin/pkg/config"

	"github.com/gin-gonic/gin"
)

func getVersionGroup(r *gin.Engine, version string) *gin.RouterGroup {
	var g *gin.RouterGroup
	if len(config.Get("api.api_domain")) == 0 {
		g = r.Group("/api/" + version)
	} else {
		g = r.Group(version)
	}
	return g
}

func InitRegistryRouter(r *gin.Engine) {
	// v1
	v1 := getVersionGroup(r, "v1")
	v1.Use(middleware.LimitIP("200-H"))
	user.UserRouterRegistryV1(v1)
	auth.AuthRouterRegistryV1(v1)
	category.CategoryRouterRegistryV1(v1)
	topic.TopicRouterRegistryV1(v1)
	link.LinkRouterRegistryV1(v1)
}
