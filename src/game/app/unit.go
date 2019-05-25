package app

type UnitType uint32

const (
	UnitType_Hero UnitType = 1 + iota
	UnitType_Creature
)

type Unit interface {
	GetPower() uint32
	ToCreature() *Creature
	ToPlayer() *Hero
	UnitType() UnitType
	Name() string
}

// 主动技能
type Skill struct {
	Id uint32 `bson:id"` // ID
	Lv uint32 `bson:lv"` // 等级
}

// 被动光环
type Aura struct {
	Id uint32 `bson:id"` // ID
	Lv uint32 `bson:lv"` // 等级
}
