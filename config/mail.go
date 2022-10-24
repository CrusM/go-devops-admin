package config

import "go-devops-admin/pkg/config"

func init() {
	config.Add("mail", func() map[string]interface{} {
		return map[string]interface{}{
			// 默认 mail 配置信息
			"smtp": map[string]interface{}{
				"host":     config.Env("mail.host", "localhost"),
				"port":     config.Env("mail.port", 1025),
				"username": config.Env("mail.username", ""),
				"password": config.Env("mail.password", ""),
			},

			"form": map[string]interface{}{
				"address": config.Env("mail.from.address", "devops@example.com"),
				"name":    config.Env("mail.from.name", "devops"),
			},
		}
	})
}
