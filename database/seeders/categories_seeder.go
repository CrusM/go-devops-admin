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
	seed.Add("SeederCategoriesTable", func(db *gorm.DB) {
		// 创建 10 个 user 对象
		categories := factories.MakeCategories(10)

		// 批量创建用户
		result := db.Table("categories").Create(&categories)

		// 记录错误
		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}
