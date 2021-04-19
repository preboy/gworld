package player

import (
	"gworld/public/protocol"
	"gworld/public/protocol/msg"
)

func (self *Player) SendNotice(notice string, flag int32) {
	res := &msg.NoticeUpdate{}
	res.Flag = flag
	res.Notice = notice
	self.SendPacket(protocol.MSG_SC_NoticeUpdate, res)
}
