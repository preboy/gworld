package session

import (
	"core/tcp"
)

type Session struct {
	socket *tcp.Socket
}

func NewSession() *Session {
	return &Session{}
}

func (self *Session) SetSocket(s *tcp.Socket) {
	self.socket = s
}

func (self *Session) OnRecvPacket(packet *tcp.Packet) {

}
