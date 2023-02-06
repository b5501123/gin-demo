package res

import "gin-demo/pkg/e"

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data interface{}
}

func Success(data interface{}) Response {
	return Response{
		e.SUCCESS,
		e.GetMsg(e.SUCCESS),
		data,
	}
}

func Error(code int, message string) Response {
	return Response{
		Code: code,
		Msg:  message,
	}
}
