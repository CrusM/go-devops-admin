package user

import (
	"go-devops-admin/app/models"
)

// 用户模型
// json:"-",指定在JSON解析器忽略字段
type User struct {
	models.BaseModel

	Name     string `json:"name,omitempty"`
	Email    string `json:"-"`
	Phone    string `json:"-"`
	Password string `json:"-"`

	models.CommonTimestampField
}
