package server

import (
	"core/event"
	"server/player"
)

func (self *Server) on_new_hour() {
	evt := event.NewEvent(event.EVT_SCHED_HOUR, nil)
	player.EachPlayer(func(plr *player.Player) {
		plr.FireEvent(evt)
	})
}

func (self *Server) on_new_day() {
	evt := event.NewEvent(event.EVT_SCHED_DAY, nil)
	player.EachPlayer(func(plr *player.Player) {
		plr.FireEvent(evt)
	})
}

func (self *Server) on_new_week() {
	evt := event.NewEvent(event.EVT_SCHED_WEEK, nil)
	player.EachPlayer(func(plr *player.Player) {
		plr.FireEvent(evt)
	})
}

func (self *Server) on_new_month() {
	evt := event.NewEvent(event.EVT_SCHED_MONTH, nil)
	player.EachPlayer(func(plr *player.Player) {
		plr.FireEvent(evt)
	})
}

func (self *Server) on_new_year() {
	evt := event.NewEvent(event.EVT_SCHED_YEAR, nil)
	player.EachPlayer(func(plr *player.Player) {
		plr.FireEvent(evt)
	})
}
