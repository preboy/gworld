package log

import (
	"fmt"
	"time"
)

var (
	gLogger *Logger
)

func Start(name string) {
	filename := name + time.Now().Format("__2006-01-02__15_04_05") + ".log"
	if gLogger == nil {
		gLogger = NewLogger(filename, true)
	}
}

func Stop() {
	if gLogger != nil {
		gLogger.Stop()
	}
}

func GetLogger() *Logger {
	return gLogger
}

// proxy function for gLogger

func Println(a ...interface{}) {
	fmt.Println(a...)
}

func Info(format string, a ...interface{}) {
	gLogger.Info(format, a...)
}

func Debug(format string, a ...interface{}) {
	gLogger.Debug(format, a...)
}

func Warning(format string, a ...interface{}) {
	gLogger.Warning(format, a...)
}

func Error(format string, a ...interface{}) {
	gLogger.Error(format, a...)
}

func Fatal(format string, a ...interface{}) {
	gLogger.Fatal(format, a...)
}

// aux function
func get_time_string() string {
	return time.Now().Format("2006-01-02 15:04:05.000")
}
