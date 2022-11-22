package requests

import (
	"go-devops-admin/app/base"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type BackupRequest struct {
	// request 字段
	// Name string `json:"name,omitempty" valid:"name"`
}

func BackupSave(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		// "name": []string{"required"},
	}

	messages := govalidator.MapData{
		// "name": []string{"required:name 为必填项"},
	}

	return base.ValidateData(data, rules, messages)
}
