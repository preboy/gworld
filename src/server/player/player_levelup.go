package player

import (
	_ "core/log"
	"public/protocol"
	"public/protocol/msg"
	"server/config"
)

func (self *Player) GetLevel() uint32 {
	return self.data.Level
}

func (self *Player) AddExp(exp uint64) {
	if exp == 0 {
		return
	}

	self.data.Exp += exp

	old_lv := self.data.Level

	for {
		conf := config.LevelupConf.Query(self.data.Level)
		if conf == nil {
			break
		}

		if conf.Exp == 0 {
			self.data.Exp = 0
			break
		}

		if self.data.Exp >= conf.Exp {
			self.data.Exp -= conf.Exp
			self.data.Level++
		} else {
			break
		}
	}

	// notice
	res := &msg.PlayerLvExpUpdate{
		Lv:  self.data.Level,
		Exp: self.data.Exp,
	}

	self.SendPacket(protocol.MSG_SC_PlayerLvExpUpdate, res)

	// upgrade event
	new_lv := self.data.Level
	if old_lv != new_lv {
		self.on_levelup(old_lv, new_lv)
	}
}

func (self *Player) on_levelup(old_lv, new_lv uint32) {
}
