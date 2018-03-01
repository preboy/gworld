package protocol

// 按照功能 分段

const (
	//start flag
	PROTO_BEGIN             uint16 = iota

	// general
	MSG_CS_PING             uint16 = 0x0001
	MSG_SC_PING             uint16 = 0x0002

	// client to/from server
	MSG_CS_LOGIN            uint16 = 0x0003
	MSG_SC_LOGIN            uint16 = 0x0004

	MSG_CS_ENTER_GAME       uint16 = 0x0005
	MSG_SC_ENTER_GAME       uint16 = 0x0006

	MSG_CS_PlayerData       uint16 = 0x0007
	MSG_SC_PlayerData       uint16 = 0x0008

	MSG_CS_GMCommand        uint16 = 0x0009
	MSG_SC_GMCommand        uint16 = 0x000A

    MSG_CS_Notice           uint16 = 0x000B
	MSG_SC_Notice           uint16 = 0x000C
    


	MSG_OTHER               uint16 = iota + 0x1000

	// end flag
	MSG_END                 uint16 = 0x2000
)
