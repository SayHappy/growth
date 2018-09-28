package models

import (
	"fmt"
	"github.com/SayHappy/growth/conf"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

var DB *gorm.DB
var Redis redis.Conn

func InitDB() error {
	db, err := gorm.Open("mysql", "root:xiangdushen001@tcp(localhost:3306)/growth")
	if err != nil {
		return err
	}
	DB = db
	DB.AutoMigrate(&User{})

	r, err := redis.Dial("tcp", conf.Configuration.Host+":"+conf.Configuration.RedisPort)
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return err
	}
	Redis = r
	return nil
}
