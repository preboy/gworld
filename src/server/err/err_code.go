package err

const (
	ERR_OK           uint32 = iota // 无错误
	ERR_LOGIN_FAILED               // 用户名、密码不正确
	ERR_END
)
