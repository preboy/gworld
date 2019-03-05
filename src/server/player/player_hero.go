package player

import (
	"public/protocol"
	"public/protocol/msg"
	"server/app"
	"server/battle"
)

func (self *Player) GetHero(id uint32) *app.Hero {
	hero, _ := self.data.Heros[id]
	return hero
}

func (self *Player) AddHero(id uint32) bool {
	hero, _ := self.data.Heros[id]
	if hero != nil {
		return true
	} else {
		hero = app.NewHero(id)
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

// 检测队伍的合法性
func (self *Player) IsValidTeam(team []uint32) (ret bool, heros []*app.Hero) {
	size := len(team)
	if size > battle.MAX_TROOP_MEMBER {
		return
	}

	ids := map[uint32]int{}

	for _, id := range team {
		if id == 0 {
			continue
		}
		if ids[id] != 0 {
			return
		}
		ids[id]++
	}

	if len(ids) == 0 {
		return
	}

	for _, id := range ids {
		hero := self.GetHero(id)
		if hero == nil {
			return
		}
		heros = append(heros, hero)
	}

	ret = true

	return
}

func (self *Player) CreateBattleTroop(team []uint32) *battle.BattleTroop {
	// 找出英雄
	ok, heros := self.IsValidTeam(team)
	if !ok {
		return nil
	}

	// todo

	// return NewBattleTroop(members ...*BattleUnit)
}
