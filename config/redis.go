package config

import "go-devops-admin/pkg/config"

// 初始化 redis 配置信息
func init() {
	config.Add("redis", func() map[string]interface{} {
		return map[string]interface{}{
			"host":     config.Env("redis.host", "127.0.0.1"),
			"port":     config.Env("redis.port", "6379"),
			"password": config.Env("redis.password", ""),

			// 缓存 cache 包使用 0, 缓存清空理应当不影响业务
			"database_cache": config.Env("redis.cache_db", 0),
			// 业务缓存用 1, 如 图片验证码，短信验证码
			"database": config.Env("redis.main_db", 1),
		}
	})
}
