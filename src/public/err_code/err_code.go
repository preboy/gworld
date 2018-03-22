package err_code

const (
	ERR_OK           uint32 = iota // 无错误
	ERR_FAILED                     // 操作失败
	ERR_LOGIN_FAILED               // 用户名、密码不正确
	ERR_LOGIN_DUP                  // 已在线，重复登录
	ERR_NOT_LOGIN                  // 未登录
	ERR_END
)
