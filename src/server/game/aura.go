package game

import (
	"server/game/config"
)

type Aura struct {
	Id    uint32 `bson:id"`    // ID
	Level uint32 `bson:level"` // 等级
}

type AuraBattle struct {
	owner  *BattleUnit
	caster *BattleUnit
	sp     *config.AuraProto
}

func NewAuraBattle(id, lv uint32) *AuraBattle {
	sp := config.GetAuraProtoConf().GetAuraProto(id, lv)
	if sp == nil {
		return nil
	}
	ab := &AuraBattle{
		sp: sp,
	}
	return ab
}

func (self *AuraBattle) Init(caster, owner *BattleUnit) {
	self.owner = owner
	self.caster = caster
}

func (self *AuraBattle) Update(time uint32) {
}

func (self *AuraBattle) IsFinish() bool {
	return true
}

func (self *AuraBattle) onStart() {
}

func (self *AuraBattle) onUpdate() {
}

func (self *AuraBattle) onFinish() {
}
