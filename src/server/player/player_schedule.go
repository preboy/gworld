package player

import (
	"core/event"
)

func (self *Player) OnSchedule(evt *event.Event) {
	self.evtMgr.Fire(evt)
}
