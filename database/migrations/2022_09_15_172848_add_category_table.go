package migrations

import (
	"database/sql"
	"go-devops-admin/app/base"
	"go-devops-admin/pkg/migrate"

	"gorm.io/gorm"
)

func init() {
	type Category struct {
		base.BaseModel

		Name        string `json:"name,omitempty"`
		Description string `json:"description,omitempty"`

		base.CommonTimestampField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Category{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Category{})
	}

	migrate.Add("2022_09_15_172848_add_category_table", up, down)
}
