package validators

import "go-devops-admin/pkg/captcha"

func ValidateCaptcha(captchaID, CaptchaAnswer string, errs map[string][]string) map[string][]string {
	if ok := captcha.NewCaptcha().VerifyCaptcha(captchaID, CaptchaAnswer); !ok {
		errs["captcha_answer"] = append(errs["captcha_answer"], "图片验证码错误")
	}
	return errs
}

// 验证 2 次输入密码是否一致
func ValidatePasswordConfirm(password, password_confirm string, errs map[string][]string) map[string][]string {
	if password != password_confirm {
		errs["password_confirm"] = append(errs["password_confirm"], "两次输入的密码不匹配！")
	}
	return errs
}

// 验证码校验
func ValidateVerifyCode(key, answer string, errs map[string][]string) map[string][]string {
	if ok := captcha.NewCaptcha().VerifyCaptcha(key, key); !ok {
		errs["verify_code"] = append(errs["verify_code"], "验证码错误")
	}
	return errs
}
