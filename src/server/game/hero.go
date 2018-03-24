package game

import (
	"server/game/battle"
	"server/game/config"
)

type Hero struct {
	// 这里的数据就是要存入DB的数据
	Id         uint32       `bson:id"`           // 配置表ID
	Level      uint32       `bson:"level"`       // 等级
	Quality    uint32       `bson:"quality"`     // 品质
	Power      uint32       `bson:"power"`       // 战斗力
	Equips     [4]Equipment `bson:"equip"`       // 武器、护甲、血玉、手套
	Skills     [2]Skill     `bson:"skills"`      // 技能(主动)
	Auras      [2]Aura      `bson:"auras"`       // 光环技能(被动)
	Status     uint32       `bson:"status"`      // 当前状态：0:闲置军营 1:战舰出征
	StatusData uint32       `bson:"status_data"` // 与当前状态相关的数据
	Dead       bool         `bson:"dead"`        // 是否死亡
}

func NewHero(id uint32) *Hero {
	proto := config.GetHeroProtoConf().GetHeroProto(id, 1)
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
	proto := config.GetHeroProtoConf().GetHeroProto(self.Id, self.Level)
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

	proto := config.GetHeroProtoConf().GetHeroProto(self.Id, self.Level)

	// 普攻
	if len(proto.Skill_common) > 0 {
		sc := proto.Skill_common[0]
		u.Skill_comm = battle.NewBattleSkill(sc.Id, sc.Lv)
	}

	// 技能
	for i := 0; i < 2; i++ {
		v := &self.Skills[i]
		skill := battle.NewBattleSkill(v.Id, v.Level)
		if skill != nil {
			u.Skill_exclusive = append(u.Skill_exclusive, skill)
		}
	}

	// 可见属性计算
	u.Prop_base = &battle.Property{
		Hp:       proto.Hp,
		Atk:      proto.Atk,
		Def:      proto.Def,
		Crit:     proto.Crit,
		CritHurt: proto.Crit_hurt,
	}

	// 加成属性计算 TODO
	u.Prop_addi = &battle.Property{}

	u.CalcProp()

	return u
}
