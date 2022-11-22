package requests

import (
	"go-devops-admin/app/base"
	"go-devops-admin/app/base/validators"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type LoginByPhoneRequest struct {
	Phone       string `json:"phone,omitempty" valid:"phone"`
	Verify_code string `json:"verify_code,omitempty" valid:"verify_code"`
}

// 表单验证
func LoginByPhone(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"phone":       []string{"required", "digits:11"},
		"verify_code": []string{"required", "digits:6"},
	}

	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项, 参数名 phone",
			"digits:手机长度必须为11位数字",
		},
		"verify_code": []string{
			"required:验证码答案必填",
			"digits:验证码长度必须为6为数字",
		},
	}
	errs := base.ValidateData(data, rules, messages)

	// 手机验证码
	_data := data.(*LoginByPhoneRequest)
	errs = validators.ValidateVerifyCode(_data.Phone, _data.Verify_code, errs)
	return errs
}

// 用户密码登录
type LoginByPasswordRequest struct {
	CaptchaID     string `json:"captcha_id,omitempty" valid:"captcha_id"`
	CaptchaAnswer string `json:"captcha_answer,omitempty" valid:"captcha_answer"`
	LoginID       string `json:"login_id,omitempty" valid:"login_id"`
	Password      string `json:"password,omitempty" valid:"password"`
}

func LoginByPassword(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"captcha_id":     []string{"required"},
		"captcha_answer": []string{"required", "digits:6"},
		"login_id":       []string{"required", "min:3"},
		"password":       []string{"required", "min:6"},
	}

	messages := govalidator.MapData{
		"captcha_id": []string{
			"required: 图片验证码的 ID 为必填项",
		},
		"captcha_answer": []string{
			"required: 图片验证码答案为必填项",
			"digits:图片验证码答案长度必须为 6 位数字",
		},
		"login_id": []string{
			"required: 登录 ID 为必填项, 支持手机号、邮箱 和 用户名",
			"min:登录 ID 长度需大于 3",
		},
		"password": []string{
			"required: 密码为必填项",
			"min:密码长度需大于 6",
		},
	}

	errs := base.ValidateData(data, rules, messages)
	_data := data.(*LoginByPasswordRequest)
	return validators.ValidateCaptcha(_data.CaptchaID, _data.CaptchaAnswer, errs)
}
