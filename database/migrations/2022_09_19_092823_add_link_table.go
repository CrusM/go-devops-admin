package migrations

import (
	"database/sql"
	"go-devops-admin/app"
	"go-devops-admin/pkg/migrate"

	"gorm.io/gorm"
)

func init() {
	type Link struct {
		app.BaseModel

		Name string `gorm:"type:varchar(255);not null"`
		URL  string `grom:"type:varchar(255);default;null"`

		app.CommonTimestampField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Link{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Link{})
	}

	migrate.Add("2022_09_19_092823_add_link_table", up, down)
}
