package log

func log_prefix(ty int) string {
	switch ty {
	case 1:
		return "[" + get_time_string() + "] "
	case 2:
		return "\033[1;36m [" + get_time_string() + "] "
	case 3:
		return "\033[1;33m [" + get_time_string() + "] "
	case 4:
		return "\033[1;31m [" + get_time_string() + "] "
	case 5:
		return "\033[1;35m [" + get_time_string() + "] "
	}
	return ""
}

func log_suffix() string {
	return "\033[m\n"
}
