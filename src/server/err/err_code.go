package err

const (
	ERR_OK           uint32 = iota // 无错误
	ERR_FAILED                     // 操作失败
	ERR_LOGIN_FAILED               // 用户名、密码不正确
	ERR_NOT_LOGIN                  // 未登录
	ERR_END
)
