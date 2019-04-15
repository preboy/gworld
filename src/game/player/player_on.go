package player

import (
	"core/event"
)

func init() {
	event.On(event.EVT_SCHED_HOUR, func(id uint32, args ...interface{}) {
		EachOnlinePlayer(func(plr *Player) {
			plr.on_new_hour()
		})
	})

	event.On(event.EVT_SCHED_DAY, func(id uint32, args ...interface{}) {
		EachOnlinePlayer(func(plr *Player) {
			plr.on_new_day()
		})
	})

	event.On(event.EVT_SCHED_WEEK, func(id uint32, args ...interface{}) {
		EachOnlinePlayer(func(plr *Player) {
			plr.on_new_week()
		})
	})

	event.On(event.EVT_SCHED_MONTH, func(id uint32, args ...interface{}) {
		EachOnlinePlayer(func(plr *Player) {
			plr.on_new_month()
		})
	})

	event.On(event.EVT_SCHED_YEAR, func(id uint32, args ...interface{}) {
		EachOnlinePlayer(func(plr *Player) {
			plr.on_new_year()
		})
	})
}

func (self *Player) on_new_hour() {
}

func (self *Player) on_new_day() {
}

func (self *Player) on_new_week() {
}

func (self *Player) on_new_month() {
}

func (self *Player) on_new_year() {
}
