package requests

// 验证请求参数

import (
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

	return validate(data, rules, messages)
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

	return validate(data, rules, messages)
}
