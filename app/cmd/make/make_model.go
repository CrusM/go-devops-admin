package make

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var CmdMakeModel = &cobra.Command{
	Use:   "model",
	Short: "Crate model file, example: make model user",
	Run:   runMakeModel,
	Args:  cobra.ExactArgs(1), // 只允许传一个参数
}

func runMakeModel(cmd *cobra.Command, args []string) {
	// 格式化模型名称
	model := makeModelFormString(args[0])

	// 拼接目标文件路径
	dir := fmt.Sprintf("app/models/%s/", model.PackageName)
	os.MkdirAll(dir, os.ModePerm)

	// 从模板中创建文件
	createFileFromStub(dir+model.PackageName+"_model.go", "model/model", model)
	createFileFromStub(dir+model.PackageName+"_util.go", "model/model_util", model)
	createFileFromStub(dir+model.PackageName+"_hooks.go", "model/model_hooks", model)

}
