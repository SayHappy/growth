package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model `json:"extra"`
	Password   string `json:"password" gorm:"default:null" `
	AvatarUrl  string `json:"avatar_url"`
	NickName   string `json:"nick_name"`                                  //昵称
	Telephone  string `json:"telephone" gorm:"unique_index;default:null"` //手机号码
	SecretKey  string `json:"secret_key"`                                 //手机号码
	UserType   int    `json:"user_type"`
	Code       string `json:"code"` // 验证码
}

func (user *User) Insert() error {
	return DB.Create(user).Error
}

func (user *User) QueryByTelephone() (*User, error) {
	telephone := user.Telephone
	var userInfo *User
	fmt.Println(DB)
	err := DB.Find(&userInfo, "telephone = ?", telephone).Error
	return userInfo, err
}
