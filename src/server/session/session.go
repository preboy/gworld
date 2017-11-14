package session

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"regexp"
	"time"
)

import (
	"github.com/gogo/protobuf/proto"
)

import (
	"core/log"
	"core/tcp"
	"public/protocol"
	"public/protocol/msg"
	"server/err"
	"server/player"
)

type Session struct {
	socket *tcp.Socket
	player *player.Player

	// session data
	verify     bool   // 验证是否通过
	account    string // 账号名
	last_touch int64  // 心跳时间
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

func (self *Session) OnRecvPacket(packet *tcp.Packet) {
	self.last_touch = time.Now().Unix()

	if packet.Opcode == uint16(protocol.MSG_PING) {
		self.on_ping(packet)
		return
	}

	if self.player != nil {
		self.player.OnRecvPacket(packet)
		return
	}

	switch packet.Opcode {
	case protocol.MSG_LOGIN:
		self.on_login(packet)
	case protocol.MSG_ENTER_GAME:
		self.on_enter_game(packet)
	default:
		fmt.Println("unknown packet in session:", packet.Opcode)
	}
}

func (self *Session) Send(data []byte) {
	self.socket.Send(data)
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
		self.Send(buf.Bytes())
	} else {
		fmt.Println("SendPacket Error:failed to Marshal obj")
	}
}

func (self *Session) Disconnect() {
	self.player = nil
	if self.socket != nil {
		self.socket.Stop()
		self.socket = nil
	}
}

// ------------------ session handler ------------------

// 心跳包
func (self *Session) on_ping(packet *tcp.Packet) {
	req := msg.PingRequest{}
	res := msg.PingResponse{}
	proto.Unmarshal(packet.Data, &req)

	res.Time = req.Time
	self.SendPacket(packet.Opcode, &res)
	fmt.Println("session: on_ping", req.Time)
}

// 登录
func (self *Session) on_login(packet *tcp.Packet) {
	req := msg.LoginRequest{}
	res := msg.LoginResponse{}
	proto.Unmarshal(packet.Data, &req)
	// TODO something
	res.ErrorCode = err.ERR_LOGIN_FAILED

	if _re.MatchString(req.Acct) {
		if req.Pass == "1" {
			self.verify = true
			res.ErrorCode = err.ERR_OK
		}
	}
	self.SendPacket(packet.Opcode, &res)
	log.GetLogger().Debug("on_login: acct=%s, pass=%s, ok=%d", req.Acct, req.Pass, res.ErrorCode)
}

// 进入游戏
func (self *Session) on_enter_game(packet *tcp.Packet) {
	fmt.Println("on_enter_game")
	if !self.verify {
		return
	}
	player.EnterGame(self.account, self)
}
