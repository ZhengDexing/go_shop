package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"time"
)

var (
	FileName string      // 日志文件名
	All      *log.Logger // 记录所有日志
	Info     *log.Logger // 重要的信息
	Warning  *log.Logger // 需要注意的信息
	Error    *log.Logger // 非常严重的问题
)

func init() {
	file, err := os.OpenFile("errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}

	All = log.New(io.MultiWriter(file, os.Stdout), "All: ", log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(io.MultiWriter(file, os.Stdout), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(io.MultiWriter(file, os.Stdout), "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

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
	Info.Println(
		fmt.Sprintf("| %3d | %13v | %15s | %s | %s |",
			statusCode, latencyTime, clientIP, reqMethod, reqUri,
		),
	)
}
