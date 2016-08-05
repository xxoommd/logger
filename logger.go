package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/kr/pretty"
	"github.com/xxoommd/beego/logs"
)

var (
	logLevel int
	logPath  string
	logger   *logs.BeeLogger
)

// Init method initializes logger, should be called before use.
func Init(lv int, path string) {
	logger = logs.NewLogger(10000)
	logLevel = lv
	logPath = path
	checkLogPath()

	logger.DelLogger("console")
	logger.SetLogger("console", fmt.Sprintf(`{"level": %d}`, logLevel))

	logger.DelLogger("file")
	logger.SetLogger("file", fmt.Sprintf(`{"filename":"%s", "level":%d, "daily": true}`, logPath, logLevel))

	logger.EnableFuncCallDepth(false)
}

func checkLogPath() {
	dir := filepath.Dir(logPath)

	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			panic(err)
		}
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
