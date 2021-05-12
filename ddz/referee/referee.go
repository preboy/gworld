package referee

import (
	"gworld/core/log"
	"gworld/core/tcp"
	"gworld/core/utils"
	"gworld/ddz/comp"
	"gworld/ddz/loop"

	"github.com/gogo/protobuf/proto"
)

var (
	_rfrs = map[string]*Referee{}
)

// ----------------------------------------------------------------------------
// init

func init() {
	loop.Register(func() {
		for _, rfr := range _rfrs {
			rfr.OnUpdate()
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
// Referee

type Referee struct {
	PID  string
	Name string
	Data *referee_data
	Sess tcp.ISession
}

func (self *Referee) OnLogin() {
	_rfrs[self.PID] = self
	log.Info("referee login")
}

func (self *Referee) OnLogout() {
	self.Sess = nil
	delete(_rfrs, self.PID)

	log.Info("referee logout")
}

func (self *Referee) OnUpdate() {
}

func (self *Referee) OnPacket(packet *tcp.Packet) {
	e, ok := _msg_executor[int32(packet.Opcode)]
	if !ok {
		log.Warning("referee Unknown packet : %v %d", self.PID, packet.Opcode)
		return
	}

	req, res := e.c()

	err := proto.Unmarshal(packet.Data, req)
	if err != nil {
		log.Error("referee proto.Unmarshal ERROR: %v, %v, %v", self.PID, packet.Opcode, err)
		return
	}

	str := utils.ObjectToString(req)
	log.Info("referee RECV packet: %v, %v, %v", self.PID, req.GetOP(), str)

	e.h(self, req, res)

	str = utils.ObjectToString(res)
	log.Info("referee SEND packet: %v, %v, %v", self.PID, res.GetOP(), str)

	self.SendMessage(res)
}

// ----------------------------------------------------------------------------
// member

func (self *Referee) GetPID() string {
	return self.PID
}

func (self *Referee) GetName() string {
	return self.Name
}

func (self *Referee) SetSession(sess tcp.ISession) {
	self.Sess = sess
}

func (self *Referee) SendMessage(msg comp.IMessage) {
	self.SendProtobufMessage(uint16(msg.GetOP()), msg)
}

func (self *Referee) SendProtobufMessage(opcode uint16, msg proto.Message) {
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
