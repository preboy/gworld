package server

import (
	"server/player"
)

func (self *Server) on_new_hour() {
	player.EachPlayer(func(plr *player.Player) {
		println("dd")
	})
}

func (self *Server) on_new_day() {

}

func (self *Server) on_new_week() {

}

func (self *Server) on_new_month() {

}

func (self *Server) on_new_year() {

}
