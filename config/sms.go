package config

import "go-devops-admin/pkg/config"

func init() {
	config.Add("sms", func() map[string]interface{} {
		return map[string]interface{}{
			// 默认用 aliyun 短信平台
			"aliyun": map[string]interface{}{
				"access_key_id":     config.Env("SMS_ALIYUN_ACCESS_ID"),
				"access_key_secret": config.Env("SMS_ALIYUN_ACCESS_SECRET"),
				"sign_name":         config.Env("SMS_ALIYUN_SIGN_NAME", "阿里云短信测试"),
				"template_code":     config.Env("SMS_ALIYUN_TEMPLATE_CODE", "SMS_111111111"),
				"endpoint":          config.Env("SMS_ALIYUN_ENDPOINT", "http://dysmsapi.aliyuncs.com/"),
			},
		}
	})
}
