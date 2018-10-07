package player

import (
	"fmt"

	"core/event"
	"core/utils"
)

// called by other go routine to push event to player.Go
func (self *Player) FireEvent(evt *event.Event) {
	self.evtMgr.Fire(evt)
}

func (self *Player) FireEventArgs(id uint32, args ...interface{}) {
	evt := event.NewEvent(id, args...)
	self.evtMgr.Fire(evt)
}

// called by player routine to exec event directly.
func (self *Player) CallEvent(evt *event.Event) {
	self.OnEvent(evt)
}

// events callback dispatcher
func (self *Player) OnEvent(evt *event.Event) {
	defer self.do_next_tick()

	fmt.Println("Player.OnEvent:", evt)

	if evt.Id == event.EVT_SCHED_SYNC_CALL {
		if f, ok := evt.Args[0].(func()); ok {
			f()
		}
	}

	// fallthrough to other modules

	utils.ExecuteSafely(func() {
		self.data.Growth.OnEvent(evt)
	})
}
