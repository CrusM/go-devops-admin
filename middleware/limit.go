package middleware

import (
	"go-devops-admin/pkg/app"
	"go-devops-admin/pkg/limiter"
	"go-devops-admin/pkg/logger"
	"go-devops-admin/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

// LimitIP 全局限流中间件, 针对 IP 进行限流
// limit 为格式化字符串, 如 5-S, 示例:
// * 5 req / seconds: "5-S"
// * 10 req / minutes: "5-M"
// * 1000 req / hours: "1000-H"
// * 2000 req / days: "2000-D"
//
func LimitIP(limit string) gin.HandlerFunc {
	if app.IsTesting() {
		limit = "1000000-H"
	}

	return func(c *gin.Context) {
		key := limiter.GetKeyIP(c)
		if ok := limitHandler(c, key, limit); !ok {
			return
		}
		c.Next()
	}
}

func LimitPreRoute(limit string) gin.HandlerFunc {
	if app.IsTesting() {
		limit = "1000000-H"
	}
	return func(c *gin.Context) {
		// 针对单个路由增加访问次数
		c.Set("limiter-once", true)
		// 针对 IP + Route 限流
		key := limiter.GetKeyRouteWithIP(c)
		if ok := limitHandler(c, key, limit); !ok {
			return
		}
		c.Next()
	}
}

func limitHandler(c *gin.Context, key string, limit string) bool {
	// 获取超额情况
	rate, err := limiter.CheckRate(c, key, limit)
	if err != nil {
		logger.LogIf(err)
		response.Abort500(c)
		return false
	}

	// 设置 header 信息
	// X-RateLimit-Limit :10000 最大访问次数
	// X-RateLimit-Remaining :9993 剩余的访问次数
	// X-RateLimit-Reset :1513784506 到该时间点，访问次数会重置为 X-RateLimit-Limit
	c.Header("X-RateLimit-Limit", cast.ToString(rate.Limit))
	c.Header("X-RateLimit-Remaining", cast.ToString(rate.Remaining))
	c.Header("X-RateLimit-Reset", cast.ToString(rate.Reset))

	// 超额
	if rate.Reached {
		// 提示用户被限流
		c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
			"message": "接口请求太频繁, 请稍后再试",
		})
		return false
	}

	return true
}
