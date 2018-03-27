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
	_creators = make(map[uint32]script_creator, 0x40)

	// 一行一行往下累积
	register_aura_script(1, NewAuraScript_1) // 回春
	register_aura_script(2, NewAuraScript_2) // 狂燥
	register_aura_script(3, NewAuraScript_3) // 合欢
	register_aura_script(4, NewAuraScript_4) // 吸血
	register_aura_script(5, NewAuraScript_5) // 掉防

}

func create_aura_script(proto *config.AuraProto) AuraScript {
	if f, ok := _creators[proto.ScriptID]; ok {
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
type AuraScript_1 struct {
	// 在这里存储每个光环自身的数据
}

func (self *AuraScript_1) OnStart(aura *BattleAura) {
	fmt.Println("AuraScript_1 OnStart")
}

func (self *AuraScript_1) OnUpdate(aura *BattleAura) {
	fmt.Println("AuraScript_1 OnUpdate")
	aura.owner.Hp += uint32(aura.proto.Param1)
	aura.owner.AddCampaignDetail(CampaignEvent_AuraEffect, int32(ProertyHP), aura.proto.Param1, 0, 0)
}

func (self *AuraScript_1) OnFinish(aura *BattleAura) {
	fmt.Println("AuraScript_1 OnFinish")
}

func (self *AuraScript_1) OnEvent(evt BattleEventType, aura *BattleAura, ctx *SkillContext) {
	fmt.Println("AuraScript_1 Event:", evt)
}

func NewAuraScript_1() AuraScript {
	return &AuraScript_1{}
}

// ============================================================
/*
	加攻
*/
type AuraScript_2 struct {
	// 在这里存储每个光环自身的数据
}

func (self *AuraScript_2) OnStart(aura *BattleAura) {
	fmt.Println("AuraScript_2 OnStart")
	aura.owner.Prop_addi.Atk += uint32(aura.proto.Param1)
	aura.owner.AddCampaignDetail(CampaignEvent_AuraEffect, int32(ProertyAtk), aura.proto.Param1, 0, 0)
}

func (self *AuraScript_2) OnUpdate(aura *BattleAura) {
	fmt.Println("AuraScript_2 OnUpdate")
}

func (self *AuraScript_2) OnFinish(aura *BattleAura) {
	fmt.Println("AuraScript_2 OnFinish")
	aura.owner.Prop_addi.Atk -= uint32(aura.proto.Param1)
	aura.owner.AddCampaignDetail(CampaignEvent_AuraEffect, int32(ProertyAtk), -aura.proto.Param1, 0, 0)
}

func (self *AuraScript_2) OnEvent(evt BattleEventType, aura *BattleAura, ctx *SkillContext) {
	fmt.Println("AuraScript_2 Event:", evt)
}

func NewAuraScript_2() AuraScript {
	return &AuraScript_2{}
}

// ============================================================
/*
	减伤
*/
type AuraScript_3 struct {
	// 在这里存储每个光环自身的数据
	times int32
}

func (self *AuraScript_3) OnStart(aura *BattleAura) {
	fmt.Println("AuraScript_3 OnStart")
}

func (self *AuraScript_3) OnUpdate(aura *BattleAura) {
	fmt.Println("AuraScript_3 OnUpdate")
}

func (self *AuraScript_3) OnFinish(aura *BattleAura) {
	fmt.Println("AuraScript_3 OnFinish")
}

func (self *AuraScript_3) OnEvent(evt BattleEventType, aura *BattleAura, ctx *SkillContext) {
	fmt.Println("AuraScript_3 Event:", evt)
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

func NewAuraScript_3() AuraScript {
	return &AuraScript_3{}
}

// ============================================================
/*
   吸血光环
*/
type AuraScript_4 struct {
	// 在这里存储每个光环自身的数据
}

func (self *AuraScript_4) OnStart(aura *BattleAura) {
	fmt.Println("AuraScript_4 OnStart")
}

func (self *AuraScript_4) OnUpdate(aura *BattleAura) {
	fmt.Println("AuraScript_4 OnUpdate")
}

func (self *AuraScript_4) OnFinish(aura *BattleAura) {
	fmt.Println("AuraScript_4 OnFinish")
}

func (self *AuraScript_4) OnEvent(evt BattleEventType, aura *BattleAura, ctx *SkillContext) {
	fmt.Println("AuraScript_4 Event:", evt)
	if evt == BattleEvent_Damage {
		aura.owner.Hp += uint32(aura.proto.Param1)
		if aura.owner.Hp > aura.owner.Prop.Hp {
			aura.owner.Hp = aura.owner.Prop.Hp
		}
	}
	aura.owner.AddCampaignDetail(CampaignEvent_AuraEffect, int32(ProertyHP), aura.proto.Param1, 0, 0)
}

func NewAuraScript_4() AuraScript {
	return &AuraScript_4{}
}

// ============================================================
/*
   掉防
*/
type AuraScript_5 struct {
	// 在这里存储每个光环自身的数据
}

func (self *AuraScript_5) OnStart(aura *BattleAura) {
	fmt.Println("AuraScript_5 OnStart")
	aura.owner.Prop_addi.Def -= uint32(aura.proto.Param1)
	aura.owner.AddCampaignDetail(CampaignEvent_AuraEffect, int32(ProertyDef), -aura.proto.Param1, 0, 0)
}

func (self *AuraScript_5) OnUpdate(aura *BattleAura) {
	fmt.Println("AuraScript_5 OnUpdate")
}

func (self *AuraScript_5) OnFinish(aura *BattleAura) {
	fmt.Println("AuraScript_5 OnFinish")
	aura.owner.Prop_addi.Def += uint32(aura.proto.Param1)
	aura.owner.AddCampaignDetail(CampaignEvent_AuraEffect, int32(ProertyDef), +aura.proto.Param1, 0, 0)
}

func (self *AuraScript_5) OnEvent(evt BattleEventType, aura *BattleAura, ctx *SkillContext) {
	fmt.Println("AuraScript_5 Event:", evt)
}

func NewAuraScript_5() AuraScript {
	return &AuraScript_5{}
}
