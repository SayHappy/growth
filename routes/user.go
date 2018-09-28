package routes

import (
	"github.com/SayHappy/growth/controllers"
	"github.com/gin-gonic/gin"
)

func initUserRoutes(router *gin.Engine) {
	api := router.Group("/growth/user/api")
	//api.GET("/", func(context *gin.Context) {
	//	context.String(http.StatusOK, "hello world")
	//})
	api.POST("/singUp", controllers.SingUp)
	api.POST("/sendMessage", controllers.SendMessage)
}
