package make

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var CmdMakeApiRequest = &cobra.Command{
	Use:   "request",
	Short: "Create api request, example: make request user",
	Run:   runMakeApiRequest,
	Args:  cobra.ExactArgs(1),
}

func runMakeApiRequest(cmd *cobra.Command, args []string) {

	model := makeModelFormString(args[0])

	// 组建目录
	dir := fmt.Sprintf("app/%s/requests/", model.PackageName)
	os.MkdirAll(dir, os.ModePerm)

	// 基于模板创建文件
	createFileFromStub(dir+model.PackageName+"_request.go", "request/request", model)

}
