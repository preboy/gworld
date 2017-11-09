package player

import (
	"encoding/binary"
	"fmt"
	"reflect"
)

import (
	"core/tcp"
	"public/protocol"
	// "server/player/msg"
)

type msg_type struct {
	Req, Res reflect.Type
}

type IMessage interface {
	OnRequest(plr *Player) bool
}

var (
	map_msgid_type = [protocol.PROTO_END]*msg_type{}
	map_type_msgid = make(map[reflect.Type]uint16)
)

// 将Packet转化为Message
func (self *Player) OnRecvPacket(packet *tcp.Packet) {
	self.q_packets <- packet
}

func (self *Player) dispatch_packet() bool {
	busy := false
	for {
		select {
		case packet := <-self.q_packets:
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
	mt := map_msgid_type[packet.Opcode]
	if mt != nil {
		req := reflect.New(mt.Req).Interface().(*mt.Req)
		res := reflect.New(mt.Res).Interface().(*mt.Res)
		// proto.Unmarshal(packet.data, obj)
		if req.OnRequest(self, res) {
			// data, err := proto.Marshal(res)
			// if err {
			// fmt.Println("fuck")
			// } else {
			// self.Send(data)
			// }
		}
	} else {
		fmt.Println("unknown packet", packet.Opcode)
	}
}

func register_protocol(opcode protocol.ProtoID, req interface{}, res interface{}) {
	if opcode >= protocol.PROTO_END {
		return
	}
	map_msgid_type[opcode] = &msg_type{
		reflect.TypeOf(req),
		reflect.TypeOf(res),
	}
	map_type_msgid[reflect.TypeOf(res)] = opcode
}

func init() {
	register_protocol(protocol.CS_LOGIN, msg.Student{}, msg.StudentResp{})
}

func (self *Player) SendPacket(obj interface{}) {
	// data, err = proto.Marshal(&obj)
	// if err {
	// return
	// }

	typ := reflect.TypeOf(obj)
	opcode, err := protocol_response[typ]
	if err == nil {
		buf := new(bytes.Buffer)
		binary.Write(buf, binary.LittleEndian, uint16(len(data)))
		binary.Write(buf, binary.LittleEndian, opcode)
		binary.Write(buf, binary.LittleEndian, data)
		self.Send(buf)
	} else {
		fmt.Println("invalid obj")
	}
}
