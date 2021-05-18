package lib

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"common_gin/common"
)

type ResponseCode int

//1000以下为通用码，1000以上为用户自定义码
const (
	SuccessCode ResponseCode = iota
)

type Response struct {
	ErrorCode ResponseCode `json:"errno"`
	ErrorMsg  string       `json:"errmsg"`
	Data      interface{}  `json:"data"`
	TraceId   interface{}  `json:"trace_id"`
	Stack     interface{}  `json:"stack"`
}

func ResponseError(c *gin.Context, code ResponseCode, err error) {
	trace, _ := c.Get("trace")
	traceContext, _ := trace.(*common.TraceContext)
	traceId := ""
	if traceContext != nil {
		traceId = traceContext.TraceId
	}
	resp := &Response{ErrorCode: code, ErrorMsg: err.Error(), Data: "", TraceId: traceId}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
	c.AbortWithError(200, err)
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	trace, _ := c.Get("trace")
	traceContext, _ := trace.(*common.TraceContext)
	traceId := ""
	if traceContext != nil {
		traceId = traceContext.TraceId
	}

	resp := &Response{ErrorCode: SuccessCode, ErrorMsg: "", Data: data, TraceId: traceId}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
}
