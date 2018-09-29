package main

import (
	"github.com/SayHappy/growth/conf"
	"github.com/SayHappy/growth/models"
	"github.com/SayHappy/growth/routes"
	"github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	conf.InitConf()
	models.InitDB()
	routes.InitRoutes(r)
	seelog.Info("启动app")
	r.Run(":8080")
}
