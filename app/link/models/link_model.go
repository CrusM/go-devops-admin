package models

import (
	"go-devops-admin/app/base"
	"go-devops-admin/pkg/database"
)

type Link struct {
	base.BaseModel

	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`

	base.CommonTimestampField
}

func (links *Link) Create() {
	database.DB.Create(&links)
}

func (links *Link) Save() (rowsAffected int64) {
	result := database.DB.Save(&links)
	return result.RowsAffected
}

func (links *Link) Delete() (rowAffected int64) {
	result := database.DB.Delete(&links)
	return result.RowsAffected
}
