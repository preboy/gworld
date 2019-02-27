package app

import (
	"server/battle"
	"server/config"
)

// 对象属性
type Creature struct {
	proto *config.Creature
}

func NewCreature(id, lv uint32) *Creature {
	proto := config.CreatureConf.Query(id, lv)
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
		Lv:       self.proto.Level,
		UnitType: uint32(self.UnitType()),
	}

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

	// 可见属性计算
	u.Prop_base = &battle.Property{
		Hp:   float64(self.proto.Hp),
		Apm:  float64(self.proto.Apm),
		Atk:  float64(self.proto.Atk),
		Def:  float64(self.proto.Def),
		Crit: float64(self.proto.Crit),
		Hurt: float64(self.proto.Hurt),
	}

	// 加成属性计算 TODO
	u.Prop_addi = &battle.Property{}

	return u
}

// ============================================================================

func CreatureTeamToBattleTroop(id uint32) *battle.BattleTroop {
	team := config.CreatureTeamConf.Query(id)
	if team == nil {
		return nil
	}

	var r11, r12, r13, r21, r22, r23 *battle.BattleUnit

	if len(team.Row11) > 0 {
		m := team.Row11[0]
		r11 = NewCreature(m.Id, m.Lv).ToBattleUnit()
	}

	if len(team.Row12) > 0 {
		m := team.Row12[0]
		r12 = NewCreature(m.Id, m.Lv).ToBattleUnit()
	}

	if len(team.Row13) > 0 {
		m := team.Row13[0]
		r13 = NewCreature(m.Id, m.Lv).ToBattleUnit()
	}

	if len(team.Row21) > 0 {
		m := team.Row21[0]
		r21 = NewCreature(m.Id, m.Lv).ToBattleUnit()
	}

	if len(team.Row22) > 0 {
		m := team.Row22[0]
		r22 = NewCreature(m.Id, m.Lv).ToBattleUnit()
	}

	if len(team.Row23) > 0 {
		m := team.Row23[0]
		r23 = NewCreature(m.Id, m.Lv).ToBattleUnit()
	}

	return battle.NewBattleTroop(r11, r12, r13, r21, r22, r23)
}
