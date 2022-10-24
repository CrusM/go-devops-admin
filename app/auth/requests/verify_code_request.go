package requests

import (
	"go-devops-admin/app"
	"go-devops-admin/app/validators"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type VerifyCodePhoneRequest struct {
	CaptchaID     string `json:"captcha_id,omitempty" valid:"captcha_id"`
	CaptchaAnswer string `json:"captcha_answer,omitempty" valid:"captcha_answer"`

	Phone string `json:"phone,omitempty" valid:"phone"`
}

// 验证表单, 返回长度为 0 即验证通过
func VerifyCodePhone(data interface{}, c *gin.Context) map[string][]string {
	// 验证规则
	rules := govalidator.MapData{
		"phone":          []string{"required", "digits:11"},
		"captcha_id":     []string{"required"},
		"captcha_answer": []string{"required", "digits:6"},
	}

	// 验证返回消息
	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号必须填写，参数名 phone",
			"digits:请填入11位手机号码",
		},
		"captcha_id": []string{
			"required:图片验证码 ID 为必填项",
		},
		"captcha_answer": []string{
			"required: 图片验证码答案为必填项",
			"digits:图片验证码长度为 6 位数字",
		},
	}

	errs := app.ValidateData(data, rules, messages)

	// 图片验证码
	_data := data.(*VerifyCodePhoneRequest)

	return validators.ValidateCaptcha(_data.CaptchaID, _data.CaptchaAnswer, errs)
}

// 邮件表单验证
type VerifyCodeEmailRequest struct {
	CaptchaID     string `json:"captcha_id,omitempty" valid:"captcha_id"`
	CaptchaAnswer string `json:"captcha_answer,omitempty" valid:"captcha_answer"`

	Email string `json:"email,omitempty" valid:"email"`
}

func VerifyCodeEmail(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"email":          []string{"required", "email"},
		"captcha_id":     []string{"required"},
		"captcha_answer": []string{"required", "digits:6"},
	}
	messages := govalidator.MapData{
		"email": []string{
			"required:Email 为必填项",
			"email: Email 格式不正确, 请提供正确的邮箱地址",
		},
		"captcha_id": []string{
			"required:图片验证码 ID 为必填项",
		},
		"captcha_answer": []string{
			"required: 图片验证码答案为必填项",
			"digits:图片验证码长度为 6 位数字",
		},
	}
	errs := app.ValidateData(data, rules, messages)

	// 图片验证码
	_data := data.(*VerifyCodeEmailRequest)
	return validators.ValidateCaptcha(_data.CaptchaID, _data.CaptchaAnswer, errs)
}
