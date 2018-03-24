package battle

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
