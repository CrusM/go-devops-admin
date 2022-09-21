package config

import "go-devops-admin/pkg/config"

// 初始化 redis 配置信息
func init() {
	config.Add("redis", func() map[string]interface{} {
		return map[string]interface{}{
			"host":     config.Env("REDIS_HOST", "127.0.0.1"),
			"port":     config.Env("REDIS_PORT", "6379"),
			"password": config.Env("REDIS_PASSWORD", ""),

			// 缓存 cache 包使用 0, 缓存清空理应当不影响业务
			"database_cache": config.Env("Redis_Cache_DB", 0),
			// 业务缓存用 1, 如 图片验证码，短信验证码
			"database": config.Env("Redis_MAIN_DB", 1),	
		}
	})
}
