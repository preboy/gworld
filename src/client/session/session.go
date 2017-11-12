package session

import (
	"time"
)
import (
	"core/tcp"
)

type Session struct {
	socket    *tcp.Socket
	q_packets chan *tcp.Packet
}

func NewSession() *Session {
	return &Session{
		q_packets: make(chan *tcp.Packet, 0x100),
	}
}

func (self *Session) SetSocket(s *tcp.Socket) {
	self.socket = s
}

func (self *Session) OnRecvPacket(packet *tcp.Packet) {
	self.q_packets <- packet
}

func (self *Session) Go() {
	go func() {
		for {
			select {
			case packet := self.q_packets:
				self.on_packet(packet)
			default:
				time.Sleep(20 * time.Millisecond)
				self.update()
			}
		}
	}()
}

func (self *Session) on_packet(packet *tcp.Packet) {

}

func (self *Session) update() {

}
