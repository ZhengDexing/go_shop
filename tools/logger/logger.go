package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var (
	file *os.File
)

type Logger struct{}

func LoggerByDay(fileName string) *Logger {
	var err error
	createDir()
	fileName = fmt.Sprintf("./logs/%s_%s%s", fileName, formatTime(), ".log")
	file, err = os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return new(Logger)
}

// 创建log文件夹
func createDir() {
	var err error
	_, err = os.Stat("logs")
	if err == nil {
		return
	}
	if os.IsNotExist(err) {
		err = os.Mkdir("logs", os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}

// 时间格式化
func formatTime() string {
	t := time.Now()
	return t.Format("2006-01-02")
}

// 记录所有日志
func (Logger) All(value ...interface{}) {
	logInfo := log.New(io.MultiWriter(file, os.Stdout), "All: ", log.Ldate|log.Ltime|log.Lshortfile)
	logInfo.Println(value)
}

// 重要的信息
func (Logger) Info(value ...interface{}) {
	logInfo := log.New(io.MultiWriter(file, os.Stdout), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	logInfo.Println(value)
}

// 需要注意的信息
func (Logger) Warning(value ...interface{}) {
	logInfo := log.New(io.MultiWriter(file, os.Stdout), "Warning: ", log.Ldate|log.Ltime|log.Lshortfile)
	logInfo.Println(value)
}

// 非常严重的问题
func (Logger) Error(value ...interface{}) {
	logInfo := log.New(io.MultiWriter(file, os.Stdout), "Error: ", log.Ldate|log.Ltime|log.Lshortfile)
	logInfo.Println(value)
}
