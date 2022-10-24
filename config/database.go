package config

import "go-devops-admin/pkg/config"

func init() {
	config.Add("database", func() map[string]interface{} {
		return map[string]interface{}{
			// 默认数据库
			"connection": config.Env("db.type", "mysql"),

			// 数据库配置
			"mysql": map[string]interface{}{
				"host":     config.Env("db.host", "127.0.0.1"),
				"port":     config.Env("db.port", "3306"),
				"database": config.Env("db.database", "admin"),
				"username": config.Env("db.username", ""),
				"password": config.Env("db.password", ""),
				"charset":  "utf8mb4",

				// 数据库连接池配置
				"max_idle_connections": config.Env("db.max_idle_connections", 10),
				"max_open_connections": config.Env("db.max_open_connections", 30),
				"max_life_seconds":     config.Env("db.max_life_seconds", 5*60),
			},

			"sqlite": map[string]interface{}{
				"database": config.Env("DB_SQL_FILE", "database/database.db"),
			},
		}

	})
}
