package lib

import (
	"github.com/gin-gonic/gin"
	"common_gin/common"
)

const (
	TRACE = iota
	DEBUG
	INFO
	WARNING
	ERROR
	FATAL
)

//错误日志
func Warn(c *gin.Context, m map[string]interface{}) {
	traceContext := GetGinTraceContext(c)
	common.LogrusLogger.WithFields(map[string]interface{}{
		"trace_id": traceContext.TraceId,
		"span_id":traceContext.SpanId,
		"tag":      WARNING,
		"msg":      m,
	}).Info()
}

//错误日志
func Info(c *gin.Context, m map[string]interface{}) {
	traceContext := GetGinTraceContext(c)
	common.LogrusLogger.WithFields(map[string]interface{}{
		"trace_id": traceContext.TraceId,
		"span_id":traceContext.SpanId,
		"tag":      INFO,
		"msg":      m,
	}).Info()
}

// 从gin的Context中获取数据
func GetGinTraceContext(c *gin.Context) *common.TraceContext {
	traceContext, exists := c.Get("trace")
	if exists {
		if tc, ok := traceContext.(*common.TraceContext); ok {
			return tc
		}
	}
	trace := common.NewTrace()
	c.Set("trace", trace)
	return trace
}
