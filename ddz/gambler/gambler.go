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
		log.Warning("gambler Unknown packet : %v %v", self.PID, packet.Opcode)
		return
	}

	req, res := e.c()

	err := proto.Unmarshal(packet.Data, req)
	if err != nil {
		log.Error("gambler proto.Unmarshal ERROR: %v %v", self.PID, packet.Opcode)
		return
	}

	str := utils.ObjectToString(req)
	log.Info("gambler RECV packet: %v, %v, %v", self.PID, req.GetOP(), str)

	e.h(self, req, res)

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
	str := utils.ObjectToString(msg)
	log.Info("gambler SEND packet: %v, %v, %v", self.PID, msg.GetOP(), str)

	self.SendProtobufMessage(uint16(msg.GetOP()), msg)
}

func (self *Gambler) SendProtobufMessage(opcode uint16, msg proto.Message) {
	if self.Sess == nil {
		return
	}

	data, err := proto.Marshal(msg)
	if err != nil {
		log.Error("proto.Marshal ERROR: %v %v %v", self.PID, opcode, err)
		return
	}

	self.Sess.SendPacket(opcode, data)
}
