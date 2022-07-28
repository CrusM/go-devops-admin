package database

import (
	"database/sql"
	"fmt"

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
