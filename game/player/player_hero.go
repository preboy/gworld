package player

import (
	"gopkg.in/mgo.v2/bson"

	"gworld/game/app"
	"gworld/game/battle"
	"gworld/public/ec"
	"gworld/public/protocol"
	"gworld/public/protocol/msg"
)

// ============================================================================
// marshal

type hero_map_t map[uint32]*app.Hero

type hero_t struct {
	Id   uint32
	Hero *app.Hero
}

func (self hero_map_t) GetBSON() (interface{}, error) {
	var arr []*hero_t

	for k, v := range self {
		arr = append(arr, &hero_t{
			Id:   k,
			Hero: v,
		})
	}

	return arr, nil
}

func (self *hero_map_t) SetBSON(raw bson.Raw) error {
	var arr []*hero_t

	err := raw.Unmarshal(&arr)
	if err != nil {
		return err
	}

	*self = make(hero_map_t)
	for _, v := range arr {
		(*self)[v.Id] = v.Hero
	}

	return nil
}

// ============================================================================

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
