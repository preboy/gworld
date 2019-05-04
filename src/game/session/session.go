package session

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"sync/atomic"

	"core"
	"core/log"
	"core/tcp"
	"core/utils"
	"game/app"
	"game/constant"
	"game/loop"
	"game/player"
	"github.com/gogo/protobuf/proto"
	"public/ec"
	"public/protocol"
	"public/protocol/msg"
)

var count_of_session uint32

type Session struct {
	socket *tcp.Socket
	player *player.Player
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
	}
}

// ============================================================================

func (self *Session) OnOpened() {
	atomic.AddUint32(&count_of_session, 1)
}

func (self *Session) OnClosed() {
	self.socket = nil

	atomic.AddUint32(&count_of_session, ^uint32(0))
	if self.player != nil {
		self.player.OnLogout()
		self.player = nil
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
	utils.ExecuteSafely(func() {
		self.player.OnRecvPacket(packet)
	})
}

// ============================================================================
//  session handler

// 心跳包
func (self *Session) on_ping(packet *tcp.Packet) {
	req := &msg.PingRequest{}
	res := &msg.PingResponse{}
	proto.Unmarshal(packet.Data, req)

	res.Time = req.Time
	self.SendPacket(protocol.MSG_SC_PingResponse, res)
}

// 登录
func (self *Session) on_auth(packet *tcp.Packet) {
	req := &msg.LoginRequest{}
	res := &msg.LoginResponse{}
	proto.Unmarshal(packet.Data, req)

	res.ErrorCode = ec.Login_Failed

	go func() {
		for {
			if self.player != nil {
				break
			}

			if !app.IsValidGameId(req.Svr) {
				break
			}

			conf := app.GetConfig()
			addr := fmt.Sprintf("http://%s:%d/auth?sdk=%s&pseudo=%s&token=%s&svr=%s",
				conf.Auth.Host,
				conf.Auth.Port,
				req.Sdk,
				req.Pseudo,
				req.Token,
				req.Svr,
			)

			ret := core.HttpGet(addr)

			var dat map[string]interface{}
			err := json.Unmarshal([]byte(ret), &dat)
			if err != nil {
				log.Error("From Auth json.Unmarshal err: %v", err)
				break
			}

			code := dat["code"].(float64)
			if int(code) == 0 {
				res.ErrorCode = ec.OK
				loop.Get().PostEventArgs(constant.Evt_Auth, req.Sdk, req.Pseudo, req.Svr, self)
			}

			log.Debug("on_auth: sdk=%s, acct=%s, Token=%s, svr=%s, ret=%d", req.Sdk, req.Pseudo, req.Token, req.Svr, int(code))

			break
		}

		self.SendPacket(protocol.MSG_SC_LoginResponse, res)
	}()
}
