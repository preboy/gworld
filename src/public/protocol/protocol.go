package protocol

type ProtoID uint16

// 按照功能 分段

const (
	//start flag
	PROTO_BEGIN ProtoID = iota

	// general
	CS_PING
	CS_PONG

	// client to/from server
	CS_LOGIN ProtoID = iota + 0x1000
	SC_LOGIN

	// end flag
	PROTO_END ProtoID = 0x1000
)
