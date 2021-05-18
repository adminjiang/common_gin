package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"common_gin/common"
	"common_gin/lib"
	"io/ioutil"
	"time"
)

// 日志记录到文件
// 请求进入日志
func RequestInLog(c *gin.Context) {
	trace := common.NewTrace()
	c.Set("trace",trace)
	c.Set("startExecTime", time.Now())
	bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes)) // Write body back
	//记录日志
	lib.Info(c, map[string]interface{}{
		"uri":    c.Request.RequestURI,
		"method": c.Request.Method,
		"args":   c.Request.PostForm,
		"body":   string(bodyBytes),
		"from":   c.ClientIP()})
}

// 请求输出日志
func RequestOutLog(c *gin.Context) {
	// after request
	endExecTime := time.Now()
	response, _ := c.Get("response")
	st, _ := c.Get("startExecTime")
	startExecTime, _ := st.(time.Time)
	lib.Info(c, map[string]interface{}{
		"uri":       c.Request.RequestURI,
		"method":    c.Request.Method,
		"args":      c.Request.PostForm,
		"from":      c.ClientIP(),
		"response":  response,
		"proc_time": endExecTime.Sub(startExecTime).Seconds(),
	})
}

func RequestLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		RequestInLog(c)
		defer RequestOutLog(c)
		c.Next()
	}
}
