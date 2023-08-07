package processor

import (
	"github.com/aimerny/kook_bot/config"
	"github.com/aimerny/kook_bot/constant"
	"github.com/aimerny/kook_bot/core/handler/receive"
	"github.com/idodo/golang-bot/kaihela/api/base"
)

var session *base.WebSocketSession

func Session() *base.WebSocketSession {
	return session
}

func InitSession(conf *config.Conf) *base.WebSocketSession {
	session = base.NewWebSocketSession(conf.Token, constant.BaseApiUrl, "./session.pid", "", 1)
	initSessionProcess()
	return session
}

func initSessionProcess() {
	session.On(base.EventReceiveFrame, &receive.FrameHandler{})
	//session.On(util.GenEventType(constant.ChannelGroup, constant.ALL), &handler.GroupEventHandler{})
	//session.On(util.GenEventType(constant.ChannelGroup, constant.KMarkDown), &handler.GroupTextEventHandler{Token: config.Config.Token, BaseUrl: "https://www.kookapp.cn/api"})

}
