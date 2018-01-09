package game

// 对象属性
type Creature struct {
	Property
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

func (self *Creature) UnitType() int {
	return UT_Creature
}

func (self *Creature) ToBattleUnit() *BattleUnit {
	return nil
}
