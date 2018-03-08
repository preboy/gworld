package battle

import (
	"core/log"
	"fmt"
	"server/game/config"
)

// ============================================================

type AuraScript interface {
	OnStart(ab *BattleAura)
	OnUpdate(ab *BattleAura)
	OnFinish(ab *BattleAura)
	OnEvent(evt BattleEvent, ab *BattleAura, ctx *SkillContext)
}

type script_creator = func() AuraScript

var _creators map[uint32]script_creator

func init() {
	RegisterAuraScripts()
}

func register_aura_script(id uint32, f script_creator) {
	if _, ok := _creators[id]; ok {
		log.Warning("dup script_creator: id = ", id)
	}
	_creators[id] = f
}

func RegisterAuraScripts() {
	_creators = make(map[uint32]script_creator, 0x100)

	// 一行一行往下累积
	register_aura_script(1, NewAuraScript_1)
}

func create_aura_script(proto *config.AuraProto) AuraScript {
	if f, ok := _creators[proto.Id]; ok {
		return f()
	} else {
		log.Error("Not Found AuraScript: AuraProto.Id = %v", proto.Id)
		return nil
	}
}

// ============================================================
/*
	这里写注释
*/
type AuraScript_1 struct {
	// 在这里存储每个光环自身的数据
}

func (self *AuraScript_1) OnStart(ab *BattleAura) {
	fmt.Println("AuraScript_1 OnStart")
}

func (self *AuraScript_1) OnUpdate(ab *BattleAura) {
	fmt.Println("AuraScript_1 OnUpdate")
}

func (self *AuraScript_1) OnFinish(ab *BattleAura) {
	fmt.Println("AuraScript_1 OnFinish")
}

func (self *AuraScript_1) OnEvent(evt BattleEvent, ab *BattleAura, ctx *SkillContext) {
	fmt.Println("AuraScript_1 Event:", evt)
}

func NewAuraScript_1() AuraScript {
	return &AuraScript_1{}
}

// ============================================================
/*
	这里写注释
*/
