package migrate

import (
	"go-devops-admin/pkg/database"

	"gorm.io/gorm"
)

type Migrator struct {
	Folder   string
	DB       *gorm.DB
	Migrator gorm.Migrator
}

type Migration struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement"`
	Migration string `gorm:"type:varchar(255);not null;unique;"`
	Batch     int
}

func NewMigration() *Migrator {
	// 初始化必要属性
	migrator := &Migrator{
		Folder:   "database/migrations/",
		DB:       database.DB,
		Migrator: database.DB.Migrator(),
	}

	// 不存在则创建
	migrator.createMigrationsTable()

	return migrator
}

func (m *Migrator) createMigrationsTable() {
	migration := Migration{}

	// 不存在则创建
	if !m.Migrator.HasTable(&migration) {
		m.Migrator.CreateTable(&migration)
	}
}
