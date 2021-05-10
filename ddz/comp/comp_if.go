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
	SendMessage(msg IMessage)
}

type IMatch interface {
	GetMID() uint32
	GetName() string

	IsOver() bool

	OnUpdate()
	OnMessage(pid string, req IMessage, res IMessage)
}

type IMessage interface {
	proto.Message
	GetOP() int32
}
