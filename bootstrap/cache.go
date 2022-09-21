package bootstrap

import (
	"fmt"
	"go-devops-admin/pkg/cache"
	"go-devops-admin/pkg/config"
)

func SetupCache() {
	// 初始化缓存专用的 redis client
	rds := cache.NewRedisStore(
		fmt.Sprintf("%v:%v", config.GetString("redis.host"), config.GetString("redis.port")),
		config.GetString("redis.username"),
		config.GetString("redis.password"),
		config.GetInt("redis.database_cache"),
	)

	cache.InitWithCacheStore(rds)
}
