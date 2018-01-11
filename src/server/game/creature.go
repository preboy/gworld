package game

import (
	"server/game/battle"
)

// 对象属性
type Creature struct {
	battle.Property
}

func NewCreature(cid uint32) *Creature {
	return nil
}

// ==================================================

func (self *Creature) ToCreature() *Hero {
	return nil
}

func (self *Creature) ToPlayer() *Creature {
	return self
}

func (self *Creature) UnitType() UnitType {
	return UnitType_Creature
}

func (self *Creature) ToBattleUnit() *battle.BattleUnit {
	return nil
}
