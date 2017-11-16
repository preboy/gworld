package log

func log_prefix(ty int) string {
	return "[" + get_time_string() + "] "
}

func log_suffix() string {
	return "\r\n"
}
