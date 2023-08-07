package controllers

import (
	"github.com/aimerny/kook_bot/constant"
	"github.com/aimerny/kook_bot/core/server/http/models"
	"github.com/bytedance/sonic"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func send(context *gin.Context) {
	var req models.SendMsgReq
	var resp models.SendMsgResp

	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(req.Content) <= 0 {
		log.Warnf("Empty msg content, skip:%v", req)
		context.JSON(http.StatusBadRequest, models.FailWithMsg(http.StatusBadRequest, "Message content is null"))
		return
	}

	client := apiWithPath(constant.SendChannelChat)
	reqData, err := sonic.Marshal(req)
	if err != nil {
		log.Warnf("Error marshal :%v, %e", req, err)
		return
	}
	respBytes, err := client.SetBody(reqData).Post()
	if err != nil {
		log.Errorf("Send message failed with req:%v", req)
		context.JSON(http.StatusInternalServerError, models.FailWithMsg(http.StatusInternalServerError, err.Error()))
		return
	}
	err = sonic.Unmarshal(respBytes, resp)
	context.JSON(http.StatusOK, models.Success(resp))
}
