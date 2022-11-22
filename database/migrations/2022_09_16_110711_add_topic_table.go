package migrations

import (
	"database/sql"
	"go-devops-admin/app/base"
	"go-devops-admin/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type User struct {
		base.BaseModel
	}
	type Category struct {
		base.BaseModel
	}

	type Topic struct {
		base.BaseModel

		Title      string `gorm:"type:varchar(255);not null;index"`
		Body       string `gorm:"type:longtext;not null"`
		UserID     string `gorm:"type:bigint;not null;index"`
		CategoryID string `gorm:"type:bigint;not null;index"`

		// 关联用户模块
		User User
		// 管理分类模块
		Category Category

		base.CommonTimestampField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Topic{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Topic{})
	}

	migrate.Add("2022_09_16_110711_add_topic_table", up, down)
}
