package make

import (
	"fmt"
	"go-devops-admin/pkg/console"

	"github.com/spf13/cobra"
)

var CmdMakeCMD = &cobra.Command{
	Use:   "cmd",
	Short: "Create a command, should be snake_case, example: make cmd backup_database",
	Run:   runMakeCMD,
	Args:  cobra.ExactArgs(1), // 只允许传一个参数
}

func runMakeCMD(cmd *cobra.Command, args []string) {
	// 格式化模型名称
	model := makeModelFormString(args[0])

	// 拼接目标文件路径
	filePath := fmt.Sprintf("app/cmd/%s.go", model.PackageName)

	// 从模板中创建文件
	createFileFromStub(filePath, "cmd", model)

	// 提示信息
	console.Success("command name:" + model.PackageName)
	console.Success("command variable name: cmd.Cmd" + model.StructName)
	console.Warning("please edit main.go's app.Commands slice to register command")
}
