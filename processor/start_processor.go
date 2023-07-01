package processor

import (
	"github.com/aimerny/kook_bot/handler"
	"github.com/idodo/golang-bot/kaihela/api/base"
)

func InitSessionProcess(session *base.WebSocketSession) {

	session.On(base.EventReceiveFrame, &handler.ReceiveFrameHandler{})
	//session.On(util.GenEventType(constant.ChannelGroup, constant.ALL), &handler.GroupEventHandler{})
	//session.On(util.GenEventType(constant.ChannelGroup, constant.KMarkDown), &handler.GroupTextEventHandler{Token: config.Config.Token, BaseUrl: "https://www.kookapp.cn/api"})

}
