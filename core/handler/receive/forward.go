package receive

import (
	"errors"
	websocket2 "github.com/aimerny/kook_bot/core/server/websocket"
	"github.com/bytedance/sonic"
	"github.com/gookit/event"
	"github.com/gorilla/websocket"
	"github.com/idodo/golang-bot/kaihela/api/base"
	kookEvent "github.com/idodo/golang-bot/kaihela/api/base/event"
	"github.com/sirupsen/logrus"
)

type FrameHandler struct {
}

func (rf *FrameHandler) Handle(e event.Event) error {
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
		for conn, status := range websocket2.Clients {
			if conn == nil || !status {
				delete(websocket2.Clients, conn)
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
