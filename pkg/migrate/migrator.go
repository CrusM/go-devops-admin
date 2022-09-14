package migrate

import (
	"go-devops-admin/pkg/console"
	"go-devops-admin/pkg/database"
	"go-devops-admin/pkg/file"
	"os"

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

func (m *Migrator) Up() {
	// 读取所有迁移文件, 确保按照时间排序
	migrateFiles := m.readAllMigrationFiles()

	// 获取当前批次的值
	batch := m.getBatch()

	// 获取所有迁移数据
	migrations := []Migration{}
	m.DB.Find(&migrations)

	// 标记位, 通过这个值判断数据是否已经更新到最新版本
	ran := false

	// 对迁移文件进行遍历, 如果没有执行过, 就执行 UP 回调

	for _, file := range migrateFiles {
		// 对比文件名称, 判断是否已经运行过
		if file.isNotMigrated(migrations) {
			m.runUpMigration(file, batch)
			ran = true
		}
	}

	if !ran {
		console.Success("database is up to date.")
	}

}

func (m *Migrator) readAllMigrationFiles() []MigrationFile {
	// 读取 database/migrations/ 目录下的所有文件
	// 默认按照文件名进行排序
	files, err := os.ReadDir(m.Folder)
	console.ExitIf(err)

	var migrateFiles []MigrationFile
	for _, f := range files {
		// 去除文件后缀 .go
		fileName := file.FileNameWithoutExtension(f.Name())

		// 通过迁移文件, 获取对象名称
		mFile := getMigrationFile(fileName)

		// 确保迁移文件可用, 再放进 migrateFiles 数组中
		if len(mFile.FileName) > 0 {
			migrateFiles = append(migrateFiles, mFile)
		}
	}
	return migrateFiles
}

func (m *Migrator) getBatch() int {
	// 默认 batch 为 1
	batch := 1
	// 取最后执行的一条迁移数据
	lastMigration := Migration{}
	m.DB.Order("id DESC").First(&lastMigration)

	// 如果有值, 则 +1
	if lastMigration.ID > 0 {
		batch = lastMigration.Batch + 1
	}
	return batch
}

func (m *Migrator) runUpMigration(file MigrationFile, batch int) {
	// 执行 up 区块的 SQL
	if file.Up != nil {
		console.Warning("migrating " + file.FileName)

		// 执行 up 方法
		file.Up(database.DB.Migrator(), database.SQL_DB)
		// 提示已迁移了哪个文件
		console.Success("migrated " + file.FileName)
	}

	// 入库
	err := database.DB.Create(&Migration{
		Migration: file.FileName,
		Batch:     batch,
	}).Error
	console.ExitIf(err)
}

// 回滚上一个操作
func (m *Migrator) RollBack() {
	// 获取最后一批次的迁移数据
	lastMigration := Migration{}
	m.DB.Order("id DESC").First(&lastMigration)

	migrations := []Migration{}
	m.DB.Where("batch = ?", lastMigration.Batch).Order("id DESC").Find(&migrations)

	// 回滚最后一批次的迁移
	if !m.RollBackMigrations(migrations) {
		console.Success("[migrations] table is empty, nothing to rollback.")
	}
}

// 回退迁移, 按照顺序执行 down 方法
func (m *Migrator) RollBackMigrations(migrations []Migration) bool {
	// 标记是否执行
	ran := false

	for _, _migration := range migrations {
		console.Warning("RollBack" + _migration.Migration)

		// 执行迁移文件
		file := getMigrationFile(_migration.Migration)
		if file.Down != nil {
			file.Down(database.DB.Migrator(), database.SQL_DB)
		}

		ran = true
		// 回退成功，删除对应的记录
		m.DB.Delete(&_migration)

		console.Success("Finish " + file.FileName)
	}
	return ran
}

// 回滚所有迁移
func (m *Migrator) Reset() {
	migrations := []Migration{}

	m.DB.Order("id DESC").Find(&migrations)

	// 回滚所有迁移
	if !m.RollBackMigrations(migrations) {
		console.Success("[migrations] table is empty, nothing to rollback.")
	}
}

// 回滚所有迁移, 并运行所有迁移
func (m *Migrator) Refresh() {
	m.Reset()

	m.Up()
}
