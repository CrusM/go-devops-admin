package config

import "go-devops-admin/pkg/config"

func init() {
	config.Add("mail", func() map[string]interface{} {
		return map[string]interface{}{
			// 默认 mail 配置信息
			"smtp": map[string]interface{}{
				"host":     config.Env("MAIL_HOST", "localhost"),
				"port":     config.Env("MAIL_PORT", 1025),
				"username": config.Env("MAIL_USERNAME", ""),
				"password": config.Env("MAIL_PASSWORD", ""),
			},

			"form": map[string]interface{}{
				"address": config.Env("EMAIL_FROM_ADDRESS", "devops@example.com"),
				"name":    config.Env("EMAIL_FROM_NAME", "devops"),
			},
		}
	})
}
