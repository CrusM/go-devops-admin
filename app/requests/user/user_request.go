package user

import (
	"go-devops-admin/app/requests"
	"go-devops-admin/pkg/auth"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type UserRequest struct {
	// request 字段
	// Name string `json:"name,omitempty" valid:"name"`
}

func UserSave(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		// "name": []string{"required"},
	}

	messages := govalidator.MapData{
		// "name": []string{"required:name 为必填项"},
	}

	return requests.ValidateData(data, rules, messages)
}

type UserUpdateProfileRequest struct {
	// request 字段
	Name         string `json:"name" valid:"name"`
	City         string `json:"city" valid:"city"`
	Introduction string `json:"introduction " valid:"introduction"`
}

func UserUpdateProfile(data interface{}, c *gin.Context) map[string][]string {
	uid := auth.CurrentUID(c)
	rules := govalidator.MapData{
		"name":         []string{"required", "alpha_num", "between:3,20", "not_exists:user,name," + uid},
		"city":         []string{"min_cn:2", "max_cn:20"},
		"introduction": []string{"min_cn:4", "max_cn:240"},
	}

	messages := govalidator.MapData{
		"name": []string{
			"required:用户名必须填写",
			"alpha_num:用户名格式错误,只允许数字和英文",
			"between:用户名长度需在 3~20 之间",
			"not_exists:用户名已经被占用",
		},
		"city": []string{
			"min_cn:城市名称需大于2个字",
			"max_cn:城市名称需大于20个字",
		},
		"introduction": []string{
			"min_cn:描述长度需大于4个字",
			"max_cn:描述长度需小于240个字",
		},
	}

	return requests.ValidateData(data, rules, messages)
}


