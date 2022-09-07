package auth

// 验证请求参数

import (
	"go-devops-admin/app/requests"
	"go-devops-admin/app/requests/validators"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type SignUpPhoneExistRequest struct {
	Phone string `json:"phone,omitempty" valid:"phone"`
}

func ValidateSignUpPhoneExist(data interface{}, c *gin.Context) map[string][]string {
	// 自定义验证规则
	rules := govalidator.MapData{
		"phone": []string{"required", "digits:11"},
	}

	// 自定义验证出错时提示
	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项,参数名phone",
			"digits:手机号长度为11位的数字",
		},
	}

	return requests.ValidateData(data, rules, messages)
}

type SignUpEmailExistRequest struct {
	Email string `json:"email,omitempty" valid:"email"`
}

func ValidateSignUpEmailExist(data interface{}, c *gin.Context) map[string][]string {
	// 定义验证规则
	rules := govalidator.MapData{
		"email": []string{"required", "min:4", "max:30", "email"},
	}
	// 定义验证错误提示
	messages := govalidator.MapData{
		"email": []string{
			"required:Email 为必填项,字段名 email",
			"min:Email 长度不能小于4",
			"max:Email 长度不能大于30",
			"email:Email 格式不正确,请提供有效的Email",
		},
	}

	return requests.ValidateData(data, rules, messages)
}

// 通过手机注册的请求信息
type SignUpUsingPhoneRequest struct {
	Phone           string `json:"phone,omitempty" valid:"phone"`
	Name            string `json:"name" valid:"name"`
	VerifyCode      string `json:"verify_code,omitempty" valid:"verify_code"`
	Password        string `json:"password,omitempty" valid:"password"`
	PasswordConfirm string `json:"password_confirm,omitempty" valid:"password_confirm"`
}

func SignUpUsingPhone(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"phone":            []string{"required", "digits:11", "not_exists:users,phone"},
		"name":             []string{"required", "alpha_num", "between:3,20", "not_exists:users,name"},
		"password":         []string{"required", "min:6"},
		"password_confirm": []string{"required"},
		"verify_code":      []string{"required", "digits:6"},
	}
	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号码为必填项, 参数名 phone",
			"digits:手机号长度为11位的数字",
		},
		"name": []string{
			"required:用户名必须填写",
			"alpha_num:用户名格式错误, 只允许数字和英文",
			"between:用户名长度需在3~20个字符之间",
		},
		"password": []string{
			"required:密码必须填写",
			"min:密码长度需大于6",
		},
		"password_confirm": []string{
			"required:确认密码必须填写",
		},
		"verify_code": []string{
			"required:验证码答案为必填项",
			"digits:验证码长度为 6 位数字",
		},
	}

	errs := requests.ValidateData(data, rules, messages)
	_data := data.(*SignUpUsingPhoneRequest)
	errs = validators.ValidatePasswordConfirm(_data.Password, _data.PasswordConfirm, errs)
	errs = validators.ValidateVerifyCode(_data.Phone, _data.VerifyCode, errs)
	return errs
}

// 通过 Email 注册账号
type SignUpUsingEmailRequest struct {
	Email           string `json:"email,omitempty" valid:"email"`
	Name            string `json:"name" valid:"name"`
	VerifyCode      string `json:"verify_code,omitempty" valid:"verify_code"`
	Password        string `json:"password,omitempty" valid:"password"`
	PasswordConfirm string `json:"password_confirm,omitempty" valid:"password_confirm"`
}

func SignUpUsingEmail(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"email":            []string{"required", "email", "not_exists:users,email"},
		"name":             []string{"required", "alpha_num", "between:3,20", "not_exists:users,name"},
		"password":         []string{"required", "min:6"},
		"password_confirm": []string{"required"},
		"verify_code":      []string{"required", "digits:6"},
	}
	messages := govalidator.MapData{
		"email": []string{
			"required:邮箱地址必填项, 参数名 email",
			"email:邮箱格式不正确, 请输入正确的邮箱地址",
		},
		"name": []string{
			"required:用户名必须填写",
			"alpha_num:用户名格式错误, 只允许数字和英文",
			"between:用户名长度需在3~20个字符之间",
		},
		"password": []string{
			"required:密码必须填写",
			"min:密码长度需大于6",
		},
		"password_confirm": []string{
			"required:确认密码必须填写",
		},
		"verify_code": []string{
			"required:验证码答案为必填项",
			"digits:验证码长度为 6 位数字",
		},
	}

	errs := requests.ValidateData(data, rules, messages)
	_data := data.(*SignUpUsingEmailRequest)
	errs = validators.ValidatePasswordConfirm(_data.Password, _data.PasswordConfirm, errs)
	errs = validators.ValidateVerifyCode(_data.Email, _data.VerifyCode, errs)
	return errs
}
