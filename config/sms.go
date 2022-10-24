package config

import "go-devops-admin/pkg/config"

func init() {
	config.Add("sms", func() map[string]interface{} {
		return map[string]interface{}{
			// 默认用 aliyun 短信平台
			"aliyun": map[string]interface{}{
				"access_key_id":     config.Env("sms.aliyun.ACCESS_ID"),
				"access_key_secret": config.Env("sms.aliyun.ACCESS_SECRET"),
				"sign_name":         config.Env("sms.aliyun.SIGN_NAME", "阿里云短信测试"),
				"template_code":     config.Env("sms.aliyun.TEMPLATE_CODE", "SMS_111111111"),
				"endpoint":          config.Env("sms.aliyun.ENDPOINT", "http://dysmsapi.aliyuncs.com/"),
			},
		}
	})
}
