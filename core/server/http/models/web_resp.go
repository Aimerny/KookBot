package models

import "net/http"

type Resp[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

func Success(data any) *Resp[any] {
	return &Resp[any]{
		Code: http.StatusOK,
		Msg:  "成功",
		Data: data,
	}
}

func FailWithMsg(code int, msg string) *Resp[any] {
	return &Resp[any]{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}

func FailWithData(code int, msg string, data any) *Resp[any] {
	return &Resp[any]{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func Fail(code int) *Resp[any] {
	return FailWithMsg(code, "请求失败")
}
