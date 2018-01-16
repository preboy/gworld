package game

import (
	"server/game/battle"
	"server/game/config"
)

// 对象属性
type Creature struct {
	proto *config.CreatureProto
}

func NewCreature(id, lv uint32) *Creature {
	proto := config.GetCreatureProtoConf().GetCreatureProto(id, lv)
	if proto != nil {
		return &Creature{
			proto: proto,
		}
	}
	return nil
}

// ==================================================

func (self *Creature) ToCreature() *Creature {
	return self
}

func (self *Creature) ToPlayer() *Hero {
	return nil
}

func (self *Creature) UnitType() UnitType {
	return UnitType_Creature
}

func (self *Creature) ToBattleUnit() *battle.BattleUnit {
	u := &battle.BattleUnit{
		UnitType:   uint32(self.UnitType()),
		Troop:      nil,
		Dead:       false,
		Skill_Curr: nil,
	}

	u.Prop = &battle.Property{
		Atk:       self.proto.Atk,
		Def:       self.proto.Def,
		Apm:       self.proto.Apm,
		Hp_cur:    self.proto.Hp,
		Hp_max:    self.proto.Hp,
		Crit:      self.proto.Crit,
		Crit_hurt: self.proto.Crit_hurt,
	}

	return u
}
