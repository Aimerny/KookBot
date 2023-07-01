package server

import (
	"fmt"
	"github.com/aimerny/kook_bot/config"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
	"sync"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     checkOrigin,
}

var Clients = make(map[*websocket.Conn]bool)
var Broadcast = make(chan []byte)

func StartWebSocketServer(wg *sync.WaitGroup) {
	defer wg.Done()

	http.HandleFunc("/ws", EchoMessage)
	logrus.Info("Ready to start ws server")
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", config.Config.Host, config.Config.Port), nil)
	if err != nil {
		logrus.Error("Failed to start a websocket server")
	}
}

func EchoMessage(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.Warning("upgrade failed!")
	}
	//add client to cache
	Clients[conn] = true
	defer conn.Close()
	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			logrus.Warning("ws read message error, remove the conn from clients:", err)
			delete(Clients, conn)
			break
		}

		//todo other op
		logrus.Infof("%s send to: %s\n", conn.RemoteAddr(), string(msg))
		if err = conn.WriteMessage(msgType, msg); err != nil {
			return
		}
	}

}

func checkOrigin(r *http.Request) bool {
	return true
}
