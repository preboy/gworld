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

// called by player routine to exec event directly.
func (self *Player) ExecEvent(evt *event.Event) {
	self.OnEvent(evt)
}

// events callback dispatcher
func (self *Player) OnEvent(evt *event.Event) {
	defer self.do_next_tick()

	fmt.Println("Player.OnEvent:", evt)

	if evt.Id == event.EVT_SCHED_SYNC_CALL {
		f := evt.Ptr.(func())
		f()
	}

}
