package player

import (
	"fmt"
)

import (
	"core/event"
)

func (self *Player) EmitEvent(evt *event.Event) {
	self.evtMgr.Fire(evt)
}

func (self *Player) FireEvent(evt *event.Event) {
	self.OnEvent(evt)
}

func (self *Player) OnEvent(evt *event.Event) {
	fmt.Println("OnPlayerEvent:", evt)
}
