package migrations

import (
	"database/sql"
	"go-devops-admin/app/base"
	"go-devops-admin/pkg/migrate"

	"gorm.io/gorm"
)

func init() {
	type BackupHost struct {
		base.BaseModel

		// FIXME()
		Hostname string `grom:"type:varchar(100);not null;"`
		Ip       string `gorm:"type:varchar(32);"`
		Env      string `gorm:"type:varchar(16);default:test"`
		Project  string `gorm:"type:varchar(255);not null;"`
		IsActive bool   `gorm:"type:bool;default;"`

		base.CommonTimestampField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&BackupHost{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&BackupHost{})
	}

	migrate.Add("2022_10_26_150444_backup_host", up, down)
}
