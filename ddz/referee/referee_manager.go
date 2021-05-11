package referee

import (
	"gworld/ddz/comp"
)

var (
	_rfr_mgr = &referee_manager{}
)

// ----------------------------------------------------------------------------
// init

func init() {
	comp.Init_RefereeManager(_rfr_mgr)
}

// ----------------------------------------------------------------------------
// player_mgr

type referee_manager struct {
}

func (self *referee_manager) NewReferee(pid string) comp.IReferee {
	plr := &Referee{
		PID: pid,
	}

	return plr
}

func (self *referee_manager) FindReferee(pid string) comp.IReferee {
	return nil
}
