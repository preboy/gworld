package session

import (
	"fmt"
	"time"
)

import (
	"core/tcp"
	"public/protocol"
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
