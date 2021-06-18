package logUtil

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

// 获取日志文件的名字
func getFileName(baseName string) string {
	var dateTime = time.Unix(time.Now().Unix(), 0).Format("2006-01-02")
	return fmt.Sprintf("%s-%s.log", baseName, dateTime)
}

// 新建一个日志
var log = logrus.New()

type LogConfig struct {
	LogDir     string
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

func New(name string, config LogConfig) *logrus.Logger {
	if config.LogDir == "" {
		config.LogDir = "./logs/"
	}
	if config.Filename == "" {
		config.Filename = getFileName(name)
	}
	if config.MaxSize == 0 {
		config.MaxSize = 500
	}
	if config.MaxBackups == 0 {
		config.MaxBackups = 3
	}
	if config.MaxAge == 0 {
		config.MaxAge = 28
	}

	var output = &lumberjack.Logger{
		Filename:   config.LogDir + config.Filename,
		MaxSize:    config.MaxSize, // megabytes
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge, //days
		Compress:   config.Compress,
	}
	log.SetOutput(io.MultiWriter(os.Stdout, output))
	log.SetFormatter(&logrus.JSONFormatter{})

	return log
}
