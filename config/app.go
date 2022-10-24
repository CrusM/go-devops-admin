package config

import (
	"go-devops-admin/pkg/config"
)

func init() {
	config.Add("app", func() map[string]interface{} {
		return map[string]interface{}{
			// 应用名称
			"name": config.Env("app.name", "devops"),
			// 当前环境，用以区分多环境，一般为 local，stage，production，test
			"env": config.Env("app.env", "local"),
			// 是否DEBUG模式
			"debug": config.Env("app.debug", false),
			// 应用服务端口
			"port": config.Env("app.port", 8080),
			// 加密回话，JWT加密
			"key": config.Env("app.key", "33446a9dcf9ea060a0a6532b166da32f304af0de"),
			// 网站地址
			"url": config.Env("app.url", "http://localhost:8080"),
			// 设置时区
			"timezone": config.Env("app.timezone", "Asia/Shanghai"),
		}
	})
}
