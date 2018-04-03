package err_code

const (
	// 通用
	ERR_OK                  uint32 = iota // 无错误
	ERR_FAILED                            // 操作失败
	ERR_LOGIN_FAILED                      // 用户名、密码不正确
	ERR_LOGIN_DUP                         // 已在线，重复登录
	ERR_NOT_LOGIN                         // 未登录
	ERR_UNKNOWN_ITEM                      // 未知道具ID
	ERR_ITEM_NOT_ENOUGH                   // 道具数量不足
	ERR_ITEM_UNUSABLE                     // 不可使用的道具
	ERR_ITEM_INVALID_HERO                 // 未知的英雄
	ERR_ITEM_INVALID_CONFIG               // 未找到配置

	// 功能错误码
	ERR_END
)
