package response

import (
	"go-devops-admin/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 响应处理工具

// 响应 200 和 JSON 数据
func JSON(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

// 响应 200 和 预设操作成功的 JSON 消息
// 主要应用于某个 没有具体返回数据 的变更操作. 如修改密码，修改手机号
func SUCCESS(c *gin.Context) {
	JSON(c, gin.H{
		"success": true,
		"message": "操作成功",
	})
}

// 响应 200 和 带 data 消息的 JSON 数据
// 主要应用于 有具体返回信息 的操作
func Data(c *gin.Context, data interface{}) {
	JSON(c, gin.H{
		"success": true,
		"data":    data,
	})
}

// 响应 201 和 带 data 消息的 JSON 数据
func Created(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    data,
	})
}

func CreatedJSON(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, data)
}

// 返回错误状态码, 未传参数 msg 时, 返回默认消息
func Abort404(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		"message": defaultMessage("数据不存在,请确定请求正确", msg...),
	})
}

func Abort403(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
		"message": defaultMessage("权限不足, 请确认您对应的权限", msg...),
	})
}

func Abort500(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"message": defaultMessage("服务器内部错误, 请稍后再试...", msg...),
	})
}

func BadRequest(c *gin.Context, err error, msg ...string) {
	logger.LogIf(err)
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"message": defaultMessage("请求解析错误，请确认请求格式是否正确.上传文件请使用 multipart 标头，参数请使用 JSON 格式.", msg...),
		"error":   err.Error(),
	})
}

// ERROR 响应 404 或者 422
func ERROR(c *gin.Context, err error, msg ...string) {
	logger.LogIf(err)

	// 判断错误类型, 数据库未找到内容 返回 404
	if err == gorm.ErrRecordNotFound {
		Abort404(c)
		return
	}

	c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
		"message": defaultMessage("请求处理失败, 请查看 error 信息", msg...),
		"error":   err.Error(),
	})
}

// 表单验证错误
func ValidationError(c *gin.Context, errors map[string][]string) {
	c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
		"message": "请求验证不通过, 具体查看errors",
		"errors":  errors,
	})
}

// 未授权错误
func Unauthorized(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"message": defaultMessage("认证失败,请重新登录...", msg...),
	})
}

// 返回默认消息
// Go 不支持参数默认值，只能使用多变参数来实现类似效果
func defaultMessage(defaultMsg string, msg ...string) (message string) {
	if len(msg) > 0 {
		message = msg[0]
	} else {
		message = defaultMsg
	}
	return
}
