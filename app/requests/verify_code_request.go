package requests

import (
	"go-devops-admin/pkg/captcha"

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

	errs := validate(data, rules, messages)

	// 图片验证码
	_data := data.(*VerifyCodePhoneRequest)
	if ok := captcha.NewCaptcha().VerifyCaptcha(_data.CaptchaID, _data.CaptchaAnswer); !ok {
		errs["captcha_answer"] = append(errs["captcha_answer"], "图片验证码错误")
	}

	return errs
}
