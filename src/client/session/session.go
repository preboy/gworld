package session

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/rand"
	"time"
)

import (
	"github.com/gogo/protobuf/proto"
)

import (
	"core/tcp"
	"core/timer"
	"public/protocol"
	"public/protocol/msg"
)

type Session struct {
	socket    *tcp.Socket
	q_packets chan *tcp.Packet
	timerMgr  *timer.TimerMgr

	tid_ping uint64
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

func (self *Session) SendPacket(opcode uint16, obj proto.Message) {
	data, err := proto.Marshal(obj)
	if err == nil {
		l := uint16(len(data))
		b := make([]byte, 0, l+2+2)
		buf := bytes.NewBuffer(b)
		binary.Write(buf, binary.LittleEndian, uint16(len(data)))
		binary.Write(buf, binary.LittleEndian, opcode)
		binary.Write(buf, binary.LittleEndian, data)
		self.socket.Send(buf.Bytes())
	} else {
		fmt.Println("SendPacket Error:failed to Marshal obj")
	}
}

func (self *Session) init() {
	self.timerMgr = timer.NewTimerMgr(self)
	self.tid_ping = self.timerMgr.CreateTimer(3000, true, nil)
}

func (self *Session) update() {
	self.timerMgr.Update()
}

func (self *Session) on_packet(packet *tcp.Packet) {

}

func (self *Session) OnTimer(id uint64) {
	if id == self.tid_ping {
		req := &msg.PingRequest{}
		r := rand.Uint32()
		req.Time = r
		self.SendPacket(protocol.MSG_PING, req)
		fmt.Println("It's time to ping", r)
	}
}
