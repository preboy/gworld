package handlers

import (
	"src/server/player"
)

type Message interface {
	OnRequest(plr *player.Player)
}

type Student struct {
	name string
	age  uint32
	man  bool
}

type StudentResp struct {
	fuck string
}

func (self *Student) OnRequest(plr *player.Player) {
	resp := StudentResp{}
	plr.send(&resp)
}
