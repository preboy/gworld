package player

import (
	"core/log"
	"core/tcp"
	"core/utils"
	"game/microsvr"
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
		log.Warning("!!! Unkonwn packat: id = %v", packet.Opcode)
	}

	svr := microsvr.FindSvrByOpcode(packet.Opcode)
	if svr != nil {
		svr.PostPacket(self, packet)
	} else {
		self.DoPacket(packet)
	}
}

func (self *Player) DoPacket(packet *tcp.Packet) {
	self._plr_lock.Lock()
	defer self._plr_lock.Unlock()

	defer func() {
		if err := recover(); err != nil {
			log.Error("PANIC on 'on_packet':", self.GetId(), packet.Opcode, packet.Data)
			log.Error("STACK TRACE:", utils.Callstack())
		}
	}()

	_funcs[packet.Opcode](self, packet)
}

func register_handler(opcode uint16, f msg_func) {
	_funcs[opcode] = f
}
