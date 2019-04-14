package session

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"regexp"
	"sync/atomic"
	"time"

	"github.com/gogo/protobuf/proto"

	"core/event"
	"core/log"
	"core/tcp"
	"game/loop"
	"game/player"
	"public/ec"
	"public/protocol"
	"public/protocol/msg"
)

var count_of_session uint32

type Session struct {
	socket *tcp.Socket
	player *player.Player

	// session data
	auth    bool   // 验证是否通过
	account string // 账号名
}

var (
	_re *regexp.Regexp
)

func init() {
	_re = regexp.MustCompile("^test[0-9]{3}$")
}

func NewSession() *Session {
	return &Session{}
}

func (self *Session) SetSocket(socket *tcp.Socket) {
	self.socket = socket
}

func (self *Session) SetPlayer(player *player.Player) {
	self.player = player
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
		fmt.Println("SendPacket Error: failed to Marshal obj")
	}
}

func (self *Session) Disconnect() {
	if self.socket != nil {
		self.socket.Stop()
		self.socket = nil
	}

	self.player = nil
}

// ============================================================================

func (self *Session) OnOpened() {
	atomic.AddUint32(&count_of_session, 1)
}

func (self *Session) OnClosed() {
	atomic.AddUint32(&count_of_session, ^uint32(0))
	if self.player != nil {
		self.player.Stop()
	}
}

// session interface impl
func (self *Session) OnRecvPacket(packet *tcp.Packet) {
	if packet.Opcode == uint16(protocol.MSG_CS_PingRequest) {
		self.on_ping(packet)
		return
	}

	if self.player != nil {
		loop.Get().PostPacket(self, packet)
		return
	}

	if packet.Opcode == protocol.MSG_CS_LoginRequest {
		self.on_auth(packet)
	} else {
		fmt.Println("unknown packet in session:", packet.Opcode)
	}
}

// running in main loop
func (self *Session) DoPacket(packet *tcp.Packet) {
	self.player.OnRecvPacket(packet)
}

// ============================================================================
//  session handler

// 心跳包
func (self *Session) on_ping(packet *tcp.Packet) {
	req := msg.PingRequest{}
	res := msg.PingResponse{}
	proto.Unmarshal(packet.Data, &req)

	res.Time = req.Time
	self.SendPacket(protocol.MSG_SC_PingResponse, &res)
}

// 登录
func (self *Session) on_auth(packet *tcp.Packet) {
	req := msg.LoginRequest{}
	res := msg.LoginResponse{}
	proto.Unmarshal(packet.Data, &req)

	if self.auth {
		return
	}

	// TODO: should go to auth server to verify
	res.ErrorCode = ec.Login_Failed

	if _re.MatchString(req.Acct) {
		if req.Pass == "1" {
			self.auth = true
			self.account = req.Acct
			res.ErrorCode = ec.OK
		}
	}

	self.SendPacket(protocol.MSG_SC_LoginResponse, &res)

	loop.Get().PushEvent(event.NewEvent(constant.Evt_Auth, self.account, self))
	// plrmgr.OnLogin(self.account)

	log.Debug("on_login: acct=%s, pass=%s, ok=%d", req.Acct, req.Pass, res.ErrorCode)
}
