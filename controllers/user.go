package controllers

import (
	"github.com/SayHappy/growth/helpers"
	"github.com/SayHappy/growth/models"
	"github.com/SayHappy/growth/servers"
	"github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendMessage(c *gin.Context) {
	user := models.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseInfo{http.StatusBadRequest, "参数错误", nil})
		return
	}
	telephone := user.Telephone
	isTelephone, _ := helpers.ValidPhone(telephone)
	if isTelephone != true {
		c.JSON(http.StatusBadRequest, models.ResponseInfo{http.StatusBadRequest, "手机号格式错误", nil})
		return
	}
	smsErr := servers.SendMsg(telephone)
	if smsErr != nil {
		c.JSON(http.StatusBadRequest, models.ResponseInfo{http.StatusBadRequest, "请求太频繁", nil})
		return
	}
	c.JSON(http.StatusOK, models.ResponseInfo{0, "发送成功", nil})
	return
}

func SingUp(c *gin.Context) {
	user := models.User{}
	err := c.BindJSON(&user)
	if err != nil {
		seelog.Error(err)
		c.JSON(http.StatusBadRequest, models.ResponseInfo{http.StatusBadRequest, "参数错误", nil})
		return
	}
	validMsg, validErr := servers.ValidUser(&user) // 验证字符格式
	if validErr != nil {
		seelog.Error(validErr)
		c.JSON(http.StatusBadRequest, models.ResponseInfo{http.StatusBadRequest, validMsg, nil})
		return
	}
	// 创建用户
	createUserMsg, createUserErr := servers.CreateUser(&user)
	if createUserErr != nil {
		seelog.Error(createUserErr)
		c.JSON(http.StatusAccepted, models.ResponseInfo{http.StatusAccepted, createUserMsg, nil})
		return
	}
	c.JSON(http.StatusOK, models.ResponseInfo{0, "注册成功", nil})
}
