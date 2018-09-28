package growth

import (
	"github.com/SayHappy/growth/conf"
	"github.com/SayHappy/growth/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r:=gin.Default()
	conf.InitConf()
	models.InitDB()
	r.Run(":8080")
}
