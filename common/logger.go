package common

import (
	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"path"
	"time"
)

var LogrusLogger *logrus.Logger

func InitLogger() {
	logFilePath := GetStringConf("base.log.log_file_path")
	logFileName := GetStringConf("base.log.log_file_name")
	// 日志文件
	fileName := path.Join(logFilePath, logFileName)
	// 实例化
	logger := logrus.New()
	// 设置日志级别
	logger.SetLevel(logrus.DebugLevel)
	// 设置 rotateLogs
	logWriter, _ := rotateLogs.New(
		// 分割后的文件名称
		fileName + ".%Y%m%d.log",
		// 生成软链，指向最新日志文件
		rotateLogs.WithLinkName(fileName),
		// 设置最大保存时间(7天)
		rotateLogs.WithMaxAge(7*24*time.Hour),
		// 设置日志切割时间间隔(1天)
		rotateLogs.WithRotationTime(24*time.Hour),
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat:"2006-01-02 15:04:05",
	})
	// 新增钩子
	logger.AddHook(lfHook)

	LogrusLogger=logger
}