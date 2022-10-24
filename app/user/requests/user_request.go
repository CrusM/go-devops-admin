package requests

import (
	"go-devops-admin/app"
	"go-devops-admin/app/validators"
	"go-devops-admin/pkg/auth"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type UserRequest struct {
	// request 字段
	// Name string `json:"name,omitempty" valid:"name"`
}

func UserSave(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		// "name": []string{"required"},
	}

	messages := govalidator.MapData{
		// "name": []string{"required:name 为必填项"},
	}

	return app.ValidateData(data, rules, messages)
}

type UserUpdateProfileRequest struct {
	// request 字段
	Name         string `json:"name" valid:"name"`
	City         string `json:"city" valid:"city"`
	Introduction string `json:"introduction " valid:"introduction"`
}

func UserUpdateProfile(data interface{}, c *gin.Context) map[string][]string {
	uid := auth.CurrentUID(c)
	rules := govalidator.MapData{
		"name":         []string{"required", "alpha_num", "between:3,20", "not_exists:user,name," + uid},
		"city":         []string{"min_cn:2", "max_cn:20"},
		"introduction": []string{"min_cn:4", "max_cn:240"},
	}

	messages := govalidator.MapData{
		"name": []string{
			"required:用户名必须填写",
			"alpha_num:用户名格式错误,只允许数字和英文",
			"between:用户名长度需在 3~20 之间",
			"not_exists:用户名已经被占用",
		},
		"city": []string{
			"min_cn:城市名称需大于2个字",
			"max_cn:城市名称需大于20个字",
		},
		"introduction": []string{
			"min_cn:描述长度需大于4个字",
			"max_cn:描述长度需小于240个字",
		},
	}

	return app.ValidateData(data, rules, messages)
}

type UserUpdateEmailRequest struct {
	Email      string `json:"email,omitempty" valid:"email" `
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
}

func UserUpdateEmail(data interface{}, c *gin.Context) map[string][]string {
	currentUser := auth.CurrentUser(c)

	rules := govalidator.MapData{
		"email": []string{
			"required",
			"email",
			"not_exists:users,email," + currentUser.GetStringID(),
			"not_in:" + currentUser.Email,
		},
		"verify_code": []string{"required", "digits:6"},
	}

	messages := govalidator.MapData{
		"email": []string{
			"required:Email 为必填项",
			"email: Email 格式不正确, 请提供正确的邮箱地址",
			"not_exists:Email 已经被占用",
			"not_in:新的 Email 和老的 Email 一致",
		},
		"verify_code": []string{
			"required:验证码答案为必填项",
			"digits:验证码长度为 6 位数字",
		},
	}

	errs := app.ValidateData(data, rules, messages)
	_data := data.(*UserUpdateEmailRequest)
	errs = validators.ValidateVerifyCode(_data.Email, _data.VerifyCode, errs)
	return errs
}

// 修改用户手机号
type UserUpdatePhoneRequest struct {
	Phone      string `json:"phone,omitempty" valid:"phone" `
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
}

func UserUpdatePhone(data interface{}, c *gin.Context) map[string][]string {
	currentUser := auth.CurrentUser(c)

	rules := govalidator.MapData{
		"phone": []string{
			"required",
			"digits:11",
			"not_exists:users,phone," + currentUser.GetStringID(),
			"not_in:" + currentUser.Phone,
		},
		"verify_code": []string{"required", "digits:6"},
	}

	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号码 为必填项",
			"digits: 手机号格式不正确, 必须为 11 位数字",
			"not_exists:手机号已被占用",
			"not_in:新的 手机号 和老的 手机号 一致",
		},
		"verify_code": []string{
			"required:验证码答案为必填项",
			"digits:验证码长度为 6 位数字",
		},
	}

	errs := app.ValidateData(data, rules, messages)
	_data := data.(*UserUpdatePhoneRequest)
	errs = validators.ValidateVerifyCode(_data.Phone, _data.VerifyCode, errs)
	return errs
}

// 修改密码
type UserUpdatePasswordRequest struct {
	Password           string `json:"password,omitempty" valid:"password"`
	NewPassword        string `json:"new_password,omitempty" valid:"new_password"`
	NewPasswordConfirm string `json:"new_password_confirm,omitempty" valid:"new_password_confirm"`
}

func UserUpdatePassword(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"password":             []string{"required", "min:6"},
		"new_password":         []string{"required", "min:6"},
		"new_password_confirm": []string{"required", "min:6"},
	}

	messages := govalidator.MapData{
		"password": []string{
			"required:密码为必填项,参数名 password",
			"min:密码长度必须大于6",
		},
		"new_password": []string{
			"required:新密码为必填项,参数名 new_password",
			"min:密码长度必须大于6",
		},
		"bew_password_confirm": []string{
			"required:确认密码框为必填项,参数名 bew_password_confirm",
			"min:密码长度必须大于6",
		},
	}

	return app.ValidateData(data, rules, messages)
}

// 上传用户头像
type UserUpdateAvatarRequest struct {
	// Govalidator 验证文件必须使用的类型 *multipart.FileHeader
	Avatar *multipart.FileHeader `valid:"avatar" form:"avatar"`
}

func UserUpdateAvatar(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		// size 单位 bytes,  5242880 bytes 为 5mb
		"file:avatar": []string{"required", "size:5242880", "ext:png,jpg,jpeg"},
	}

	messages := govalidator.MapData{
		"avatar": []string{
			"required:必须上传图片",
			"ext:头像只能上传 png,jpg,jpeg 格式的图片",
			"size:头像图片大小不能超过 5mb",
		},
	}
	return app.ValidateFile(c, data, rules, messages)
}
