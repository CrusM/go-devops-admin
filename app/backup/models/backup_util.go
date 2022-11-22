package models

import (
	"go-devops-admin/pkg/app"
	"go-devops-admin/pkg/database"
	"go-devops-admin/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(id string) (backups BackupHost) {
	database.DB.Where("id", id).First(&backups)
	return
}

func GetBy(field, value string) (backups BackupHost) {
	database.DB.Where("? = ?", field, value).First(&backups)
	return
}

func All() (backups []BackupHost) {
	database.DB.Find(&backups)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(BackupHost{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

// 分页查询
func Paginate(c *gin.Context, pageSize int) (Backups []BackupHost, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(BackupHost{}),
		&Backups,
		app.V1URL(database.TableName(&BackupHost{})),
		pageSize,
	)
	return
}
