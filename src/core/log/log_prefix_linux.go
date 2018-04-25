package log

import (
	"fmt"
	"runtime"
)

func log_prefix(ty int) string {
	switch ty {
	case 1:
		return "[" + get_time_string() + "] "
	case 2:
		return "\033[1;36m[" + get_time_string() + "] "
	case 3:
		return "\033[1;33m[" + get_time_string() + "] "
	case 4:
		return "\033[1;31m[" + get_time_string() + "] "
	case 5:
		return "\033[1;35m[" + get_time_string() + "] "
	}
	return ""
}

func log_suffix() string {
	_, file, line, _ := runtime.Caller(3)
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			file = file[i+1:]
			break
		}
	}
	return fmt.Sprintf("\033[m [%v:%v]", file, line)
}
