package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
)

func initStaticRoutes(router *gin.Engine) {
	wd, _ := os.Getwd()
	fmt.Println(filepath.Join(wd, "./static"))
	router.Static("/static", filepath.Join(wd, "./static"))
}
