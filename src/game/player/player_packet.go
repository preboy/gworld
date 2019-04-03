package player

import (
	"core/log"
	"core/tcp"
	"core/utils"
)

type msg_func = func(*Player, *tcp.Packet)

var (
	_funcs = map[uint16]msg_func{}
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
		func() {
			defer func() {
				if err := recover(); err != nil {
					log.Error("PANIC on 'on_packet':", self.GetId(), packet.Opcode, packet.Data)
					log.Error("STACK TRACE:", utils.Callstack())
				}
			}()
			f(self, packet)
		}()
	} else {
		log.Warning("!!! Unkonwn packat: id = %v", packet.Opcode)
	}
}

func register_handler(opcode uint16, f msg_func) {
	_funcs[opcode] = f
}
