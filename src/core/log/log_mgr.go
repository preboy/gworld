package log

import (
	"time"
)

var (
	gLogger *Logger
)

func Start(name string) {
	filename := name + time.Now().Format("__2006-01-02__15_04_05") + ".log"
	if gLogger == nil {
		gLogger = NewLogger(filename)
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

func get_time_string() string {
	return time.Now().Format("2006-01-02 15:04:05.000")
}
