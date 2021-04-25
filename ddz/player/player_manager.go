package player

import (
	"gworld/ddz/comp"
)

var (
	_plr_mgr = &player_manager{}
)

// ----------------------------------------------------------------------------
// init

func init() {
	comp.Init_PlayerManager(_plr_mgr)
}

// ----------------------------------------------------------------------------
// player_mgr

type player_manager struct {
}

func (self *player_manager) FindPlayer(pid string) comp.IPlayer {
	return nil
}
