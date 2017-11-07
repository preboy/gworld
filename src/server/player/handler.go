package clients

import (
	"fmt"
	"reflect"
)

// opcode

var CS_PING uint16 = 0x1
var SC_PING uint16 = 0x2

const (
	MAX_PACKET_COUNT = 0x1000
)

type HANDLER = func(packet *Packet)

type Message interface {
	OnRequest(plr *Player)
}

var msg_handler [MAX_PACKET_COUNT]reflect.Type

func RegisterMessageType(opcode uint16, msg Message) {
	msg_handler[opcode] = reflect.TypeOf(msg)
}

func Dispatcher(packet *Packet, plr *Player) {
	fmt.Println("new packet", packet.code, len(packet.data))
	ty := msg_handler[packet.code]
	if ty != nil {
		msg := reflect.New(ty)
		obj := msg.Interface().(Message)
		// proto.Marshal(msg.Interface().(Message))
		obj.OnRequest(plr)
	} else {
		fmt.Println("invalid packet ", packet.code, packet.data)
	}
}

type Student struct {
	name string
	age  uint32
	man  bool
}

func (self *Student) OnRequest(plr *Player) {

}

func init() {
	RegisterMessageType(CS_PING, &Student{})
}
