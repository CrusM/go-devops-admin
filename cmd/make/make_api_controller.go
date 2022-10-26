package make

import (
	"fmt"
	"go-devops-admin/pkg/console"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var CmdMakeApiController = &cobra.Command{
	Use:   "api",
	Short: "Create api controller, example: make api v1/user",
	Run:   runMakeApiController,
	Args:  cobra.ExactArgs(1),
}

func runMakeApiController(cmd *cobra.Command, args []string) {
	// 处理参数, 要求附带 API 版本(v1 或 v2)
	array := strings.Split(args[0], "/")
	if len(array) != 2 {
		console.Exit("api controller name format: v1/auth")
	}

	apiVersion, apiName := array[0], array[1]
	model := makeModelFormString(apiName)

	// 组建目录
	dir := fmt.Sprintf("app/%s/controllers/%s/", model.PackageName, apiVersion)
	os.MkdirAll(dir, os.ModePerm)

	apiModel := make(map[string]string)
	apiModel["{{apiVersion}}"] = apiVersion

	// 基于模板创建文件
	fmt.Printf("%v \n", model)
	createFileFromStub(dir+model.PackageName+"_controller.go", "controller/controller", model, apiModel)

}
