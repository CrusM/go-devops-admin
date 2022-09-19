package seeders

import (
	"fmt"
	"go-devops-admin/database/factories"
	"go-devops-admin/pkg/console"
	"go-devops-admin/pkg/logger"
	"go-devops-admin/pkg/seed"

	"gorm.io/gorm"
)

func init() {
	// 添加 seeder
	seed.Add("SeederTopicsTable", func(db *gorm.DB) {
		// 创建 10 个 user 对象
		topics := factories.MakeTopics(10)

		// 批量创建用户
		result := db.Table("topics").Create(&topics)

		// 记录错误
		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}
