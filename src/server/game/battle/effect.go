package battle

import (
	"core/log"
	"fmt"
)

// ============================================================

type Effect interface {
	OnStart(ab *AuraBattle)
	OnUpdate(ab *AuraBattle)
	OnFinish(ab *AuraBattle)
	OnEvent(evt BattleEvent, ab *AuraBattle, sc *SkillContext)
}

type effect_creator = func() Effect

var ecs map[uint32]effect_creator

func init() {
	ecs = make(map[uint32]effect_creator, 0x100)
}

func RegisterEffect(i uint32, f effect_creator) {
	if _, ok := ecs[i]; ok {
		log.Warning("dup effect_creator: id = ", i)
	}
	ecs[i] = f
}

func LoadEffects() {
	// 一行一行往下累积
	RegisterEffect(1, NewEffect_1)
}

func NewEffect(i uint32) Effect {
	if fn, ok := ecs[i]; ok {
		return fn()
	}
	return nil
}

// ============================================================
/*
	这里写注释
*/
type Effect_1 struct {
}

func (self *Effect_1) OnStart(ab *AuraBattle) {
	fmt.Println("eff 1 OnStart")
}

func (self *Effect_1) OnUpdate(ab *AuraBattle) {
	fmt.Println("eff 1 OnUpdate")
}

func (self *Effect_1) OnFinish(ab *AuraBattle) {
	fmt.Println("eff 1 OnFinish")
}

func (self *Effect_1) OnEvent(evt BattleEvent, ab *AuraBattle, sc *SkillContext) {

}

func NewEffect_1() Effect {
	return &Effect_1{}
}

// ============================================================
/*
	这里写注释
*/
