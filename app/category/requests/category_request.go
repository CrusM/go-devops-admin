package requests

import (
	"go-devops-admin/app/base"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type CategoryRequest struct {
	// request 字段
	Name        string `json:"name" valid:"name"`
	Description string `json:"description,omitempty" valid:"description"`
}

func CategorySave(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"name":        []string{"required", "min_cn:2", "max_cn:8", "not_exists:categories,name"},
		"description": []string{"min_cn:3", "max_cn:255"},
	}

	messages := govalidator.MapData{
		"name": []string{
			"required:name 为必填项",
			"min_cn:分类名称长度至少 2 个字",
			"max_cn:分类名称长度最多 8 个字",
			"not_exists:分类名称已经存在",
		},
		"description": []string{
			"min_cn:分类描述最少 3 个字",
			"max_cn:分类描述最多 255 个字",
		},
	}

	return base.ValidateData(data, rules, messages)
}
