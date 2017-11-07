package player

import (
	"encoding/binary"
	"fmt"
	"reflect"
)

import (
	"core/tcp"
	"public/protocol"
	"server/player/handlers"
)

var (
	protocol_request  [protocol.PROTO_END]reflect.Type
	protocol_response map[reflect.Type]protocol.PROTO_END
)

// 将Packet转化为Message
func (self *Player) OnRecvPacket(packet *tcp.Packet) {
	self.q_packets <- packet
}

func (self *Player) dispatch_packet() {
	busy := false
	for {
		select {
		case packet := q_packets:
			self.on_dispatch_packet(packet)
			busy = true
		default:
			break
		}
	}
	return busy
}

func (self *Player) on_dispatch_packet(packet *tcp.Packet) {
	// 根据opcode生成对应的结构体对象，跳动对象的Handler方法
	typ := protocol_request[packet.opcode]
	if typ {
		obj := reflect.New(typ).Interface().(*typ)
		// proto.Unmarshal(packet.data, obj)
		obj.OnRequest(self)
	} else {
		fmt.Println("unknown packet", packet.opcode)
	}
}

func register_protocol_request(opcode uint16, obj interface{}) {
	if opcode < protocol.PROTO_END {
		protocol_request[opcode] = reflect.TypeOf(obj)
	}
}

func register_protocol_response(opcode uint16, obj interface{}) {
	if opcode < protocol.PROTO_END {
		protocol_response[reflect.TypeOf(obj)] = opcode
	}
}

func init() {
	register_protocol_request(protocol.CS_LOGIN, handlers.Student{})
}

func init() {
	register_protocol_response(protocol.SC_LOGIN, handlers.StudentResp{})
}

func (self *Socket) SendPacket(opcode uint16, data []byte) {
}

func (self *Player) send(obj interface{}) {
	// data, err = proto.Marshal(&obj)
	// if err {
	// return
	// }

	typ := reflect.TypeOf(obj)
	opcode, ok := protocol_response[typ]
	if err == nil {
		buf := new(bytes.Buffer)
		binary.Write(buf, binary.LittleEndian, uint16(len(data)))
		binary.Write(buf, binary.LittleEndian, opcode)
		binary.Write(buf, binary.LittleEndian, data)
		plr.socket.Send(buf)
	}
}
