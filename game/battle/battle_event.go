package battle

import (
	"gworld/public/protocol/msg"
)

// 释放技能
func (self *Battle) BattlePlayEvent_Cast(who *BattleUnit, id, lv uint32, targets []uint32) {
	self.Result.Skill = append(self.Result.Skill, &msg.BattleEventSkill{
		Time:    self.time,
		Caster:  who.Pos,
		Skill:   &msg.BattleSkill{id, lv},
		Targets: targets,
	})
}

// 光环事件
func (self *Battle) BattlePlayEvent_Aura(who *BattleUnit, caster *BattleUnit, id, lv uint32, obtain bool) {
	self.Result.Aura = append(self.Result.Aura, &msg.BattleEventAura{
		Time:   self.time,
		Owner:  who.Pos,
		Caster: caster.Pos,
		Aura:   &msg.BattleAura{id, lv},
		Obtain: obtain,
	})
}

// 伤害事件
func (self *Battle) BattlePlayEvent_Hurt(who *BattleUnit, target *BattleUnit, hurt uint32, crit uint32, typ uint32) {
	self.Result.Hurt = append(self.Result.Hurt, &msg.BattleEventHurt{
		Time:   self.time,
		Caster: who.Pos,
		Target: target.Pos,
		Hurt:   hurt,
		Crit:   crit,
		Type:   typ,
	})
}

// 光环效果事件
func (self *Battle) BattlePlayEvent_Effect(who *BattleUnit, caster *BattleUnit, typ AuraEffectType, arg1, arg2, arg3, arg4 int32) {
	self.Result.Effect = append(self.Result.Effect, &msg.BattleEventAuraEffect{
		Time:   self.time,
		Owner:  who.Pos,
		Caster: caster.Pos,
		Type:   uint32(typ),
		Arg1:   arg1,
		Arg2:   arg2,
		Arg3:   arg3,
		Arg4:   arg4,
	})
}
