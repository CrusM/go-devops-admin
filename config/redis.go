package config

import "go-devops-admin/pkg/config"

// 初始化 redis 配置信息
func init() {
	config.Add("redis", func() map[string]interface{} {
		return map[string]interface{}{
			"host":     config.Env("REDIS_HOST", "127.0.0.1"),
			"port":     config.Env("REDIS_PORT", "6379"),
			"password": config.Env("REDIS_PASSWORD", ""),

			// 数据库
			"database": config.Env("Redis_MAIN_DB", 1),
		}
	})
}
