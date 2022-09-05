package requests

import (
	"go-devops-admin/app/requests/validators"

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
	errs := validate(data, rules, messages)

	// 手机验证码
	_data := data.(*LoginByPhoneRequest)
	errs = validators.ValidateVerifyCode(_data.Phone, _data.Verify_code, errs)
	return errs
}
