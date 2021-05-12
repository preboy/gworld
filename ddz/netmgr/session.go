package netmgr

import (
	"bytes"
	"encoding/binary"

	"gworld/core/tcp"
	"gworld/core/utils"
	"gworld/ddz/comp"
	"gworld/ddz/loop"
)

type session struct {
	Id     uint32
	mgr    *session_mgr_t
	socket *tcp.Socket
	player comp.IPlayer
}

// ----------------------------------------------------------------------------

func new_session() *session {
	return &session{
		Id: utils.SeqU32(),
	}
}

func (self *session) SetMgr(mgr *session_mgr_t) {
	self.mgr = mgr
}

func (self *session) SetSocket(socket *tcp.Socket) {
	self.socket = socket
}

func (self *session) SetPlayer(player comp.IPlayer) {
	self.player = player
	player.SetSession(self)
}

func (self *session) SendPacket(opcode uint16, data []byte) {
	if self.socket == nil {
		return
	}

	l := uint16(len(data))
	b := make([]byte, 0, l+2+2)
	buf := bytes.NewBuffer(b)
	binary.Write(buf, binary.LittleEndian, l)
	binary.Write(buf, binary.LittleEndian, opcode)
	binary.Write(buf, binary.LittleEndian, data)
	self.socket.Send(buf.Bytes())
}

func (self *session) Disconnect() {
	if self.socket != nil {
		self.socket.Stop()
		self.socket = nil
	}
}

// ----------------------------------------------------------------------------
// impl for ISession

func (self *session) OnOpened() {
	self.mgr.AddSession(self)
}

func (self *session) OnClosed() {
	self.mgr.DelSession(self)

	loop.Post(func() {
		if self.player != nil {
			self.player.OnLogout()
		}
	})

	self.socket = nil
	self.player = nil
}

// session interface impl
func (self *session) OnRecvPacket(packet *tcp.Packet) {
	self.mgr.OnRecvPacket(self, packet)
}
