package gambler

import (
	"gworld/ddz/comp"
)

var (
	_gbr_mgr = &gambler_manager{}
)

// ----------------------------------------------------------------------------
// init

func init() {
	comp.Init_GamblerManager(_gbr_mgr)
}

// ----------------------------------------------------------------------------
// player_mgr

type gambler_manager struct {
}

func (self *gambler_manager) NewGambler(pid string) comp.IGambler {
	plr := &Gambler{
		PID: pid,
	}

	return plr
}

func (self *gambler_manager) FindGambler(pid string) comp.IGambler {
	return nil
}
