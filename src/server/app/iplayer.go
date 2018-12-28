package app

import (
	"github.com/gogo/protobuf/proto"
)

type IPlayer interface {
	GetId() string

	GetItem(id uint32) uint64
	SetItem(id uint32, cnt uint64)
	AddItem(id uint32, cnt uint64)
	SubItem(id uint32, cnt uint64)

	SendPacket(opcode uint16, obj proto.Message)
}
