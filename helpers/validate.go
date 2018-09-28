package helpers

import (
	"regexp"
)

const (
	PhoneRegx    = "^(13[0-9]|14[579]|15[0-3,5-9]|16[6]|17[0135678]|18[0-9]|19[89])\\d{8}$" // 手机号
	PasswordRegx = "/^[a-zA-Z0-9]{8,}$/"                                                    // 8位密码
	EmailRegx    = "^[A-Za-z0-9\u4e00-\u9fa5]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$"         // 邮箱
)

func ValidPhone(phone string) (bool, error) {
	return regexp.Match(PhoneRegx, []byte(phone))
}

func ValidPassword(passowrd string) (bool, error) {
	return regexp.Match(PasswordRegx, []byte(passowrd))
}

func ValidEmail(email string) (bool, error) {
	return regexp.Match(EmailRegx, []byte(email))
}
