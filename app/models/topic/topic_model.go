package topic

import (
	"go-devops-admin/app/models"
	"go-devops-admin/app/models/category"
	"go-devops-admin/app/models/user"
	"go-devops-admin/pkg/database"
)

type Topic struct {
	models.BaseModel

	Title      string `json:"title,omitempty"`
	Body       string `json:"body,omitempty"`
	UserID     string `json:"user_id,omitempty"`
	CategoryID string    `json:"category_id,omitempty"`

	// 关联用户模块
	User user.User `json:"user"`
	// 管理分类模块
	Category category.Category `json:"category"`

	models.CommonTimestampField
}

func (topics *Topic) Create() {
	database.DB.Create(&topics)
}

func (topics *Topic) Save() (rowsAffected int64) {
	result := database.DB.Save(&topics)
	return result.RowsAffected
}

func (topics *Topic) Delete() (rowAffected int64) {
	result := database.DB.Delete(&topics)
	return result.RowsAffected
}
