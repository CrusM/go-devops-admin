package link

import (
    "go-devops-admin/app/requests"

    "github.com/gin-gonic/gin"
    "github.com/thedevsaddam/govalidator"
)

type LinkRequest struct {
    // request 字段
    // Name string `json:"name,omitempty" valid:"name"`
}

func LinkSave(data interface{}, c *gin.Context) map[string][]string {
    rules := govalidator.MapData{
    // "name": []string{"required"},
    }

    messages := govalidator.MapData{
        // "name": []string{"required:name 为必填项"},
    }

    return requests.ValidateData(data, rules, messages)
}