package player

import (
	"core/event"
)

func (self *Player) EmitEvent(evt *event.Event) {
	self.evtMgr.Fire(evt)
}

func (self *Player) OnEvent(evt *event.Event) {

}
