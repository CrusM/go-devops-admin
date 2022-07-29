package requests

// 处理请求数据和表单验证

import (
	"fmt"
	"net/http"

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
	if err := c.ShouldBind(&obj); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "请求解析错误,请确认请求格式正确. 上传文件请使用 multipart 标头, 参数请使用JSON格式。",
			"error":   err.Error(),
		})
		fmt.Println(err.Error())
		return false
	}

	// 验证表单
	errs := handler(obj, c)

	// 判断是否通过验证
	if len(errs) > 0 {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "请求验证不通过, 具体查看errors",
			"errors":  errs,
		})
		return false
	}

	return true
}

func validate(data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		Messages:      messages,
		TagIdentifier: "valid",
	}

	return govalidator.New(opts).ValidateStruct()
}
