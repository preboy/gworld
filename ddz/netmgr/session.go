package netmgr

import (
	"bytes"
	"encoding/binary"
	"strconv"
	"sync"
	"sync/atomic"

	"gworld/core/tcp"
	"gworld/ddz/loop"
	"gworld/ddz/player"
)

var (
	_seq      = uint32(1)
	_sessions = map[uint32]*session{}
	_lock     = sync.Mutex{}
	_chunks   = make(chan *chunk, 0x1000)
)

type chunk struct {
	s *session
	p *tcp.Packet
}

type session struct {
	Id     uint32
	socket *tcp.Socket
	player *player.Player
}

// ----------------------------------------------------------------------------

func NewSession() *session {
	seq := atomic.AddUint32(&_seq, 1)
	return &session{
		Id: seq,
	}
}

func (self *session) SetSocket(socket *tcp.Socket) {
	self.socket = socket
}

func (self *session) SetPlayer(player *player.Player) {
	self.player = player
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

func (self *session) OnOpened() {
	_lock.Lock()
	defer _lock.Unlock()

	_sessions[self.Id] = self
}

func (self *session) OnClosed() {
	_lock.Lock()
	defer _lock.Unlock()

	delete(_sessions, self.Id)

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
	if self.player != nil {
		_chunks <- &chunk{self, packet}
	} else {
		pid := strconv.Itoa(int(self.Id))
		plr := player.NewPlayer(pid)
		self.SetPlayer(plr)

		loop.Post(func() {
			plr.OnLogin()
		})
	}
}

// ----------------------------------------------------------------------------
// local

func update_chunks() {
	for {
		select {
		case c := <-_chunks:
			c.s.player.OnPacket(c.p)
			loop.DoNext()
		default:
			return
		}
	}
}
