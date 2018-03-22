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
	OnEvent(evt BattleEventType, ab *BattleAura, ctx *SkillContext)
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
	register_aura_script(2001, NewAuraScript_2001) // 回春
	register_aura_script(2002, NewAuraScript_2002) // 回春
	register_aura_script(2003, NewAuraScript_2003) // 回春
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
	回血
*/
type AuraScript_2001 struct {
	// 在这里存储每个光环自身的数据
}

func (self *AuraScript_2001) OnStart(aura *BattleAura) {
	fmt.Println("AuraScript_2001 OnStart")
}

func (self *AuraScript_2001) OnUpdate(aura *BattleAura) {
	fmt.Println("AuraScript_2001 OnUpdate")
	aura.owner.Hp += uint32(aura.proto.Param1)
	aura.owner.AddCampaignDetail(CampaignEvent_AuraEffect, int32(ProertyHP), aura.proto.Param1, 0, 0)
}

func (self *AuraScript_2001) OnFinish(aura *BattleAura) {
	fmt.Println("AuraScript_2001 OnFinish")
}

func (self *AuraScript_2001) OnEvent(evt BattleEventType, aura *BattleAura, ctx *SkillContext) {
	fmt.Println("AuraScript_2001 Event:", evt)
}

func NewAuraScript_2001() AuraScript {
	return &AuraScript_2001{}
}

// ============================================================
/*
	加攻
*/
type AuraScript_2002 struct {
	// 在这里存储每个光环自身的数据
}

func (self *AuraScript_2002) OnStart(aura *BattleAura) {
	fmt.Println("AuraScript_2002 OnStart")
	aura.owner.Prop_addi.Atk += uint32(aura.proto.Param1)
	aura.owner.AddCampaignDetail(CampaignEvent_AuraEffect, int32(ProertyAtk), aura.proto.Param1, 0, 0)
}

func (self *AuraScript_2002) OnUpdate(aura *BattleAura) {
	fmt.Println("AuraScript_2002 OnUpdate")
}

func (self *AuraScript_2002) OnFinish(aura *BattleAura) {
	fmt.Println("AuraScript_2002 OnFinish")
	aura.owner.Prop_addi.Atk -= uint32(aura.proto.Param1)
	aura.owner.AddCampaignDetail(CampaignEvent_AuraEffect, int32(ProertyAtk), -aura.proto.Param1, 0, 0)
}

func (self *AuraScript_2002) OnEvent(evt BattleEventType, aura *BattleAura, ctx *SkillContext) {
	fmt.Println("AuraScript_2002 Event:", evt)
}

func NewAuraScript_2002() AuraScript {
	return &AuraScript_2002{}
}

// ============================================================
/*
	减伤
*/
type AuraScript_2003 struct {
	// 在这里存储每个光环自身的数据
	times int32
}

func (self *AuraScript_2003) OnStart(aura *BattleAura) {
	fmt.Println("AuraScript_2003 OnStart")
}

func (self *AuraScript_2003) OnUpdate(aura *BattleAura) {
	fmt.Println("AuraScript_2003 OnUpdate")
}

func (self *AuraScript_2003) OnFinish(aura *BattleAura) {
	fmt.Println("AuraScript_2003 OnFinish")
}

func (self *AuraScript_2003) OnEvent(evt BattleEventType, aura *BattleAura, ctx *SkillContext) {
	fmt.Println("AuraScript_2003 Event:", evt)
	if evt == BattleEvent_AftDef {
		if self.times >= aura.proto.Param1 {
			return
		}
		self.times++
		if ctx.damage.hurt > uint32(aura.proto.Param2) {
			ctx.damage.hurt -= uint32(aura.proto.Param2)
		} else {
			ctx.damage.hurt = 0
		}
		if self.times >= aura.proto.Param1 {
			aura.finish = true
			return
		}
	}
	aura.owner.AddCampaignDetail(CampaignEvent_AuraEffect, int32(ProertyHurtDec), aura.proto.Param2, 0, 0)
}

func NewAuraScript_2003() AuraScript {
	return &AuraScript_2003{}
}
