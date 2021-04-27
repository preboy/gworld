package comp

import (
	"gworld/core/tcp"

	"github.com/gogo/protobuf/proto"
)

// ----------------------------------------------------------------------------
// local

type ISession interface {
	SendPacket(opcode uint16, data []byte)
}

type IPlayer interface {
	GetPID() string

	OnLogin()
	OnLogout()

	OnPacket(packet *tcp.Packet)
	SendMessage(msg Message)
}

type Message interface {
	proto.Message
	GetOP() int32
}
