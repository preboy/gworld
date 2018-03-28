package battle

import (
	"server/game/config"
)

type Property struct {
	Hp       uint32 // HP上限
	Atk      uint32 // 攻击
	Def      uint32 // 防御
	Crit     uint32 // 暴击
	CritHurt uint32 // 暴击伤害
}

type AttrType uint32

const (
	_                 AttrType = 0 + iota // 角色属性类型
	AttrType_HP                           // 1 HP
	AttrType_Atk                          // 2 攻击
	AttrType_Def                          // 3 防御
	AttrType_Crit                         // 4 暴击
	AttrType_CritHurt                     // 5 暴击伤害
)

func (self *Property) AddAttrs(attrs []*config.AttrConf) {
	for _, v := range attrs {
		switch AttrType(v.Id) {
		case AttrType_HP:
			{
				self.Hp += v.Val
			}
		case AttrType_Atk:
			{
				self.Atk += v.Val
			}
		case AttrType_Def:
			{
				self.Def += v.Val
			}
		case AttrType_Crit:
			{
				self.Crit += v.Val
			}
		case AttrType_CritHurt:
			{
				self.CritHurt += v.Val
			}
		default:
		}
	}
}

func (self *Property) AddProperty(p *Property) {
	self.Hp += p.Hp
	self.Atk += p.Atk
	self.Def += p.Def
	self.Crit += p.Crit
	self.CritHurt += p.CritHurt
}
