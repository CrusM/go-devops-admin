package user

import (
	"go-devops-admin/app/models"
	"go-devops-admin/pkg/database"
	"go-devops-admin/pkg/hash"

	"github.com/spf13/cast"
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

func (userModel *User) Create() {
	database.DB.Create(&userModel)
}

// 密码是否正确
func (userModel *User) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, userModel.Password)
}

// 获取字符串格式 ID
func (userModel *User) GetStringID() string {
	return cast.ToString(userModel.ID)
}
