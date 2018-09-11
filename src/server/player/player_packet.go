package player

import (
	"core/tcp"
	"public/protocol"
)

type msg_func = func(*Player, *tcp.Packet)

var (
	_funcs = [protocol.MSG_END]msg_func{}
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
			self.on_packet(packet)
			busy = true
		default:
			return busy
		}
	}
	return busy
}

func (self *Player) on_packet(packet *tcp.Packet) {
	defer self.do_next_tick()

	f := _funcs[packet.Opcode]
	if f != nil {
		f(self, packet)
	}
}

func register_handler(opcode uint16, f msg_func) {
	if opcode >= protocol.MSG_END {
		return
	}
	_funcs[opcode] = f
}
