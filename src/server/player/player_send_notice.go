package player

import (
	"public/protocol"
	"public/protocol/msg"
)

func (self *Player) SendNotice(notice string, flag int32) {
	res := &msg.NoticeUpdate{}
	res.Flag = flag
	res.Notice = notice
	self.SendPacket(protocol.MSG_SC_NoticeUpdate, res)
}
