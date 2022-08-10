package auth

import (
	v1 "go-devops-admin/app/http/controllers/api/v1"
	"go-devops-admin/app/requests"
	"go-devops-admin/pkg/captcha"
	"go-devops-admin/pkg/logger"
	"go-devops-admin/pkg/response"

	"github.com/gin-gonic/gin"
)

// 验证码控制器
type VerifyCodeController struct {
	v1.BaseAPIController
}

// 显示图片验证码
func (vc *VerifyCodeController) ShowCaptcha(c *gin.Context) {
	// 生成图片验证码
	id, b64s, err := captcha.NewCaptcha().GenerateCaptcha()

	// 记录错误日志
	logger.LogIf(err)

	// 返回消息给用户
	response.JSON(c, gin.H{
		"captcha_id":    id,
		"captcha_image": b64s,
	})
}

// 发送手机验证码
func (vc *VerifyCodeController) SendUsingPhone(c *gin.Context) {
	// 验证表单
	request := requests.VerifyCodePhoneRequest{}
	if ok := requests.Validate(c, &request, requests.VerifyCodePhone); !ok {
		return
	}

	// 发送 SMS
	if ok := captcha.NewVerifyCode().SendSMS(request.Phone); !ok {
		response.Abort500(c, "发送短信失败")
	} else {
		response.SUCCESS(c)
	}
}
