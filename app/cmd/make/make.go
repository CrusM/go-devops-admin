package make

import (
	"embed"
	"fmt"
	"go-devops-admin/pkg/console"
	"go-devops-admin/pkg/file"
	"go-devops-admin/pkg/str"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
)

type Model struct {
	TableName          string
	StructName         string
	StructNamePlural   string
	VariableName       string
	VariableNamePlural string
	PackageName        string
}

//go:embed stubs
var stubsFS embed.FS

var CmdMake = &cobra.Command{
	Use:   "make",
	Short: "Generate file and code",
}

func init() {
	// 注册 make 子命令
	CmdMake.AddCommand(
		CmdMakeCMD,
		CmdMakeModel,
		CmdMakeApiController,
		CmdMakeApiRequest,
	)
}

// 格式化用户输入内容
func makeModelFormString(name string) (model Model) {
	model.StructName = str.Singular(strcase.ToCamel(name))
	model.StructNamePlural = str.Plural(model.StructName)
	model.TableName = str.Snake(model.StructNamePlural)
	model.VariableName = str.LowerCamel(model.TableName)
	model.PackageName = str.Snake(model.StructName)
	model.VariableNamePlural = str.Plural(model.VariableName)
	return
}

// 读取 .stub 文件内容并进行变量替换
// 最后一个参数可选, 如若传参 应传 map[string]string 类型, 作为附加的变量替换
func createFileFromStub(filePath string, stubName string, model Model, variables ...interface{}) {
	// 实现最后一个参数可选
	replaces := make(map[string]string)

	if len(variables) > 0 {
		replaces = variables[0].(map[string]string)
	}

	// 目标文件存在
	if file.Exists(filePath) {
		console.Exit(filePath + "already exists!")
	}

	modelData, err := stubsFS.ReadFile("stubs/" + stubName + ".stub")
	if err != nil {
		console.Exit(err.Error())
	}
	modelStub := string(modelData)

	// 添加默认替换的变量
	replaces["{{VariableName}}"] = model.VariableName
	replaces["{{StructName}}"] = model.StructName
	replaces["{{StructNamePlural}}"] = model.StructNamePlural
	replaces["{{TableName}}"] = model.TableName
	replaces["{{PackageName}}"] = model.PackageName
	replaces["{{VariableNamePlural}}"] = model.VariableNamePlural

	// 对模板内容进行替换
	for search, replace := range replaces {
		modelStub = strings.ReplaceAll(modelStub, search, replace)
	}

	// 存储到目标文件中
	err = file.Put([]byte(modelStub), filePath)
	if err != nil {
		console.Exit(err.Error())
	}

	// 提示成功
	console.Success(fmt.Sprintf("[%s] created", filePath))
}
