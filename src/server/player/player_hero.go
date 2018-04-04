package player

import (
	"public/protocol"
	"public/protocol/msg"
	"server/game"
)

func (self *Player) GetHero(id uint32) *game.Hero {
	hero, _ := self.data.Heros[id]
	return hero
}

func (self *Player) AddHero(id uint32) bool {
	hero, _ := self.data.Heros[id]
	if hero != nil {
		return true
	} else {
		hero = game.NewHero(id)
	}

	if hero == nil {
		return false
	}

	self.data.Heros[id] = hero
	return true
}

func (self *Player) UpdateHeroToClient(id uint32) {
	res := msg.HeroInfoUpdateResponse{}
	res.Hero = self.GetHero(id).ToMsg()
	self.SendPacket(protocol.MSG_SC_HeroInfoUpdate, &res)
}
