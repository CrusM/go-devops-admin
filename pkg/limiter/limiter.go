package limiter

import (
	"go-devops-admin/pkg/config"
	"go-devops-admin/pkg/logger"
	"go-devops-admin/pkg/redis"
	"strings"

	"github.com/gin-gonic/gin"
	limiterLib "github.com/ulule/limiter/v3"
	sRedis "github.com/ulule/limiter/v3/drivers/store/redis"
)

// 处理限流逻辑

func GetKeyIP(c *gin.Context) string {
	return c.ClientIP()
}

func GetKeyRouteWithIP(c *gin.Context) string {
	return routeToKeyString(c.FullPath() + c.ClientIP())
}

func routeToKeyString(routeName string) string {
	routeName = strings.ReplaceAll(routeName, "/", "-")
	routeName = strings.ReplaceAll(routeName, ":", "_")
	return routeName
}

// 检测请求是否超额
func CheckRate(c *gin.Context, key string, formatted string) (limiterLib.Context, error) {
	// 实例化依赖的 limiter 包的 limiter.Rate 对象
	var context limiterLib.Context
	rate, err := limiterLib.NewRateFromFormatted(formatted)
	if err != nil {
		logger.LogIf(err)
		return context, err
	}

	// 初始化存储, 使用程序中公用的 Redis.Store 对象
	store, err := sRedis.NewStoreWithOptions(redis.Redis.Client, limiterLib.StoreOptions{
		// 设置 limiter 前缀, 保持 redis 中数据整洁
		Prefix: config.GetString("app.name") + ":limiter",
	})
	if err != nil {
		logger.LogIf(err)
		return context, err
	}

	//
	limiterObj := limiterLib.New(store, rate)

	// 获取限流结果
	if c.GetBool("limiter-once") {
		// Peek 取结果不增加访问次数
		return limiterObj.Peek(c, key)
	} else {
		// 确保多个路由组里面调用 limitIP 进行限流时, 只增加一次访问次数
		c.Set("limiter-once", true)
		// Get 取结果且增加一次访问次数
		return limiterObj.Get(c, key)
	}
}
