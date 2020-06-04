package v1

import (
	"github.com/gin-gonic/gin"
)

func ResponseJSON(ctx *gin.Context, errCode int, msg string, data interface{}) {
	ctx.JSON(200, Response{
		Code: errCode,
		Msg:  msg,
		Data: data,
	})
	ctx.Abort()
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type Pager struct {
	Cursor  int64 `json:"cursor"`
	Size    int64 `json:"size"`
	HasMore bool  `json:"has_more"`
}
