package utils

type MqttResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func MResult(code int, data interface{}, msg string) string{
	return DataByJsonStr(MqttResponse{code, msg,data})
}

func Success(code int) string{
	return MResult(code, map[string]interface{}{}, "成功")
}

func SuccessWithData(code int,data interface{}) string{
	return MResult(code, data, "成功")
}

//500	执行中服务器内部错误
func MFailInternalServerError(code int,message string) string{
	return MResult(code, map[string]interface{}{}, message)
}
//501   参数错误
func MFailNotImplemented(code int, message string)string{
	return MResult(code, map[string]interface{}{}, message)
}
