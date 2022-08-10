package config

import "go-devops-admin/pkg/config"

func init() {
	config.Add("verifyCode", func() map[string]interface{} {
		return map[string]interface{}{
			// 验证码长度
			"code_length": config.Env("VERIFY_CODE_LENGTH", 6),
			// 过期时间, 单位分钟
			"expire_time": config.Env("VERIFY_CODE_EXPIRE", 15),

			// debug 模式
			"debug_expire_time": 10080,
			"debug_code":        123456,

			"debug_phone_prefix": "000",
			"debug_email_suffix": "@testing.com",
		}
	})
}
