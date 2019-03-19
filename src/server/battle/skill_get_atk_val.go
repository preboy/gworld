package battle

import (
	"fmt"
)

func (self *BattleSkill) get_attack_value_first_target() float64 {
	atk := self.caster.Prop.Value(PropType_Atk)
	atk *= self.proto.Target_major
	atk += get_extra_atk(self.proto.Extra_major)
	return atk
}

func (self *BattleSkill) get_attack_value_second_target() float64 {
	atk := self.caster.Prop.Value(PropType_Atk)
	atk *= self.proto.Target_minor
	atk += get_extra_atk(self.proto.Extra_minor)
	return atk
}

func (self *BattleSkill) get_extra_atk(confs []*config.ExtraAtkConf) float64 {
	for _, conf := range confs {

		switch conf.Typ {
		case 1:
			return conf.Val
		case 2:
			return conf.Val * self.caster.Prop.Value(PropType_Atk)

		default:
			log.Error("known extra atk type", conf.Typ)
		}
	}
}
