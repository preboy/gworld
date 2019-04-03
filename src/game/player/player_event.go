package player

import (
	"fmt"

	"core/event"
	"core/log"
	"core/utils"
	"game/constant"
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

	defer func() {
		if err := recover(); err != nil {
			log.Error("PANIC on 'OnEvent':", self.GetId(), evt.Id, evt.Args)
			log.Error("STACK TRACE:", utils.Callstack())
		}
	}()

	fmt.Println("Player.OnEvent:", evt)

	switch evt.Id {
	case event.EVT_SCHED_SYNC_CALL:
		{
			if f, ok := evt.Args[0].(func()); ok {
				f()
			}
		}
		// 击杀怪物
	case constant.EVT_PLR_KILL_MONSTER:
		{
			if mid, ok := evt.Args[0].(int32); ok {
				self.data.Quest.OnKill(mid)
			}
		}
	default:
	}

	// fallthrough to other modules

	utils.ExecuteSafely(func() {
		self.data.Growth.OnEvent(evt)
	})
}
