package app

import (
	"github.com/gogo/protobuf/proto"

	"game/battle"
)

type IPlayer interface {
	GetId() string
	GetLevel() uint32

	GetItem(id uint32) uint64
	SetItem(id uint32, cnt uint64)
	AddItem(id uint32, cnt uint64)
	SubItem(id uint32, cnt uint64)

	CreateBattleTroop([]uint32) (*battle.BattleTroop, int)

	SendPacket(opcode uint16, obj proto.Message)
}
