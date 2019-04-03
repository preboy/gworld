package player

import (
	"core/event"
)

func (self *Player) OnSchedule(evt *event.Event) {
	self.evtMgr.Fire(evt)
}

func (self *Player) on_new_hour() {
	println("Player.on_new_hour")
}

func (self *Player) on_new_day() {
	println("Player.on_new_day")
}

func (self *Player) on_new_week() {
	println("Player.on_new_week")
}

func (self *Player) on_new_month() {
	println("Player.on_new_month")
}

func (self *Player) on_new_year() {
	println("Player.on_new_year")
}
