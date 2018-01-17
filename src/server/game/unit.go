package game

type UnitType uint32

const (
	UnitType_Hero UnitType = 1 + iota
	UnitType_Creature
)

type Unit interface {
	GetPower() uint32
	ToCreature() *Creature
	ToPlayer() *Hero
	UnitType() int
	Name() string
}

// 主动技能
type Skill struct {
	Id       uint32 `bson:id"`        // ID
	Level    uint32 `bson:level"`     // 等级
	EffectId uint32 `bson:effect_id"` // 技能附加效果ID
}

// 被动光环
type Aura struct {
	Id    uint32 `bson:id"`    // ID
	Level uint32 `bson:level"` // 等级
}
