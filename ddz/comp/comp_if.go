package comp

import (
	"gworld/core/tcp"

	"github.com/gogo/protobuf/proto"
)

// ----------------------------------------------------------------------------
// local

type IPlayer interface {
	GetPID() string
	GetName() string

	SetSession(tcp.ISession)

	OnLogin()
	OnLogout()

	OnPacket(packet *tcp.Packet)
	SendMessage(msg IMessage)
}

type IGambler interface {
	IPlayer
}

type IReferee interface {
	IPlayer
}

type IMatch interface {
	GetMID() uint32
	GetName() string

	IsOver() bool

	OnUpdate()
	OnMessage(pid string, req IMessage, res IMessage)

	Sit(pid string) bool
}

type IMessage interface {
	proto.Message
	GetOP() int32
}
