package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

//Ins Instance
var Ins Instance

//Instance 日志
type Instance struct {
	logger *logrus.Logger
}

func init() {
	Ins.Init()
}

//Fields For logrus.Fields
func (log *Instance) Fields(fields map[string]interface{}) logrus.Fields {
	return fields
}

//Logger 获取logger
func (log *Instance) Logger() *logrus.Logger {
	return log.logger
}

//Init 初始化
func (log *Instance) Init() {
	logPath := "./logs"
	_, err := os.Stat(logPath)
	if err != nil {
		if !os.IsExist(err) {
			os.Mkdir(logPath, 0777)
		}
	}
	logFile := logPath + "/" + time.Now().Format("20060102150405") + ".log"
	_, err = os.Stat(logFile)
	if err != nil {
		if !os.IsExist(err) {
			_, err := os.Create(logFile)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
	w, err := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println(err.Error())
	}

	log.logger = logrus.New()

	log.logger.SetOutput(w)
	//设置日志级别
	log.logger.SetLevel(logrus.DebugLevel)

	//设置日志格式
	log.logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
}

//Gin 替换gin日志
func (log *Instance) Gin() gin.HandlerFunc {

	return func(c *gin.Context) {
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
		reqURI := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		//日志格式
		log.logger.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqURI,
		)
	}
}
