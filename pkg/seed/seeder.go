package seed

import "gorm.io/gorm"

type SeederFunc func(*gorm.DB)

// 对应每一个 	database/seeders 目录下的 Seeder 文件
type Seeder struct {
	Func SeederFunc
	Name string
}

// 存放所有 Seeder
var seeders []Seeder

// 注册到 seeders 数组中
func Add(name string, fn SeederFunc) {
	seeders = append(seeders, Seeder{
		Name: name,
		Func: fn,
	})
}

// 按照顺序执行的 Seeder 数组
// 支持一些必须按照顺序执行的 Seeder
// 例如 Topic 的创建必须依赖 user
// 所以 TopicSeeder 应该在 UserSeeder 后执行
var orderedSeederNames []string

// 设置 【按照顺序执行的 Seeder 数组】
func SetRunOrder(names []string) {
	orderedSeederNames = names
}
