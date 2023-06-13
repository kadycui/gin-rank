package utils

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"time"
)

var (
	Logger   = logrus.New() // 初始化日志对象
	LogEntry *logrus.Entry
)

func init() {
	// 写入日志文件
	logPath := "logs/web" // 日志存放路径
	// 日志分隔：1. 每天产生的日志写在不同的文件；2. 只保留一定时间的日志（例如：一星期）
	Logger.SetLevel(logrus.DebugLevel) // 设置日志级别
	logWriter, _ := rotatelogs.New(
		logPath+"-%Y%m%d.log",                     // 日志文件名格式
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 最多保留7天之内的日志
		rotatelogs.WithRotationTime(24*time.Hour), // 一天保存一个日志文件
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter, // info级别使用logWriter写日志
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	Hook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05", // 格式日志时间
	})
	Logger.AddHook(Hook)
	LogEntry = logrus.NewEntry(Logger).WithField("service", "gin-rank")
}
