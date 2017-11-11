package player

import (
	"github.com/gogo/protobuf/proto"
)

type ISession interface {
	SendPacket(uint16, proto.Message)
	SetPlayer(*Player)
	Disconnect()
}

func (self *Player) SetSession(s ISession) {
	self.s = s
}

func (self *Player) Disconnect() {
	self.s.Disconnect()
	self.s = nil
}

func (self *Player) SendPacket(opcode uint16, obj proto.Message) {
	if self.s != nil {
		self.s.SendPacket(opcode, obj)
	}
}
