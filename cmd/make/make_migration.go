package make

// 根据模板生成迁移文件

import (
	"fmt"
	"go-devops-admin/pkg/app"
	"go-devops-admin/pkg/console"

	"github.com/spf13/cobra"
)

var CmdMakeMigration = &cobra.Command{
	Use:   "migration",
	Short: "Create a migration file, example: make migration add_user_table",
	Run:   runMakeMigration,
	Args:  cobra.ExactArgs(1),
}

func runMakeMigration(cmd *cobra.Command, args []string) {
	// 格式化日期
	timeStr := app.TimeNowInTimezone().Format("2006_01_02_150405")

	model := makeModelFormString(args[0])
	fileName := timeStr + "_" + model.PackageName

	filePath := fmt.Sprintf("database/migrations/%s.go", fileName)
	createFileFromStub(filePath, "migration/migration", model, map[string]string{"{{FileName}}": fileName})
	console.Success("Migration file created, after modify it, use `migrate up` to migrate database.")
}
