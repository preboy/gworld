package session

import (
	"time"
)
import (
	"core/tcp"
	"core/timer"
)

type Session struct {
	socket    *tcp.Socket
	q_packets chan *tcp.Packet
	timeMgr   *timer.TimerMgr
}

func NewSession() *Session {
	s := &Session{
		q_packets: make(chan *tcp.Packet, 0x100),
	}

	s.init()
	return s
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
			case packet := <-self.q_packets:
				self.on_packet(packet)
			default:
				time.Sleep(20 * time.Millisecond)
				self.update()
			}
		}
	}()
}

func (self *Session) init() {
	self.timeMgr = timer.NewTimerMgr(self)
}

func (self *Session) on_packet(packet *tcp.Packet) {

}

func (self *Session) OnTimer(id uint64) {

}

func (self *Session) update() {

}
