package servers

import (
	"errors"
	"fmt"
	"github.com/SayHappy/growth/conf"
	"github.com/SayHappy/growth/helpers"
	"github.com/SayHappy/growth/models"
	"net/http"
	"strconv"
)

func SendMsg(telephone string) error { // 发送短信
	lastedCode, _ := models.Redis.Do("GET", telephone)
	if lastedCode != nil {
		return errors.New("请勿频繁发送")
	}
	code := helpers.Random(10000, 99999) // 五位短信
	codeStr := strconv.Itoa(code)
	codeInfo := "{\"code\":" + codeStr + "}" // 拼接json
	err := helpers.SendSms(conf.Configuration.ALiDaYu.AccessKeyID, conf.Configuration.ALiDaYu.AccessKeySecret, telephone, conf.Configuration.ALiDaYu.SignName, codeInfo, conf.Configuration.ALiDaYu.TemplateParam)
	if err != nil {
		fmt.Println(err)
		return errors.New("发送失败")
	}
	models.Redis.Do("SET", telephone, codeStr, "300") // 设置过期时间
	return nil
}

func ValidUser(user *models.User) (string, error) {
	telephone := user.Telephone
	password := user.Password
	code := user.Code
	validCode, _ := models.Redis.Do("GET", telephone) // 设置过期时间
	isPhone, _ := helpers.ValidPhone(telephone)
	isPassword, _ := helpers.ValidPassword(password)
	if !isPhone && !isPassword && code == "" && code != validCode {
		return "账号密码格式有误", errors.New(string(http.StatusBadRequest))
	}
	return "", nil
}

func CreateUser(user *models.User) (string, error) {
	userInfo, err := user.QueryByTelephone()
	fmt.Println(userInfo)
	if err == nil {
		return "已注册，请直接登录", errors.New("已注册，请直接登录")
	}
	createUserErr := user.Insert()
	if createUserErr != nil {
		return "创建失败", errors.New("创建失败")
	}
	return "", nil
}
