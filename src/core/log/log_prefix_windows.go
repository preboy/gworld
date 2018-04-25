package log

import (
	"fmt"
	"runtime"
)

func log_prefix(ty int) string {
	return "[" + get_time_string() + "] "
}

func log_suffix() string {
	_, file, line, _ := runtime.Caller(3)
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			file = file[i+1:]
			break
		}
	}
	return fmt.Sprintf(" [%v:%v]", file, line)
}
