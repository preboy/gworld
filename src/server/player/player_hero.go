package player

import ()

type Equipment struct {
	Quality uint32 `bson:quality"` // 品质
	Level   uint32 `bson:level"`   // 等级
}

// 主动技能
type Skill struct {
	Id    uint32 `bson:id"`    // ID
	Level uint32 `bson:level"` // 等级
}

// 被动技能
type Passivity struct {
	Id    uint32 `bson:id"`    // ID
	Level uint32 `bson:level"` // 等级
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
	Skills     [2]Skill     `bson:"Spell"`       // 技能
	Passives   [2]Passivity `bson:"Spell"`       // 被动技能
	Status     uint32       `bson:"status"`      // 当前状态：0:闲置军营 1:战舰出征
	StatusData uint32       `bson:"status_data"` // 与当前状态相关的数据
	Dead       bool         `bson:"dead"`        // 是否死亡

	// 非存盘数据
	atk       uint32  `bson:"-"` // 攻击
	def       uint32  `bson:"-"` // 防御
	apm       uint32  `bson:"-"` // 攻速
	hp_cur    uint32  `bson:"-"` // hp  hp主要用于远程战舰战斗
	hp_max    uint32  `bson:"-"` // hpmax
	crit      uint32  `bson:"-"` // 暴击
	crit_hurt uint32  `bson:"-"` // 暴伤百分比
	auras     []*Aura `bson:"-"` // 光环
}
