package bootstrap

import (
	"errors"
	"fmt"

	"go-devops-admin/app/models/user"
	"go-devops-admin/pkg/config"
	"go-devops-admin/pkg/database"
	"go-devops-admin/pkg/logger"
	"time"

	"gorm.io/driver/mysql"
	// "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// 初始化数据库和ORM
func SetupDB() {
	var dbConfig gorm.Dialector
	switch config.Get("database.connection") {
	case "mysql":
		// 构建dns信息, 数据库连接字符串
		dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&multiStatements=true&loc=Local",
			config.Get("database.mysql.username"),
			config.Get("database.mysql.password"),
			config.Get("database.mysql.host"),
			config.Get("database.mysql.port"),
			config.Get("database.mysql.database"),
			config.Get("database.mysql.charset"),
		)
		dbConfig = mysql.New(mysql.Config{
			DSN: dsn,
		})
	case "sqlite":
		// 初始化 sqlite
		// database := config.Get("database.sqlite.database")
		// dbConfig = sqlite.Open(database)
		panic(errors.New("sqlite does not support"))
	default:
		panic(errors.New("database connection not found"))
	}

	//  数据库连接， 并设置GORM的日志模式
	database.Connect(dbConfig, logger.NewGormLogger())

	// 设置最大空闲连接数
	database.SQL_DB.SetMaxIdleConns(config.GetInt("database.mysql.max_idle_connections"))
	// 设置最大连接数
	database.SQL_DB.SetMaxOpenConns(config.GetInt("database.mysql.max_open_connections"))
	// 设置最大超时时间
	database.SQL_DB.SetConnMaxLifetime(time.Duration(config.GetInt("database.mysql.max_life_seconds")) * time.Second)

	// 自动迁移
	AutoMigrate()
}

func AutoMigrate() {
	database.DB.AutoMigrate(&user.User{})
}
