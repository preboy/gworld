package player

import (
	"core/event"
	"game/constant"
)

func init() {
	event.On(event.EVT_SCHED_DAY, func(evt uint32, args ...interface{}) {
		EachOnlinePlayer(func(plr *Player) {
			event.Fire(constant.Evt_Plr_DataReset, plr)
		})
	})
}
