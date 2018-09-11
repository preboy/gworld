package player

import (
	"fmt"
)

import (
	_ "core/timer"
)

func (self *Player) CreateTimer(i uint64, r bool, f func()) uint64 {
	return self.timerMgr.CreateTimer(i, r, f)
}

func (self *Player) NextTick(fn func()) {
	self.tf = append(self.tf, fn)
}

func (self *Player) CancelTimer(id uint64) {
	self.timerMgr.CancelTimer(id)
}

func (self *Player) OnTimer(id uint64) {
	defer self.do_next_tick()

	fmt.Println("Player.OnTimer:", id)
}
