package game

import (
	"server/game/battle"
	"server/game/config"
)

type Hero struct {
	// 这里的数据就是要存入DB的数据
	Id           uint32   `bson:id"`              // 配置表ID
	Level        uint32   `bson:"level"`          // 等级(决定基础属性)
	Exp          uint32   `bson:"exp"`            // 当前经验
	Refine       uint32   `bson:"refine_lv"`      // 精炼等级(额外提升属性)
	RefineTimes  uint32   `bson:"refine_times"`   // 普通精炼失败次数
	RefineSuper  bool     `bson:"refine_super"`   // 是否超级精炼(超级精炼失败则精炼等级归0，且失败无次数累计，但属性更强)
	Active       [2]Skill `bson:"skill_active"`   // 技能(主动)
	Passive      [4]Skill `bson:"skill_passive"`  // 技能(主动)
	Power        uint32   `bson:"power"`          // 战斗力
	Status       uint32   `bson:"status"`         // 当前状态：0:闲置军营 1:战舰出征
	LifePoint    uint32   `bson:"life_point"`     // 生命点数(外战中每死一回掉点，为0时无法再出战)
	LifePointMax uint32   `bson:"life_point_max"` // 生命点数上限
}

func NewHero(id uint32) *Hero {
	proto := config.GetHeroProto(id, 1)
	if proto == nil {
		return nil
	}

	hero := &Hero{
		Id:    id,
		Level: 1,
	}
	return hero
}

// ==================================================

func (self *Hero) ToCreature() *Creature {
	return nil
}

func (self *Hero) ToPlayer() *Hero {
	return self
}

func (self *Hero) UnitType() UnitType {
	return UnitType_Hero
}

func (self *Hero) Name() string {
	proto := config.GetHeroProto(self.Id, self.Level)
	return proto.Name
}

func (self *Hero) ToBattleUnit() *battle.BattleUnit {
	u := &battle.BattleUnit{
		Base:       self,
		Id:         self.Id,
		Lv:         self.Level,
		UnitType:   uint32(self.UnitType()),
		Troop:      nil,
		Skill_curr: nil,
	}

	proto := config.GetHeroProto(self.Id, self.Level)

	// 可见属性计算
	u.Prop_base = &battle.Property{
		Hp:       proto.Hp,
		Atk:      proto.Atk,
		Def:      proto.Def,
		Crit:     proto.Crit,
		CritHurt: proto.Crit_hurt,
	}

	// 普攻
	if len(proto.Skill_common) > 0 {
		sc := proto.Skill_common[0]
		u.Skill_comm = battle.NewBattleSkill(sc.Id, sc.Lv)
	}

	// 主动技能
	for i := 0; i < 2; i++ {
		v := &self.Active[i]
		skill := battle.NewBattleSkill(v.Id, v.Level)
		if skill != nil {
			u.Skill_exclusive = append(u.Skill_exclusive, skill)
		}
	}

	// 被动技能
	for i := 0; i < 4; i++ {
		v := &self.Passive[i]
		proto := config.GetSkillProto(v.Id, v.Level)
		if proto != nil {
			if proto.Passive == 1 {
				u.Prop_base.AddAttrs(proto.Attrs)
			} else {
				skill := battle.NewBattleSkill(v.Id, v.Level)
				if skill != nil {
					u.Skill_exclusive = append(u.Skill_exclusive, skill)
				}
			}
		}
	}

	// 装备精炼
	// 英雄专精

	// 加成属性计算
	u.Prop_addi = &battle.Property{}

	u.CalcProp()

	return u
}
