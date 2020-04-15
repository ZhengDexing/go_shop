package filter

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_shop/tools/logger"
	"time"
)

var log = logger.LoggerByDay("routers")

// 记录请求，耗时，状态码等信息
func RoutersLogger(c *gin.Context) {
	// 开始时间
	startTime := time.Now()
	// 处理请求
	c.Next()
	// 结束时间
	endTime := time.Now()
	// 执行时间
	latencyTime := endTime.Sub(startTime)
	// 请求方式
	reqMethod := c.Request.Method
	// 请求路由
	reqUri := c.Request.RequestURI
	// 状态码
	statusCode := c.Writer.Status()
	// 请求IP
	clientIP := c.ClientIP()
	// 日志格式
	log.Info(
		fmt.Sprintf("| %3d | %13v | %15s | %s | %s |",
			statusCode, latencyTime, clientIP, reqMethod, reqUri,
		),
	)
}
