package app

import (
	"gworld/game/battle"
	"gworld/game/config"
)

// 对象属性
type Creature struct {
	proto *config.Creature
}

func NewCreature(id uint32) *Creature {
	proto := config.CreatureConf.Query(id)
	if proto != nil {
		return &Creature{
			proto: proto,
		}
	}
	return nil
}

// ============================================================================

func (self *Creature) ToCreature() *Creature {
	return self
}

func (self *Creature) ToPlayer() *Hero {
	return nil
}

func (self *Creature) UnitType() UnitType {
	return UnitType_Creature
}

func (self *Creature) Name() string {
	return self.proto.Name
}

func (self *Creature) ToBattleUnit() *battle.BattleUnit {
	u := &battle.BattleUnit{
		Base:     self,
		Id:       self.proto.Id,
		UnitType: uint32(self.UnitType()),
		Prop:     battle.NewPropertyGroup(),
	}

	// ------------------------------------------------------------------------
	// 装入属性

	for { // 等级
		u.Prop.AddProps(self.proto.Props)
		break
	}

	// ------------------------------------------------------------------------
	// 装入技能

	// 普攻
	if len(self.proto.SkillCommon) > 0 {
		s := self.proto.SkillCommon[0]
		u.Skill_comm = battle.NewBattleSkill(s.Id, s.Lv)
	}

	// 技能
	for _, v := range self.proto.SkillExtra {
		s := battle.NewBattleSkill(v.Id, v.Lv)
		if s != nil {
			u.Skill_battle = append(u.Skill_battle, s)
		}
	}

	return u
}

// ============================================================================

func CreatureTeamToBattleTroop(id uint32) *battle.BattleTroop {
	team := config.CreatureTeamConf.Query(id)
	if team == nil {
		return nil
	}

	var r11, r12, r13, r21, r22, r23 *battle.BattleUnit

	if c := NewCreature(team.Row11); c != nil {
		r11 = c.ToBattleUnit()
	}

	if c := NewCreature(team.Row12); c != nil {
		r12 = c.ToBattleUnit()
	}

	if c := NewCreature(team.Row13); c != nil {
		r13 = c.ToBattleUnit()
	}

	if c := NewCreature(team.Row21); c != nil {
		r21 = c.ToBattleUnit()
	}

	if c := NewCreature(team.Row22); c != nil {
		r22 = c.ToBattleUnit()
	}

	if c := NewCreature(team.Row23); c != nil {
		r23 = c.ToBattleUnit()
	}

	return battle.NewBattleTroop(r11, r12, r13, r21, r22, r23)
}
