package requests

import (
	"go-devops-admin/app/base"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type TopicRequest struct {
	// request 字段
	// Name string `json:"name,omitempty" valid:"name"`
	Title      string `json:"title,omitempty" valid:"title"`
	Body       string `json:"body,omitempty" valid:"body"`
	CategoryID string `json:"category_id,omitempty" valid:"category_id"`
}

func TopicSave(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"title":       []string{"required", "min_cn:3", "max_cn:40"},
		"body":        []string{"required", "min_cn:10", "max_cn:50000"},
		"category_id": []string{"required", "exists:categories,id"},
	}

	messages := govalidator.MapData{
		"title": []string{
			"required:帖子标题必填项, 参数名 title",
			"min_cn:帖子标题长度需大于 3 个字",
			"max_cn:帖子标题不能超过 40 个字",
		},
		"body": []string{
			"required:帖子内容为必填项, 参数名 body",
			"min_cn:帖子内容长度需大于 10 个字",
			"max_cn:帖子内容不能超过 50000 个字",
		},
		"category_id": []string{
			"required:帖子分类为必填项, 参数名 category_id",
			"exists:帖子分类未找到",
		},
	}

	return base.ValidateData(data, rules, messages)
}
