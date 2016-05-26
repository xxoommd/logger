package logger

import (
	"fmt"
	"os"
	"runtime"

	"github.com/astaxie/beego/logs"
	"github.com/kr/pretty"
)

var (
	logger   *logs.BeeLogger
	opLogger *logs.BeeLogger
)

func init() {
	logDir := "./logs"
	if err := os.Mkdir(logDir, os.ModePerm); err != nil && !os.IsExist(err) {
		fmt.Printf("create dir %s fail : %s", logDir, err.Error())
		return
	}

	logger = logs.NewLogger(10000)
	opLogger = logs.NewLogger(10000)

	logger.EnableFuncCallDepth(true)
	logger.SetLogFuncCallDepth(3)

	opLogger.EnableFuncCallDepth(true)
	opLogger.SetLogFuncCallDepth(3)

	SetLogger(logs.LevelDebug)
}

// SetLogger method
func SetLogger(level int) {
	logger.SetLogger("console", fmt.Sprintf(`{"level": %d}`, level))
	logger.SetLogger("file", fmt.Sprintf(`{"filename":"logs/server.log", "level":%d}`, level))
	opLogger.SetLogger("file", fmt.Sprintf(`{"filename": "logs/op.log", "level":%d}`, level))
}

// Emergency method
func Emergency(format string, v ...interface{}) { logger.Emergency(format, v...) }

// Alert method
func Alert(format string, v ...interface{}) { logger.Alert(format, v...) }

// Critical method
func Critical(format string, v ...interface{}) { logger.Critical(format, v...) }

// Error method
func Error(format string, v ...interface{}) { logger.Error(format, v...) }

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
	Error("Panic: "+format, v...)
	// Error(format, v...)
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

// OpLog method
func OpLog(format string, v ...interface{}) { opLogger.Debug(format, v...) }

// PrettyPrint method
func PrettyPrint(v ...interface{}) {
	Debug("%# v", pretty.Formatter(v))
}

