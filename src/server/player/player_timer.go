package player

import (
	_ "core/timer"
)

func (self *Player) CreateTimer(i uint64, r bool, f func()) uint64 {
	return self.timerMgr.CreateTimer(i, r, f)
}

func (self *Player) CancelTimer(id uint64) {
	self.timerMgr.CancelTimer(id)
}

func (self *Player) OnTimer(id uint64) {

}
