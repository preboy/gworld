package player

import (
	"gworld/core/log"
	"gworld/core/tcp"
)

type msg_func = func(*Player, *tcp.Packet)

var (
	_funcs = map[uint16]msg_func{}
)

// 将Packet转化为Message
func (self *Player) OnRecvPacket(packet *tcp.Packet) {
	fn := _funcs[packet.Opcode]

	if fn != nil {
		fn(self, packet)
	} else {
		log.Warning("!!! Unkonwn packat: id = %v", packet.Opcode)
	}
}

func register_handler(opcode uint16, f msg_func) {
	_funcs[opcode] = f
}
