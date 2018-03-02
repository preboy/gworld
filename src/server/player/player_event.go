package player

import (
	"fmt"
)

import (
	"core/event"
)

// called by other go routine to push event to player.Go
func (self *Player) FireEvent(evt *event.Event) {
	self.evtMgr.Fire(evt)
}

func (self *Player) OnEvent(evt *event.Event) {
	fmt.Println("Player.OnEvent:", evt)

	if evt.Id == event.EVT_SCHED_SYNC_CALL {
		f := evt.Ptr.(func())
		f()
	}

}
