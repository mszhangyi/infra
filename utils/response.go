package utils

import (
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	LoginBinding = 601 //用户需绑定
)

func Result(code int, data interface{}, msg string) []byte{
	// 开始时间
	return DataByJsonByte(Response{
		code,
		data,
		msg,
	})
}

func Ok() []byte{
	return Result(1, map[string]interface{}{}, "操作成功")
}
func OkWithData(data interface{})[]byte {
	return Result(1, data, "操作成功")
}

//---------------------------------------------------------			Fail
//403 服务器拒绝请求
func FailForbidden(message string) []byte{
	return Result(http.StatusForbidden, map[string]interface{}{}, "操作失败")
}
//401	鉴权
func FailUnauthorized(message string) []byte{
	return Result(http.StatusUnauthorized, map[string]interface{}{}, message)
}
//404  找不到资源
func FailNotFound ( message string) []byte{
	return Result(http.StatusNotFound, map[string]interface{}{}, message)
}
//500	执行中服务器内部错误
func FailInternalServerError(message string) []byte{
	return Result(http.StatusInternalServerError, map[string]interface{}{}, message)
}
//501   参数错误
func FailNotImplemented(message string) []byte{
	return Result(http.StatusNotImplemented, map[string]interface{}{}, message)
}
//502   参数错误
func FailStatusBadGateway(message string) []byte{
	return Result(http.StatusBadGateway, map[string]interface{}{}, message)
}
//601
func FailLoginBinding(message string) []byte{
	return Result(LoginBinding, map[string]interface{}{}, message)
}
