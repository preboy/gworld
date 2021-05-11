package gambler

import (
	"gworld/ddz/comp"
)

var (
	_gbr_mgr = &gambler_manager{
		gamblers: map[string]*Gambler{},
	}
)

// ----------------------------------------------------------------------------
// init

func init() {
	comp.Init_GamblerManager(_gbr_mgr)
}

// ----------------------------------------------------------------------------
// player_mgr

type gambler_manager struct {
	gamblers map[string]*Gambler
}

func (self *gambler_manager) NewGambler(pid string) comp.IGambler {
	if _, ok := self.gamblers[pid]; ok {
		return nil
	}

	plr := &Gambler{
		PID: pid,
	}

	plr.Init()

	self.gamblers[pid] = plr

	return plr
}

func (self *gambler_manager) FindGambler(pid string) comp.IGambler {
	return self.gamblers[pid]
}

func (self *gambler_manager) ExistGambler(name string) bool {
	for _, v := range self.gamblers {
		if v.Name == name {
			return true
		}
	}

	return false
}
