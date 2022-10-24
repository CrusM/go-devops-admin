package requests

import (
	"go-devops-admin/app"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type LinkRequest struct {
	// request 字段
	Name string `json:"name,omitempty" valid:"name"`
	URL  string `json:"url,omitempty" valid:"url"`
}

func LinkSave(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"name": []string{"required", "max_cn:10", "min_cn:2"},
		"url":  []string{"required", "url"},
	}

	messages := govalidator.MapData{
		"name": []string{
			"required:name 为必填项",
			"min_cn:链接名称不能小于2个字",
			"min_cn:链接名称不能小于10个字",
		},
		"url": []string{
			"required:url 为必填项",
			"url:url 格式不正确",
		},
	}

	return app.ValidateData(data, rules, messages)
}
