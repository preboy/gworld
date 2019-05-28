package app

import (
	"game/battle"
	"game/config"
	"public/protocol/msg"
)

type Hero struct {
	// 这里的数据就是要存入DB的数据
	Id           uint32   `bson:"id"`             // 配置表ID
	Lv           uint32   `bson:"lv"`             // 等级(决定基础属性)
	Talent       uint32   `bson:"talent"`         // 天赋等级
	Aptitude     uint32   `bson:"aptitude"`       // 资质等级
	RefineLv     uint32   `bson:"refine_lv"`      // 精炼等级(额外提升属性)
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
	proto := config.HeroConf.Query(id)
	if proto == nil {
		return nil
	}

	hero := &Hero{
		Id: id,
		Lv: 1,
	}

	if proto.Skill1 != 0 {
		hero.Active[0] = Skill{
			Id: proto.Skill1,
			Lv: 1,
		}
	}
	if proto.Skill2 != 0 {
		hero.Active[1] = Skill{
			Id: proto.Skill2,
			Lv: 1,
		}
	}

	return hero
}

// ============================================================================

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
	proto := config.HeroConf.Query(self.Id)
	return proto.Name
}

func (self *Hero) ToBattleUnit() *battle.BattleUnit {
	u := &battle.BattleUnit{
		Base:     self,
		Id:       self.Id,
		Lv:       self.Lv,
		UnitType: uint32(self.UnitType()),
		Prop:     battle.NewPropertyGroup(),
	}

	// ------------------------------------------------------------------------
	conf := config.HeroConf.Query(self.Id)

	// 装入速度
	u.Prop.AddProp(1, 0, float64(conf.Apm))

	// 装入属性
	for {
		// 等级基础属性
		prop_conf := config.HeroPropConf.Query(self.Id, self.Lv)
		u.Prop.AddProps(prop_conf.Props)

		// 资质加成
		if self.Aptitude > 0 {
			rate := float64(self.Aptitude) / 100
			for _, v := range prop_conf.Props {
				u.Prop.AddProp(v.Id, 2, v.Val*rate)
			}
		}

		break
	}

	// 天赋
	for {
		if self.Talent > 0 {
			talent_conf := config.HeroTalentConf.Query(conf.Talent, self.Talent)
			u.Prop.AddProps(talent_conf.Props)
		}

		break
	}

	// 精炼
	for {
		if self.RefineLv == 0 {
			break
		}

		if self.RefineSuper {
			proto := config.RefineSuperConf.Query(self.RefineLv)
			u.Prop.AddProps(proto.Props)
		} else {
			proto := config.RefineNormalConf.Query(self.RefineLv)
			u.Prop.AddProps(proto.Props)
		}

		break
	}

	// 被动技能
	for {
		for i := 0; i < 4; i++ {
			v := &self.Passive[i]
			proto := config.SkillProtoConf.Query(v.Id, v.Lv)
			if proto != nil {
				u.Prop.AddProps(proto.Prop_Passive)
			}
		}

		break
	}

	// ------------------------------------------------------------------------
	// 装入技能

	// 普攻
	for {
		proto := config.HeroConf.Query(self.Id)
		if len(proto.SkillCommon) > 0 {
			sc := proto.SkillCommon[0]
			u.Skill_comm = battle.NewBattleSkill(sc.Id, sc.Lv)
		}
		break
	}

	// 主动技能
	for i := 0; i < 2; i++ {
		v := &self.Active[i]
		skill := battle.NewBattleSkill(v.Id, v.Lv)
		if skill != nil {
			u.Skill_battle = append(u.Skill_battle, skill)
		}
	}

	// 被动技能
	for i := 0; i < 4; i++ {
		v := &self.Passive[i]
		proto := config.SkillProtoConf.Query(v.Id, v.Lv)
		if proto != nil {
			u.Skill_Passive = append(u.Skill_Passive, proto)
		}
	}

	return u
}

func (self *Hero) ToMsg() *msg.Hero {
	conf := config.HeroConf.Query(self.Id)
	hero := &msg.Hero{
		Id:           self.Id,
		Lv:           self.Lv,
		Apm:          conf.Apm,
		Talent:       self.Talent,
		Aptitude:     self.Aptitude,
		RefineLv:     self.RefineLv,
		RefineTimes:  self.RefineTimes,
		RefineSuper:  self.RefineSuper,
		Power:        self.Power,
		Status:       self.Status,
		LifePoint:    self.LifePoint,
		LifePointMax: self.LifePointMax,
	}

	for i := 0; i < 2; i++ {
		hero.Active = append(hero.Active, &msg.Skill{
			Id: self.Active[i].Id,
			Lv: self.Active[i].Lv,
		})
	}

	for i := 0; i < 4; i++ {
		hero.Passive = append(hero.Passive, &msg.Skill{
			Id: self.Passive[i].Id,
			Lv: self.Passive[i].Lv,
		})
	}

	return hero
}
