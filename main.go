package main

import (
	"github.com/aimerny/kook_bot/config"
	"github.com/aimerny/kook_bot/core/processor"
	"github.com/aimerny/kook_bot/core/server/http/controllers"
	"github.com/aimerny/kook_bot/core/server/websocket"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

var GlobalWaitGroup sync.WaitGroup

func main() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(log.InfoLevel)

	Config, err := config.LoadConf()
	if err != nil {
		log.Errorf("Load conf failed: %s", err.Error())
		time.Sleep(1 * time.Second)
		return
	}
	GlobalWaitGroup = sync.WaitGroup{}

	session := processor.InitSession(Config)
	GlobalWaitGroup.Add(1)
	go websocket.StartWebSocketServer(&GlobalWaitGroup)
	GlobalWaitGroup.Add(1)
	go controllers.StartWebHttpServer(&GlobalWaitGroup)

	//start ws proxy to connect and broadcast
	session.Start()
	log.Info("Connected to session!")
	//wait all goroutine done
	GlobalWaitGroup.Wait()

}
