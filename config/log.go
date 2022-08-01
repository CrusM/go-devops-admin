package config

import "go-devops-admin/pkg/config"

func init() {
	config.Add("log", func() map[string]interface{} {
		return map[string]interface{}{
			"level": config.Get("LOG_LEVEL", "debug"),
			// 日志滚动类型，可选(推荐daily):
			// single： 单独文件
			// daily: 按照日期每日一个文件
			"type": config.Get("LOG_TYPE", "daily"),
			/* ----------- 滚动日志配置 ----------------- */
			// 日志文件
			"filename": config.Env("LOG_NAME", "storage/logs/logs.log"),
			// 每个日志文件保存最大尺寸, 单位M
			"max_size": config.Env("LOG_MAX_SIZE", 100),
			// 最多保存日志文件数, 0为不限制, 根据max_age删除
			"max_backup": config.Env("LOG_MAX_BACKUP", 5),
			// 日志保存的最长时间,单位天, 0为不限制
			"max_age": config.Env("LOG_MAX_AGE", 7),
			// 是否压缩(压缩日志不方便查看，但是节省空间，默认不压缩)
			"compress": config.Env("LOG_COMPRESS"),
		}
	})
}
