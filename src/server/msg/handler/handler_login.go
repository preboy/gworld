package handler

import (
	"server/player"
)

type Student struct {
	name string
	age  uint32
	man  bool
}

type StudentResp struct {
	fuck string
}

func (self *Student) OnRequest(plr *player.Player, res msg.IMessage) bool {
	resp := res.(*StudentResp)

	return true
}
