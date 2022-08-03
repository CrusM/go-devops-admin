package config

import "go-devops-admin/pkg/config"

func init() {
	config.Add("captcha", func() map[string]interface{} {
		return map[string]interface{}{
			"height":    80,
			"width":     240,
			"length":    6,
			"max_skew":  0.7,
			"dot_count": 80,
			// 过期时间, 单位 分钟
			"expire_time": 15,
			// debug 模式下验证码过期时间, 方便调试
			"debug_expire_time": 10080,

			// 非生产环境用这个 key 跳过验证
			"test_key": "captcha_testing_key",
		}
	})
}
