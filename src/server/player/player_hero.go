package player

import ()

// 属性
type Property struct {
	Atk       uint32 // 攻击
	Def       uint32 // 防御
	Apm       uint32 // 手速
	Hp_cur    uint32 // HP当前
	Hp_max    uint32 // HP上限
	Crit      uint32 // 暴击
	Crit_hurt uint32 // 暴伤
}

type Equipment struct {
	Quality uint32 `bson:quality"` // 品质
	Level   uint32 `bson:level"`   // 等级
}

// 主动技能
type Skill struct {
	Id     uint32 `bson:id"`      // ID
	Level  uint32 `bson:level"`   // 等级
	SpecId uint32 `bson:spec_id"` // 技能附加ID
}

type Aura struct {
	Id    uint32 `bson:id"`    // ID
	Level uint32 `bson:level"` // 等级
}

type Hero struct {
	// 这里的数据就是要存入DB的数据
	Pid        uint32       `bson:pid"`          // 配置表ID
	Uid        uint32       `bson:uid"`          // 唯一ID
	Level      uint32       `bson:"level"`       // 等级
	Power      uint32       `bson:"power"`       // 战斗力
	Equips     [4]Equipment `bson:"equip"`       // 武器、护甲、血玉、手套
	Skills     [2]Skill     `bson:"skills"`      // 技能(主动)
	Auras      [2]Aura      `bson:"auras"`       // 光环(被动)
	Status     uint32       `bson:"status"`      // 当前状态：0:闲置军营 1:战舰出征
	StatusData uint32       `bson:"status_data"` // 与当前状态相关的数据
	Dead       bool         `bson:"dead"`        // 是否死亡

	// 非存盘数据
	prob Property `bson:"-"` // 总体属性

	auras []*Aura `bson:"-"` // 光环
}
