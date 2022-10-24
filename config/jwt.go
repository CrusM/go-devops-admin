package config

import "go-devops-admin/pkg/config"

func init() {
	config.Add("jwt", func() map[string]interface{} {
		return map[string]interface{}{
			// 过期时间
			"expire_time": config.Env("jwt.expire_time", 120),

			// 最大允许刷新时间
			"max_refresh_time": config.Env("jwt.max_refresh_time", 86400),

			// 调试模式过期时间
			"debug_expire_time": config.Env("jwt.debug_expire_time", 86400),
		}
	})
}
