package player

import (
	"core/tcp"
	"server/msg"
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
			msg.OnPacket(packet, self)
			busy = true
		default:
			break
		}
	}
	return busy
}
