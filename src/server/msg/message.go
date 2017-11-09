package msg

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
)

import (
	"core/tcp"
	"github.com/gogo/protobuf/proto"
	"public/protocol"
	"server/player"
	"server/player/msg/handler"
)

var (
	map_msgid_type = [protocol.PROTO_END]*msg_type{}
	map_type_msgid = make(map[reflect.Type]protocol.ProtoID)
)

type msg_type struct {
	Req, Res reflect.Type
}

type IPlayerMessage interface {
	Send(data []byte)
}

// type IMessage interface {
// 	OnRequest(plr *Player) bool
// }

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
	register_protocol(protocol.CS_LOGIN, handler.Student{}, handler.StudentResp{})
}

func OnPacket(packet *tcp.Packet, plr IPlayerMessage) {
	// 根据opcode生成对应的结构体对象，跳动对象的Handler方法
	mt := map_msgid_type[packet.Opcode]
	if mt != nil {
		req := reflect.New(mt.Req).Interface().(*mt.Req)
		res := reflect.New(mt.Res).Interface().(*mt.Res)
		proto.Unmarshal(packet.data, req)

		if req.OnRequest(plr, res) {
			data, err := proto.Marshal(res)
			if err != nil {
				fmt.Println("fuck")
			} else {
				plr.Send(packet.Opcode, data)
			}
		}
	} else {
		fmt.Println("unknown packet", packet.Opcode)
	}
}

// func SendPacket(obj interface{}) {
// 	data, err := proto.Marshal(&obj)
// 	if err != nil {
// 		return
// 	}

// 	typ := reflect.TypeOf(obj)
// 	opcode, ok := map_type_msgid[typ]
// 	if ok {
// 		buf := new(bytes.Buffer)
// 		binary.Write(buf, binary.LittleEndian, uint16(len(data)))
// 		binary.Write(buf, binary.LittleEndian, opcode)
// 		binary.Write(buf, binary.LittleEndian, data)
// 		self.Send(buf.Bytes())
// 	} else {
// 		fmt.Println("invalid obj")
// 	}
// }
