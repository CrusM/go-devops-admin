package seeders

import "go-devops-admin/pkg/seed"

func Initialize() {
	// 触发加载本目录下所有文件中的 init 方法

	// 指定优先于同目录下的其他文件运行
	seed.SetRunOrder([]string{
		"SeedUsersTable",
	})
}
