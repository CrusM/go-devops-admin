package database

import (
	"database/sql"
	"errors"
	"fmt"
	"go-devops-admin/pkg/config"

	"gorm.io/gorm"
	GormLogger "gorm.io/gorm/logger"
)

var DB *gorm.DB
var SQL_DB *sql.DB

// 连接数据库

func Connect(config gorm.Dialector, _logger GormLogger.Interface) {
	// 使用 gorm.Open 连接数据库

	var err error
	DB, err = gorm.Open(config, &gorm.Config{
		Logger: _logger,
	})
	// 处理错误
	if err != nil {
		fmt.Println(err.Error())
	}

	// 获取底层的sql db
	SQL_DB, err = DB.DB()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func CurrentDatabase() (dbName string) {
	return DB.Migrator().CurrentDatabase()
}

// 删除所有表
func DeleteAllTable() error {
	var err error

	switch config.Get("database.connection") {
	case "mysql":
		err = deleteMysqlTables()
	case "sqlite":
		err = deleteAllSqliteTables()
	default:
		panic(errors.New("database connection not supported! supported mysql or sqlite"))
	}

	return err
}

func deleteAllSqliteTables() error {
	tables := []string{}

	// 读取数据库中所有的表
	err := DB.Select(&tables, "SELECT name FROM sqlite_master WHERE table='table'").Error
	if err != nil {
		return err
	}

	// 删除所有表
	for _, table := range tables {
		err := DB.Migrator().DropTable(table)
		if err != nil {
			return err
		}
	}

	return nil
}

func deleteMysqlTables() error {
	dbName := CurrentDatabase()
	tables := []string{}

	err := DB.Table("information_schema.tables").
		Where("table_schema = ?", dbName).
		Pluck("table_name", &tables).Error

	if err != nil {
		return err
	}

	// 暂时关闭外键检测
	DB.Exec("SET foreign_key_checks = 0;")

	// 删除所有表
	for _, table := range tables {
		err := DB.Migrator().DropTable(table)
		if err != nil {
			return err
		}
	}

	// 开启外键检测
	DB.Exec("SET foreign_key_checks = 1;")
	return nil
}
