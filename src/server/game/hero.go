package game

import (
	"server/game/battle"
)

type Hero struct {
	// 这里的数据就是要存入DB的数据
	Pid        uint32       `bson:pid"`          // 配置表ID
	Uid        uint32       `bson:uid"`          // 唯一ID
	Level      uint32       `bson:"level"`       // 等级
	Power      uint32       `bson:"power"`       // 战斗力
	Equips     [4]Equipment `bson:"equip"`       // 武器、护甲、血玉、手套
	Skills     [2]Skill     `bson:"skills"`      // 技能(主动)
	Auras      [2]Aura      `bson:"auras"`       // 光环技能(被动)
	Status     uint32       `bson:"status"`      // 当前状态：0:闲置军营 1:战舰出征
	StatusData uint32       `bson:"status_data"` // 与当前状态相关的数据
	Dead       bool         `bson:"dead"`        // 是否死亡
}

func NewHero(cid uint32) *Hero {
	return nil
}

// ==================================================

func (self *Hero) ToCreature() *Creature {
	return nil
}

func (self *Hero) ToPlayer() *Hero {
	return self
}

func (self *Hero) UnitType() UnitType {
	return UnitType_Hero
}

func (self *Hero) ToBattleUnit() *battle.BattleUnit {
	return nil
}
