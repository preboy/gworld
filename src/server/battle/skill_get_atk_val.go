package battle

import (
	"core/log"
	"server/config"
)

func (self *BattleSkill) get_attack_for_target_major() float64 {
	atk := self.caster.Prop.Value(PropType_Atk)
	atk *= self.proto.Ratio_major
	atk += self.get_extra_atk(self.proto.Extra_major)
	return atk
}

func (self *BattleSkill) get_attack_for_target_minor() float64 {
	atk := self.caster.Prop.Value(PropType_Atk)
	atk *= self.proto.Ratio_minor
	atk += self.get_extra_atk(self.proto.Extra_minor)
	return atk
}

// ----------------------------------------------------------------------------

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

	return 0
}
