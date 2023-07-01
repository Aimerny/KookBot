package handler

import (
	"errors"
	"github.com/aimerny/kook_bot/server"
	"github.com/bytedance/sonic"
	"github.com/gookit/event"
	"github.com/gorilla/websocket"
	"github.com/idodo/golang-bot/kaihela/api/base"
	kookEvent "github.com/idodo/golang-bot/kaihela/api/base/event"
	"github.com/sirupsen/logrus"
)

type ReceiveFrameHandler struct {
}

func (rf *ReceiveFrameHandler) Handle(e event.Event) error {
	logrus.WithField("event", e).Info("Received Frame")
	if _, ok := e.Data()[base.EventDataFrameKey]; !ok {
		return errors.New("data has no frame field")
	}
	frame := e.Data()[base.EventDataFrameKey].(*kookEvent.FrameMap)
	data, err := sonic.Marshal(frame.Data)
	if err != nil {
		logrus.Error(err)
		return err
	}
	go func() {
		for conn, status := range server.Clients {
			if conn == nil || !status {
				delete(server.Clients, conn)
				continue
			}

			if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
				logrus.Warningf("Send msg failed to client: %v", conn)
				return
			}
		}
	}()
	return nil
}
