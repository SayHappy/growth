package models

import (
	"fmt"
	"github.com/SayHappy/growth/conf"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB
var Redis redis.Conn

func InitDB() error {
	db, err := gorm.Open("mysql", "root:xiangdushen001@tcp(localhost:3306)/growth")
	if err != nil {
		return err
	}
	DB = db
	DB.AutoMigrate(&User{},&Exam{},&Examinee{},&Question{},&Paper{})

	r, err := redis.Dial("tcp", conf.Configuration.Host+":"+conf.Configuration.RedisPort)
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return err
	}
	Redis = r
	return nil
}
