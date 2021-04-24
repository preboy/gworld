package player

import (
	"gworld/core/log"
	"gworld/core/tcp"
	"gworld/core/utils"
	"gworld/ddz/comp"
	"gworld/ddz/loop"

	"github.com/gogo/protobuf/proto"
)

var (
	_plrs = map[string]*Player{}
)

// ----------------------------------------------------------------------------
// init

func init() {
	loop.Register(func() {
		for _, plr := range _plrs {
			plr.OnUpdate()
		}
	})
}

// ----------------------------------------------------------------------------
// export

func Init() {
}

func Release() {
}

func NewPlayer(pid string) *Player {
	plr := &Player{
		PID: pid,
	}

	return plr
}

// ----------------------------------------------------------------------------
// Player

type Player struct {
	PID  string
	Data *PlayerData
	Sess comp.ISession
}

func (self *Player) OnLogin() {
	_plrs[self.PID] = self
}

func (self *Player) OnLogout() {
	delete(_plrs, self.PID)
}

func (self *Player) OnUpdate() {
}

func (self *Player) OnPacket(packet *tcp.Packet) {
	e, ok := _msg_executor[int(packet.Opcode)]
	if !ok {
		log.Warning("Unknown packet : %s %d", self.PID, packet.Opcode)
		return
	}

	req, res := e.c()

	err := proto.Unmarshal(packet.Data, req)
	if err != nil {
		log.Error("proto.Unmarshal ERROR: %s %d", self.PID, packet.Opcode)
		return
	}

	str := utils.ObjectToString(req)
	log.Info("RECV packet: %s, %d, %s", self.PID, packet.Opcode, str)

	e.h(req, res)

	str = utils.ObjectToString(res)
	log.Info("SEND packet: %s, %d, %s", self.PID, packet.Opcode, str)

	self.SendMessage(2, res)
}

func (self *Player) SendMessage(opcode uint16, msg proto.Message) {
	if self.Sess == nil {
		return
	}

	data, err := proto.Marshal(msg)
	if err == nil {
		log.Error("proto.Marshal ERROR: %s %d", self.PID, opcode)
		return
	}

	self.Sess.SendPacket(opcode, data)
}

func (self *Player) SetSession(sess comp.ISession) {
	self.Sess = sess
}
