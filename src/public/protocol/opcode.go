package protocol

// 按照功能 分段

const (

	//start flag
	PROTO_BEGIN uint16 = iota

	// general
	MSG_CS_PING uint16 = 0x0001
	MSG_SC_PING uint16 = 0x0002

	// client to/from server
	MSG_CS_LOGIN uint16 = 0x0003
	MSG_SC_LOGIN uint16 = 0x0004

	MSG_CS_ENTER_GAME uint16 = 0x0005
	MSG_SC_ENTER_GAME uint16 = 0x0006

	MSG_CS_PlayerData uint16 = 0x0007
	MSG_SC_PlayerData uint16 = 0x0008

	MSG_CS_GMCommand uint16 = 0x0009
	MSG_SC_GMCommand uint16 = 0x000A

	MSG_CS_Notice uint16 = 0x000B
	MSG_SC_Notice uint16 = 0x000C

	MSG_CS_MakeBattle uint16 = 0x000D
	MSG_SC_MakeBattle uint16 = 0x000E

	MSG_SC_ItemCntChanged uint16 = 0x000F

	MSG_CS_UseItem uint16 = 0x0010
	MSG_SC_UseItem uint16 = 0x0011

	MSG_CS_MarketBuy uint16 = 0x0012
	MSG_SC_MarketBuy uint16 = 0x0013

	MSG_CS_HeroLevelup uint16 = 0x0014
	MSG_SC_HeroLevelup uint16 = 0x0015

	MSG_CS_HeroRefine uint16 = 0x0016
	MSG_SC_HeroRefine uint16 = 0x0017

	MSG_SC_HeroInfoUpdate uint16 = 0x0018

	MSG_CS_QuestList   uint16 = 0x0020
	MSG_SC_QuestList   uint16 = 0x0021
	MSG_CS_QuestOp     uint16 = 0x0022
	MSG_SC_QuestOp     uint16 = 0x0023
	MSG_SC_QuestUpdate uint16 = 0x0024

	MSG_SC_PlayerLvExpUpdate uint16 = 0x0025

	// ------------------------------------------------------------------------

	MSG_OTHER uint16 = iota + 0x1000

	// end flag
	MSG_END uint16 = 0x2000
)
