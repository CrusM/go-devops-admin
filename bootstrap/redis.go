package bootstrap

import (
	"fmt"
	"go-devops-admin/pkg/config"
	"go-devops-admin/pkg/redis"
)

// 初始化 redis
func SetupRedis() {
	// 建立 redis 连接
	redis.ConnectRedis(
		fmt.Sprintf("%v:%v", config.GetString("redis.host"), config.GetString("redis.port")),
		config.GetString("redis.username"),
		config.GetString("redis.password"),
		config.GetInt("redis.database"),
	)
}
