package comp

import (
	"gworld/core/tcp"
)

// ----------------------------------------------------------------------------
// local

type ISession interface {
	SendPacket(opcode uint16, data []byte)
}

type IPlayer interface {
	GetID() string

	OnLogin()
	OnLogout()
	OnPacket(packet *tcp.Packet)
}
