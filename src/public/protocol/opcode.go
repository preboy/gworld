package protocol

// 按照功能 分段

const (
	//start flag
	PROTO_BEGIN uint16 = iota

	// general
	MSG_PING

	// client to/from server
	MSG_LOGIN
	MSG_ENTER_GAME
	MSG_PlayerData

	MSG_OTHER uint16 = iota + 0x1000

	// end flag
	MSG_END uint16 = 0x2000
)
