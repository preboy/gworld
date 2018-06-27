package battle

import (
	"server/game/config"
)

type Property struct {
	Hp   float64 // HP上限
	Apm  float64 // 速度
	Atk  float64 // 攻击
	Def  float64 // 防御
	Crit float64 // 暴击
	Hurt float64 // 暴伤
}

type PropType uint32

const (
	_                PropType = 0 + iota // 角色属性类型
	PropType_HP                          // 1 HP
	PropType_Apm                         // 2 速度
	PropType_Atk                         // 3 攻击
	PropType_Def                         // 4 防御
	PropType_Crit                        // 5 暴击
	PropType_Hurt                        // 6 暴伤
	PropType_PctHP                       // 7  HP加成
	PropType_PctApm                      // 8  速度加成
	PropType_PctAtk                      // 9  攻击加成
	PropType_PctDef                      // 10 防御加成
	PropType_PctCrit                     // 11 暴击加成
	PropType_PctHurt                     // 12 暴伤加成
)

func (self *Property) Clear() {
	self.Hp = 0
	self.Apm = 0
	self.Atk = 0
	self.Def = 0
	self.Crit = 0
	self.Hurt = 0
}

func (self *Property) AddConf(props []*config.PropConf) {
	for _, v := range props {
		switch PropType(v.Id) {
		case PropType_HP:
			{
				self.Hp += v.Val
			}
		case PropType_Apm:
			{
				self.Apm += v.Val
			}
		case PropType_Atk:
			{
				self.Atk += v.Val
			}
		case PropType_Def:
			{
				self.Def += v.Val
			}
		case PropType_Crit:
			{
				self.Crit += v.Val
			}
		case PropType_Hurt:
			{
				self.Hurt += v.Val
			}
		default:
		}
	}
}

func (self *Property) AddProperty(prop *Property) {
	self.Hp += prop.Hp
	self.Apm += prop.Apm
	self.Atk += prop.Atk
	self.Def += prop.Def
	self.Crit += prop.Crit
	self.Hurt += prop.Hurt
}
