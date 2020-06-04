package code

const (
	Success      = 200
	Error        = 500
	InvalidParam = 400
)

var msgMap = map[int]string{
	Success:      "ok",
	Error:        "fail",
	InvalidParam: "请求参数错误",
}

func GetMsg(code int) string {
	msg, ok := msgMap[code]
	if ok {
		return msg
	}
	return msgMap[Error]
}
