package player

import (
	"game/app"
	"game/battle"
	"public/ec"
	"public/protocol"
	"public/protocol/msg"
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
	res := &msg.HeroInfoUpdateResponse{}
	res.Hero = self.GetHero(id).ToMsg()
	self.SendPacket(protocol.MSG_SC_HeroInfoUpdateResponse, res)
}

func (self *Player) IsValidTeam(team []uint32) (ret bool, err int) {
	if len(team) > battle.MAX_TROOP_MEMBER {
		err = ec.BATTLE_Hero_Cnt_Exceed
		return
	}

	ids := map[uint32]int{}

	for _, id := range team {
		if id == 0 {
			continue
		}
		if ids[id] != 0 {
			err = ec.BATTLE_Hero_Present
			return
		}
		ids[id]++

		if self.GetHero(id) == nil {
			err = ec.BATTLE_Hero_NotExist
			return
		}
	}

	if len(ids) == 0 {
		err = ec.BATTLE_Hero_Zero
		return
	}

	ret = true
	err = ec.OK

	return
}

// 检测队伍的合法性
func (self *Player) CreateBattleTroop(team []uint32) (*battle.BattleTroop, int) {

	if ok, err := self.IsValidTeam(team); !ok {
		return nil, err
	}

	var units = make([]*battle.BattleUnit, battle.MAX_TROOP_MEMBER)

	for i, id := range team {
		if id == 0 {
			continue
		}

		units[i] = self.GetHero(id).ToBattleUnit()
	}

	return battle.NewBattleTroop(units...), ec.OK
}
