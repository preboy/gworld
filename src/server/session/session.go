package session

import (
	"fmt"
	"time"
)

import (
	"core/tcp"
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

func NewSession() *Session {
	return &Session{}
}

func (self *Session) SetSocket(socket *tcp.Socket) *Session {
	self.socket = socket
}

func (self *Session) SetPlayer(player *player.Player) *Session {
	self.player = player
}

func (self *Session) OnRecvPacket(packet *tcp.Packet) {
	self.last_touch = time.Now().Unix()

	if packet.opcode == protocol.CS_PING {
		self.on_ping(packet)
		return
	}

	if self.player != nil {
		self.player.OnRecvPacket(packet)
		return
	}

	switch packet.opcode {
	case protocol.CS_LOGIN:
		self.on_login(packet)
	case protocol.CS_ENTER_GAME:
		self.on_enter_game(packet)
	default:
		fmt.Println("unknown packet in session:", packet.opcode)
	}
}

func (self *Session) Send(data []byte) {
	self.socket.Send(data)
}

// ------------------ session handler ------------------

// 心跳包
func (self *Session) on_ping(packet *tcp.Packet) {
	fmt.Println("on_ping")
}

// 登录
func (self *Session) on_login(packet *tcp.Packet) {
	fmt.Println("on_login")
	// self.verify = true
	// self.account = "zcg"
}

// 进入游戏
func (self *Session) on_enter_game(packet *tcp.Packet) {
	fmt.Println("on_enter_game")
	if !self.verify {
		return
	}

	player.EnterGame(self.account, self)
}
