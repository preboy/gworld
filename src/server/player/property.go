package player

import ()

// 对象属性
type Property struct {
	Atk       uint32 // 攻击
	Def       uint32 // 防御
	Apm       uint32 // 手速
	Hp_cur    uint32 // HP当前
	Hp_max    uint32 // HP上限
	Crit      uint32 // 暴击
	Crit_hurt uint32 // 暴伤
}
