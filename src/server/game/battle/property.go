package battle

type Property struct {
	Atk       uint32 // 攻击
	Def       uint32 // 防御
	Hp_cur    uint32 // HP当前
	Hp_max    uint32 // HP上限
	Crit      uint32 // 暴击
	Crit_hurt uint32 // 暴伤
}

func (self *Property) GetPower() uint32 {
	return 2
}
