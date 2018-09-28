package models

import (
	"github.com/jinzhu/gorm"
	"time"
	"fmt"
)

type User struct {
	gorm.Model `json:"extra"`
	Password   string    `json:"password" gorm:"default:null" `
	AvatarUrl  string    `json:"avatar_url"`
	NickName   string    `json:"nick_name"`                                  //昵称
	Telephone  string    `json:"telephone" gorm:"unique_index;default:null"` //手机号码
	SecretKey  string    `json:"secret_key" gorm:"default:null"`             //手机号码
	UserType   int       `json:"user_type"`
	OutTime    time.Time `json:"out_time"`
	Code       string    `json:"code"` // 验证码
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
