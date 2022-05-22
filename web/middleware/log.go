package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"time"
)

var logMiddleware gin.HandlerFunc = func(c *gin.Context) {
	log := logrus.StandardLogger()
	// 请求方式
	reqMethod := c.Request.Method
	// 请求路由
	reqUri := c.Request.RequestURI
	// HTTP版本
	reqProto := c.Request.Proto
	log.Infof(`-> %s "%s" %s <-`, reqMethod, reqUri, reqProto)
	// 请求Body
	var buf bytes.Buffer
	tee := io.TeeReader(c.Request.Body, &buf)
	requestBody, _ := ioutil.ReadAll(tee)
	c.Request.Body = ioutil.NopCloser(&buf)
	log.Infof(`-> %s`, requestBody)
	// 开始时间
	startTime := time.Now()
	// 处理请求
	c.Next()
	// 结束时间
	endTime := time.Now()
	// 执行时间
	latencyTime := endTime.Sub(startTime)
	// 状态码
	statusCode := c.Writer.Status()
	// 日志等级
	log.WithFields(logrus.Fields{
		"statusCode": statusCode,
		"latency":    latencyTime,
	}).Infof(`==> %s "%s" <==`, reqMethod, reqUri)
}
