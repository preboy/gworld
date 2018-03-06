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

func (self *Creature) Name() string {
	return self.proto.Name
}

func (self *Creature) ToBattleUnit() *battle.BattleUnit {
	u := &battle.BattleUnit{
		Base:       self,
		UnitType:   uint32(self.UnitType()),
		Troop:      nil,
		Dead:       false,
		Skill_curr: nil,
	}

	// 攻速
	// u.Rest_time_last = uint32(60000 / self.proto.Apm)

	// 普攻
	if len(self.proto.Skill_common) > 0 {
		sc := self.proto.Skill_common[0]
		u.Skill_comm = battle.NewSkillBattle(sc.Id, sc.Lv)
	}

	// 技能
	for _, v := range self.proto.Skill_extra {
		skill := battle.NewSkillBattle(v.Id, v.Lv)
		if skill != nil {
			u.Skill_extra = append(u.Skill_extra, skill)
		}
	}

	// 光环
	for _, v := range self.proto.Auras {
		aura := battle.NewAuraBattle(v.Id, v.Lv)
		if aura != nil {
			u.Auras = append(u.Auras, aura)
		}
	}

	// 基本属性
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

// ==================================================

func CreatureTeamToBattleTroop(id uint32) *battle.BattleTroop {
	team := config.GetCreatureTeamConf().GetCreatureTeam(id)
	if team == nil {
		return nil
	}

	var lp, rp, c, lg, rg *battle.BattleUnit

	if len(team.L_Pioneer) > 0 {
		m := team.L_Pioneer[0]
		lp = NewCreature(m.Id, m.Lv).ToBattleUnit()
	}

	if len(team.R_Pioneer) > 0 {
		m := team.R_Pioneer[0]
		rp = NewCreature(m.Id, m.Lv).ToBattleUnit()
	}

	if len(team.Commander) > 0 {
		m := team.Commander[0]
		c = NewCreature(m.Id, m.Lv).ToBattleUnit()
	}

	if len(team.L_Guarder) > 0 {
		m := team.L_Guarder[0]
		lg = NewCreature(m.Id, m.Lv).ToBattleUnit()
	}

	if len(team.R_Guarder) > 0 {
		m := team.R_Guarder[0]
		rg = NewCreature(m.Id, m.Lv).ToBattleUnit()
	}

	return battle.NewBattleTroop(lp, rp, c, lg, rg)
}
