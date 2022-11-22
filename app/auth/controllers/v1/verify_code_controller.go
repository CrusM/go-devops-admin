package v1

import (
	authRequest "go-devops-admin/app/auth/requests"
	"go-devops-admin/app/base"
	"go-devops-admin/pkg/captcha"
	"go-devops-admin/pkg/logger"
	"go-devops-admin/pkg/response"

	"github.com/gin-gonic/gin"
)

// 验证码控制器
type VerifyCodeController struct {
	base.BaseAPIController
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
	request := authRequest.VerifyCodePhoneRequest{}
	if ok := base.Validate(c, &request, authRequest.VerifyCodePhone); !ok {
		return
	}

	// 发送 SMS
	if ok := captcha.NewVerifyCode().SendSMS(request.Phone); !ok {
		response.Abort500(c, "发送短信失败")
	} else {
		response.SUCCESS(c)
	}
}

// 发送 Email 验证码
func (vc *VerifyCodeController) SendUsingEmail(c *gin.Context) {
	// 验证表单
	request := authRequest.SignUpEmailExistRequest{}

	if ok := base.Validate(c, &request, authRequest.ValidateSignUpEmailExist); !ok {
		return
	}

	err := captcha.NewVerifyCode().SendEmail(request.Email)
	if err != nil {
		response.Abort500(c, "发送 Email 邮件失败~")
	} else {
		response.SUCCESS(c)
	}

}
