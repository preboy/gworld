package protocol

// 按照功能 分段

const (
	//start flag
	PROTO_BEGIN uint16 = iota

	// general
	MSG_CS_PING
	MSG_SC_PING

	// client to/from server
	MSG_CS_LOGIN
	MSG_SC_LOGIN

	MSG_CS_ENTER_GAME
	MSG_SC_ENTER_GAME

	MSG_CS_PlayerData
	MSG_SC_PlayerData

	MSG_OTHER uint16 = iota + 0x1000

	// end flag
	MSG_END uint16 = 0x2000
)
