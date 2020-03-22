package util

const (
	Success     = 0
	Forbidden   = 403
	ServerError = 500
)

var resultText = map[int]string{
	Success:     "成功",
	Forbidden:   "访问权限错误",
	ServerError: "服务器未知错误",
}

func RespOut(code int, data interface{}) map[string]interface{} {
	r := make(map[string]interface{})
	r["code"] = code
	r["data"] = data
	r["message"] = resultText[code]
	return r
}
