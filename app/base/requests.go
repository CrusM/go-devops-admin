package base

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

type PaginationRequest struct {
	Sort     string `json:"sort" valid:"sort" form:"sort"`
	Order    string `json:"order" valid:"order" form:"order"`
	PageSize string `json:"page_size" valid:"page_size" form:"page_size"`
}

func Pagination(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"sort":      []string{"in:id,created_at,updated_at"},
		"order":     []string{"in:asc,desc"},
		"page_size": []string{"numeric_between:2,100"},
	}

	messages := govalidator.MapData{
		"sort":      []string{"in:排序字段仅支持 id,created_at,updated_at"},
		"order":     []string{"in:排序规则仅支持 asc(正序),desc(倒序)"},
		"page_size": []string{"numeric_between:每页条数的值介于 2~100 之间"},
	}
	return ValidateData(data, rules, messages)
}
