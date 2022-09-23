package requests

// 处理请求数据和表单验证

import (
	"go-devops-admin/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

// 验证函数类型
type ValidateFunc func(interface{}, *gin.Context) map[string][]string

//  控制器里调用示例
// if ok := requests.Validate(c, &request.UserSaveRequest{}, requests.UserSave); !ok {
// 		return
// }
func Validate(c *gin.Context, obj interface{}, handler ValidateFunc) bool {
	// 解析请求，支持JSON数据,表单请求和URL query
	if err := c.ShouldBind(obj); err != nil {
		response.BadRequest(c, err)
		return false
	}

	// 验证表单
	errs := handler(obj, c)

	// 判断是否通过验证
	if len(errs) > 0 {
		response.ValidationError(c, errs)
		return false
	}

	return true
}

func ValidateData(data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		Messages:      messages,
		TagIdentifier: "valid",
	}

	return govalidator.New(opts).ValidateStruct()
}

func ValidateFile(c *gin.Context, data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {
	opts := govalidator.Options{
		Request:       c.Request,
		Rules:         rules,
		Messages:      messages,
		TagIdentifier: "valid",
	}
	// Validate 方法来验证文件
	return govalidator.New(opts).Validate()
}
