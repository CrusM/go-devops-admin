package sms

import (
	"go-devops-admin/pkg/config"
	"sync"
)

//  短信消息 结构体
type Message struct {
	Template string
	Data     map[string]string

	Content string
}

// 发送短信的操作类
type SMS struct {
	Driver Driver
}

// 单例模式
var once sync.Once

// 内部使用的 SMS 对象
var internalSMS *SMS

// 单例模式获取

func NewSMS() *SMS {
	once.Do(func() {
		internalSMS = &SMS{
			Driver: &AliYunGo{},
		}
	})
	return internalSMS
}

func (sms *SMS) Send(phone string, message Message) bool {
	return sms.Driver.Send(phone, message, config.GetStringMapString("sms.aliyun"))
}
