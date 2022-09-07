package auth

import (
	"go-devops-admin/app/requests"
	"go-devops-admin/app/requests/validators"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type ResetPasswordByPhoneRequest struct {
	Phone      string `json:"phone,omitempty" valid:"phone"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
	Password   string `json:"password,omitempty" valid:"password"`
}

func ResetPasswordByPhone(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"phone":       []string{"required", "digits:11"},
		"verify_code": []string{"required", "digits:6"},
		"password":    []string{"required", "min:6"},
	}
	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项, 参数名 phone",
			"digits:手机号必须为 11 位长度的数字",
		},
		"verify_code": []string{
			"required:验证码必须填写",
			"digits:验证码必须是6位数字",
		},
		"password": []string{
			"required:密码为必填项,参数名password",
			"min:密码长度必须大于6",
		},
	}
	errs := requests.ValidateData(data, rules, messages)
	_data := data.(*ResetPasswordByPhoneRequest)
	return validators.ValidateVerifyCode(_data.Phone, _data.VerifyCode, errs)
}

type ResetPasswordByEmailRequest struct {
	Email      string `json:"email,omitempty" valid:"email"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
	Password   string `json:"password,omitempty" valid:"password"`
}

func ResetPasswordByEmail(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"email":       []string{"required", "min:4", "email"},
		"verify_code": []string{"required", "digits:6"},
		"password":    []string{"required", "min:6"},
	}
	messages := govalidator.MapData{
		"email": []string{
			"required:Email 为必填项, 参数名 email",
			"min:Email 长度需大于4	",
			"email:Email 格式不正确,请提供有效的邮箱地址",
		},
		"verify_code": []string{
			"required:验证码必须填写",
			"digits:验证码必须是6位数字",
		},
		"password": []string{
			"required:密码为必填项,参数名password",
			"min:密码长度必须大于6",
		},
	}
	errs := requests.ValidateData(data, rules, messages)
	_data := data.(*ResetPasswordByEmailRequest)
	return validators.ValidateVerifyCode(_data.Email, _data.VerifyCode, errs)
}
