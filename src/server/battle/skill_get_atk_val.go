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

func (self *BattleSkill) get_extra_atk(confs []*config.ExtraAtkConf) (ret float64) {
	for _, conf := range confs {
		switch conf.Typ {
		case 1:
			ret += conf.Val
		case 2:
			ret += conf.Val * self.caster.Prop.Value(PropType_Atk) / 100
		default:
			log.Error("unknown extra atk type = %d", conf.Typ)
		}
	}

	return ret
}
