package player

import (
	// "core/log"

	"core/event"
	"game/config"
	"game/constant"
	"public/protocol"
	"public/protocol/msg"
)

func (self *Player) GetLevel() uint32 {
	return self.data.Lv
}

func (self *Player) AddExp(exp uint64) {
	if exp == 0 {
		return
	}

	self.data.Exp += exp
	old_lv := self.data.Lv

	for {
		conf := config.LevelupConf.Query(self.data.Lv)
		if conf == nil {
			break
		}

		if conf.Exp == 0 {
			self.data.Exp = 0
			break
		}

		if self.data.Exp >= conf.Exp {
			self.data.Exp -= conf.Exp
			self.data.Lv++
		} else {
			break
		}
	}

	// notice
	self.SendPacket(protocol.MSG_SC_PlayerExpUpdate, &msg.PlayerExpUpdate{
		Lv:  self.data.Lv,
		Exp: self.data.Exp,
	})

	// level upgrade event
	new_lv := self.data.Lv
	if old_lv != new_lv {
		event.Fire(constant.Evt_Plr_LevelUp, old_lv, new_lv)
	}
}
