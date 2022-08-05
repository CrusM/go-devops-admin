package sms

import (
	"encoding/json"
	"go-devops-admin/pkg/logger"

	openApi "github.com/alibabacloud-go/darabonba-openapi/client"
	aliSms "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AliYunGo struct{}

func CreateClient(accessKeyId *string, accessKeySecret *string, endPoint *string) (_result *aliSms.Client, _err error) {
	config := &openApi.Config{
		// 您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = endPoint
	_result, _err = aliSms.NewClient(config)
	return _result, _err
}

func (s *AliYunGo) Send(phone string, message Message, config map[string]string) bool {
	client, _err := CreateClient(tea.String(config["access_key_id"]), tea.String(config["access_key_secret"]), tea.String(config["endpoint"]))
	if _err != nil {
		logger.ErrorString("阿里云[短信]", "登录短信接口失败", _err.Error())
	}

	templateParam, err := json.Marshal(message.Data)
	if err != nil {
		logger.ErrorString("短信[阿里云]", "解析绑定错误", err.Error())
		return false
	}

	sendSmsRequest := &aliSms.SendSmsRequest{
		TemplateCode:  tea.String(message.Template),
		PhoneNumbers:  tea.String(phone),
		SignName:      tea.String(config["sign_name"]),
		TemplateParam: tea.String(string(templateParam)),
	}

	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		_, _err = client.SendSmsWithOptions(sendSmsRequest, runtime)
		if _err != nil {
			return _err
		}
		logger.DebugString("短信[阿里云]", "发送成功", "")
		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			logger.ErrorString("短信[阿里云]", "服务商返回错误", string(tryErr.Error()))
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
			logger.ErrorString("短信[阿里云]", "发送失败", tryErr.Error())
		}
		// 如有需要，请打印 error
		util.AssertAsString(error.Message)
		return false
	}
	return true
}
