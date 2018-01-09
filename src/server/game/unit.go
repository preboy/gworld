package game

const (
	UT_Hero = 1 + iota
	UT_Creature
)

type Unit interface {
	GetPower() uint32
	ToCreature() *Creature
	ToPlayer() *Hero
	UnitType() int
}
