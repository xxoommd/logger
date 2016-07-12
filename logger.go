package logger

import (
	"fmt"
	"os"
	"runtime"

	"github.com/kr/pretty"
	"github.com/xxoommd/beego/logs"
)

var (
	logger *logs.BeeLogger
)

func init() {
	logDir := "./logs"
	if err := os.Mkdir(logDir, os.ModePerm); err != nil && !os.IsExist(err) {
		fmt.Printf("create dir %s fail : %s", logDir, err.Error())
		return
	}

	logger = logs.NewLogger(10000)
	SetLoggerLevel(logs.LevelDebug)
}

// SetLoggerLevel method
func SetLoggerLevel(level int) {
	logger.DelLogger("console")
	logger.SetLogger("console", fmt.Sprintf(`{"level": %d}`, level))

	logger.DelLogger("file")
	logger.SetLogger("file", fmt.Sprintf(`{"filename":"logs/server.log", "level":%d, "daily": true}`, level))

	if level == logs.LevelDebug {
		logger.EnableFuncCallDepth(true)
		logger.SetLogFuncCallDepth(2)
	} else {
		logger.EnableFuncCallDepth(false)
	}
}

// Emergency method
func Emergency(format string, v ...interface{}) { logger.Emergency(format, v...) }

// Alert method
func Alert(format string, v ...interface{}) { logger.Alert(format, v...) }

// Critical method
func Critical(format string, v ...interface{}) { logger.Critical(format, v...) }

// Err method
func Err(format string, v ...interface{}) { logger.Error(format, v...) }

// Notice method
func Notice(format string, v ...interface{}) { logger.Notice(format, v...) }

// Debug method
func Debug(format string, v ...interface{}) { logger.Debug(format, v...) }

// Warning method
func Warning(format string, v ...interface{}) { logger.Warn(format, v...) }

// Info method
func Info(format string, v ...interface{}) { logger.Info(format, v...) }

// Panic method
func Panic(format string, v ...interface{}) {
	Err(format, v...)
	errMsg := fmt.Sprintf(format, v...)
	for i := 1; ; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		Critical("%s:%v", file, line)
	}
	panic(errMsg)
}

// PrettyPrint method
func PrettyPrint(v ...interface{}) {
	Debug("%# v", pretty.Formatter(v))
}
