package migrations

import (
	"database/sql"
	"go-devops-admin/app/models"
	"go-devops-admin/pkg/migrate"

	"gorm.io/gorm"
)

func init() {
	type User struct {
		models.BaseModel

		Name string `gorm:"varchar(255);not null;index"`
		Email    string `gorm:"type:varchar(255);index;default:null"`
        Phone    string `gorm:"type:varchar(20);index;default:null"`
        Password string `gorm:"type:varchar(255)"`

        models.CommonTimestampField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&User{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&User{})
	}

	migrate.Add("2022_09_14_094025_add_user_table", up, down)
}