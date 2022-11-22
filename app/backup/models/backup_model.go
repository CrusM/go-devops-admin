package models

import (
	"go-devops-admin/app/base"
	"go-devops-admin/pkg/database"
)

type BackupHost struct {
	base.BaseModel

	Hostname string `json:"hostname,omitempty"`
	Ip       string `json:"ip,omitempty"`
	Env      string `json:"env,omitempty"`
	Project  string `json:"project,omitempty"`
	IsActive bool   `json:"is_active,omitempty"`

	base.CommonTimestampField
}

func (backups *BackupHost) Create() {
	database.DB.Create(&backups)
}

func (backups *BackupHost) Save() (rowsAffected int64) {
	result := database.DB.Save(&backups)
	return result.RowsAffected
}

func (backups *BackupHost) Delete() (rowAffected int64) {
	result := database.DB.Delete(&backups)
	return result.RowsAffected
}
