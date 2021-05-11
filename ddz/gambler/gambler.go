package gambler

import (
	"gworld/core/log"
	"gworld/core/tcp"
	"gworld/core/utils"
	"gworld/ddz/comp"
	"gworld/ddz/loop"

	"github.com/gogo/protobuf/proto"
)

var (
	_gbrs = map[string]*Gambler{}
)

// ----------------------------------------------------------------------------
// init

func init() {
	loop.Register(func() {
		for _, plr := range _gbrs {
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

// ----------------------------------------------------------------------------
// Gambler

type Gambler struct {
	PID  string
	Name string
	Data *gambler_data
	Sess tcp.ISession
}

func (self *Gambler) OnLogin() {
	_gbrs[self.PID] = self
}

func (self *Gambler) OnLogout() {
	self.Sess = nil
	delete(_gbrs, self.PID)
}

func (self *Gambler) OnUpdate() {
}

func (self *Gambler) OnPacket(packet *tcp.Packet) {
	e, ok := _msg_executor[int32(packet.Opcode)]
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

	e.h(self, req, res)

	str = utils.ObjectToString(res)
	log.Info("SEND packet: %s, %d, %s", self.PID, packet.Opcode, str)

	self.SendMessage(res)
}

// ----------------------------------------------------------------------------
// member
func (self *Gambler) Init() {
	self.Data = &gambler_data{}
}

func (self *Gambler) GetPID() string {
	return self.PID
}

func (self *Gambler) GetName() string {
	return self.Name
}

func (self *Gambler) SetSession(sess tcp.ISession) {
	self.Sess = sess
}

func (self *Gambler) SendMessage(msg comp.IMessage) {
	self.SendProtobufMessage(uint16(msg.GetOP()), msg)
}

func (self *Gambler) SendProtobufMessage(opcode uint16, msg proto.Message) {
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
