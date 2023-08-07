package controllers

import (
	"fmt"
	"github.com/aimerny/kook_bot/config"
	"github.com/aimerny/kook_bot/constant"
	"github.com/gin-gonic/gin"
	"github.com/idodo/golang-bot/kaihela/api/helper"
	"github.com/sirupsen/logrus"
	"sync"
)

func StartWebHttpServer(wg *sync.WaitGroup) {

	defer wg.Done()
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	messageRoute := router.Group("/message")
	{
		messageRoute.POST("/send", send)
	}
	manageRoute := router.Group("/manage")
	{
		manageRoute.POST("/kick")
	}
	logrus.Infof("Http Post API started at %s:%d", config.Config.Host, config.Config.HttpApiPort)
	err := router.Run(fmt.Sprintf("%s:%d", config.Config.Host, config.Config.HttpApiPort))
	if err != nil {
		logrus.Errorf("Start Http post api server failed: %v", err)
		panic(err)
	}

}

func apiWithPath(uri string) *helper.ApiHelper {
	return helper.NewApiHelper(uri, config.Config.Token, constant.BaseApiUrl, "", "")
}
