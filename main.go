package main

import (
	"github.com/aimerny/kook_bot/config"
	"github.com/aimerny/kook_bot/processor"
	"github.com/aimerny/kook_bot/server"
	"github.com/idodo/golang-bot/kaihela/api/base"
	log "github.com/sirupsen/logrus"
	"sync"
)

const (
	BASE_API_URL = "https://www.kookapp.cn/api"
)

var GlobalWaitGroup sync.WaitGroup

func main() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(log.InfoLevel)

	Config, err := config.LoadConf()
	if err != nil {
		log.Error("Load conf failed!")
		panic(err)
	}
	GlobalWaitGroup = sync.WaitGroup{}

	session := base.NewWebSocketSession(Config.Token, BASE_API_URL, "./session.pid", "", 1)
	processor.InitSessionProcess(session)
	GlobalWaitGroup.Add(1)
	go server.StartWebSocketServer(&GlobalWaitGroup)
	//start ws proxy to connect and broadcast
	session.Start()
	log.Info("Connected to session!")
	//wait all goroutine done
	GlobalWaitGroup.Wait()

}
