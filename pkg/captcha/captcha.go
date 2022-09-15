package captcha

import (
	"go-devops-admin/pkg/app"
	"go-devops-admin/pkg/config"
	"go-devops-admin/pkg/redis"
	"sync"

	"github.com/mojocn/base64Captcha"
)

// 处理图片验证码逻辑

type Captcha struct {
	Base64Captcha *base64Captcha.Captcha
}

// 确保 internalCaptcha 对象只初始化一次
var once sync.Once

// 内部使用 Captcha 对象
var internalCaptcha *Captcha

// 单例获取模式
func NewCaptcha() *Captcha {
	once.Do(func() {
		// 初始化 Captcha 对象
		internalCaptcha = &Captcha{}

		// 使用全局 redis 对象, 并配置 key 前缀
		store := RedisStore{
			RedisClient: redis.Redis,
			KeyPrefix:   config.GetString("app.name") + ":captcha:",
		}

		// 配置 base64Captcha 驱动信息
		device := base64Captcha.NewDriverDigit(
			config.GetInt("captcha.height"),
			config.GetInt("captcha.width"),
			config.GetInt("captcha.length"),
			config.GetFloat64("captcha.max_skew"), // 图片最大斜角
			config.GetInt("captcha.dot_count"),    // 图片背景混淆点数
		)

		// 实例化 base64Captcha 并赋值给内部使用的 internalCaptcha 对象
		internalCaptcha.Base64Captcha = base64Captcha.NewCaptcha(device, &store)
	})
	return internalCaptcha
}

// 生成图片验证码
func (c *Captcha) GenerateCaptcha() (id string, b64s string, err error) {
	return c.Base64Captcha.Generate()
}

// 验证验证码是否正确
func (c *Captcha) VerifyCaptcha(id string, answer string) (match bool) {
	// 方便本地和 API 自动测试
	if !app.IsProduction() && id == config.GetString("captcha.testing_key") {
		return true
	}
	// 第三个参数是验证后是否删除, 使用 false
	// 用户多次提交时, 防止表单错误
	return c.Base64Captcha.Verify(id, answer, false)
}
