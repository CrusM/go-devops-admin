package config

import (
	"go-devops-admin/pkg/config"
)

func init() {
	config.Add("app", func() map[string]interface{} {
		return map[string]interface{}{
			// 应用名称
			"name": config.Env("APP_NAME", "devops"),
			// 当前环境，用以区分多环境，一般为 local，stage，production，test
			"env": config.Env("APP_ENV", "local"),
			// 是否DEBUG模式
			"debug": config.Env("APP_DEBUG", false),
			// 应用服务端口
			"port": config.Env("APP_PORT", 8080),
			// 加密回话，JWT加密
			"key": config.Env("APP_KEY", "33446a9dcf9ea060a0a6532b166da32f304af0de"),
			// 网站地址
			"url": config.Env("APP_URL", "http://localhost:8080"),
			// 设置时区
			"timezone": config.Env("TIMEZONE", "Asia/Shanghai"),
		}
	})
}
