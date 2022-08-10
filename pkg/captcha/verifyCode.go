package captcha

import (
	"fmt"
	"go-devops-admin/pkg/app"
	"go-devops-admin/pkg/config"
	"go-devops-admin/pkg/helpers"
	"go-devops-admin/pkg/logger"
	"go-devops-admin/pkg/mail"
	"go-devops-admin/pkg/redis"
	"go-devops-admin/pkg/sms"
	"strings"
)

type VerifyCode struct {
	Store Store
}

// 单例操作
// var once sync.Once
var internalVerifyCode *VerifyCode

func NewVerifyCode() *VerifyCode {
	once.Do(func() {
		internalVerifyCode = &VerifyCode{
			Store: &RedisStore{
				RedisClient: redis.Redis,
				KeyPrefix:   config.GetString("app.name") + ":verifyCode:",
			},
		}
	})
	return internalVerifyCode
}

func (vc *VerifyCode) SendSMS(phone string) bool {
	// 生成验证码
	code := vc.generateVerifyCode(phone)
	// 本地调试代码
	if !app.IsProduction() && strings.HasPrefix(phone, config.GetString("verifyCode.debug_phone_prefix")) {
		return true
	}

	// 发送短信
	return sms.NewSMS().Send(phone, sms.Message{
		Template: config.GetString("sms.aliyun.template_code"),
		Data:     map[string]string{"code": code},
	})
}

// 发送邮件验证码
func (vc *VerifyCode) SendEmail(email string) error {
	// 生成验证码
	code := vc.generateVerifyCode(email)
	if !app.IsProduction() && strings.HasSuffix(email, config.GetString("verifyCode.debug_email_suffix")) {
		return nil
	}

	content := fmt.Sprintf("<h1>您的 Email 验证码是: %v", code)
	// 发送邮件
	mail.NewMailer().Send(mail.Email{
		From: mail.From{
			Address: config.GetString("mail.from.address"),
			Name:    config.GetString("mail.from.name"),
		},
		To:      []string{email},
		Subject: "Email 验证码",
		HTML:    []byte(content),
	})
	return nil

}

// 检查验证码是否正确
func (vc *VerifyCode) CheckAnswer(key string, answer string) bool {
	logger.DebugJson("验证码", "检查验证码", map[string]string{key: answer})

	// 非 生产环境 调试
	if !app.IsProduction() && (strings.HasSuffix(key, config.GetString("verifyCode.debug_email_suffix")) || strings.HasPrefix(key, config.GetString("verifyCode.debug_phone_prefix"))) {
		return true
	}
	return vc.Store.Verify(key, answer, false)
}

func (vc *VerifyCode) generateVerifyCode(key string) string {
	// 生成随机验证码
	code := helpers.RandomNumber(config.GetInt("verifyCode.code_length"))

	// 本地调试
	if app.IsLocal() {
		code = config.GetString("verifyCode.debug_code")
	}

	logger.DebugJson("验证码", "生成验证码", map[string]string{key: code})
	// 将验证码存入 redis, 并设置过期时间
	vc.Store.Set(key, code)
	return ""
}
