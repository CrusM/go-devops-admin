package models

import (
	"go-devops-admin/app/base"
	"go-devops-admin/pkg/database"
)

type Category struct {
	base.BaseModel

	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`

	base.CommonTimestampField
}

func (categories *Category) Create() {
	database.DB.Create(&categories)
}

func (categories *Category) Save() (rowsAffected int64) {
	result := database.DB.Save(&categories)
	return result.RowsAffected
}

func (categories *Category) Delete() (rowAffected int64) {
	result := database.DB.Delete(&categories)
	return result.RowsAffected
}
