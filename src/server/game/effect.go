package game

import (
	"core/log"
	"fmt"
)

type EffectCtx struct {
	//
}

// ============================================================

type Effect interface {
	OnStart(ab *AuraBattle)
	OnUpdate(ab *AuraBattle)
	OnFinish(ab *AuraBattle)
	OnEvent(evt BattleEvent, ab *AuraBattle, sc *SkillContext)
}

type effect_func = func() Effect

var effects map[uint32]effect_func

func init() {
	effects = make(map[uint32]effect_func, 0x100)
}

func RegisterEffect(i uint32, f effect_func) {
	if _, ok := effects[i]; ok {
		log.Warning("dup eff: id = ", i)
	}
	effects[i] = f
}

func LoadEffects() {
	// 一行一行往下累积
	RegisterEffect(1, NewEffect_1)
}

func NewEffect(i uint32) Effect {
	if fn, ok := effects[i]; ok {
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
