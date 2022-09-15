package app

import (
	"go-devops-admin/pkg/config"
	"time"
)

func IsLocal() bool {
	return config.Get("app.env") == "local"
}

func IsProduction() bool {
	return config.Get("app.env") == "production"
}

func IsTesting() bool {
	return config.Get("app.env") == "testing"
}

// 获取当前时间, 支持时区
func TimeNowInTimezone() time.Time {
	chinaTimezone, _ := time.LoadLocation(config.GetString("app.timezone", "Asia/Shanghai"))
	return time.Now().In(chinaTimezone)
}

// 传参 PATH 拼接 URL
func URL(path string) string {
	return config.Get("app.url") + path
}

// 拼接带 v1 标示的 URL

func V1URL(path string) string {
	return URL("/v1" + path)
}
