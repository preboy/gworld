package protocol

type ProtoID uint16

// 按照功能 分段

const (
	//start flag
	PROTO_BEGIN ProtoID = iota

	// general
	CS_PING
	SS_PONG

	// client to/from server
	CS_LOGIN
	SC_LOGIN
	CS_ENTER_GAME

	CS_OTHER ProtoID = iota + 0x1000

	// end flag
	PROTO_END ProtoID = 0x1000
)
