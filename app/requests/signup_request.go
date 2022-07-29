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

	// 初始化配置
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		Messages:      messages,
		TagIdentifier: "valid", // 模型的 struct 标签标识符
	}

	// 开始验证
	return govalidator.New(opts).ValidateStruct()
}
