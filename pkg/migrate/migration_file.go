package migrate

import (
	"database/sql"

	"gorm.io/gorm"
)

// migrationFunc 定义 up 和 down 回调方法
type migrationFunc func(gorm.Migrator, *sql.DB)

// 代表单个迁移文件
type MigrationFile struct {
	Up       migrationFunc
	Down     migrationFunc
	FileName string
}

var migrationFiles []MigrationFile

// 新增一个迁移文件, 所有的迁移文件都需要调用此方法注册

func Add(name string, up migrationFunc, down migrationFunc) {
	migrationFiles = append(migrationFiles, MigrationFile{
		FileName: name,
		Up:       up,
		Down:     down,
	})
}

func getMigrationFile(filename string) MigrationFile {
	for _, file := range migrationFiles {
		if filename == file.FileName {
			return file
		}
	}
	return MigrationFile{}
}

func (m MigrationFile) isNotMigrated(migrations []Migration) bool {
	for _, migration := range migrations {
		if migration.Migration == m.FileName {
			return false
		}
	}
	return true
}
