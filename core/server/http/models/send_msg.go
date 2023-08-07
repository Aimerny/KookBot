package models

type SendMsgReq struct {
	MessageType  int    `json:"type"`
	TargetId     string `json:"target_id"`
	Content      string `json:"content"`
	Quote        string `json:"quote"`
	Nonce        string `json:"nonce"`
	TempTargetId string `json:"temp_target_id"`
}

type SendMsgResp struct {
	MsgId        string `json:"msg_id"`
	MsgTimestamp int    `json:"msg_timestamp"`
	Nonce        string `json:"nonce"`
}
