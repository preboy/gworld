package battle

import (
	"core/log"
	"fmt"
	"server/config"
)

// ============================================================================

type AuraEffectType uint32

const (
	_               AuraEffectType = 0 + iota // 光环效果类型
	AET_PropChanged                           // 属性变化： arg1:属性类型  arg2: 属性变化量
)

type AuraScript interface {
	OnStart(aura *BattleAura)
	OnEvent(aura *BattleAura, evt BattleCalcEvent, ctx *SkillContext)
	OnUpdate(aura *BattleAura)
	OnFinish(aura *BattleAura)
}

var _creators map[uint32]func() AuraScript

func init() {
	// 脚本ID  创建函数
	_creators = map[uint32]func() AuraScript{
		1: NewAuraScript_1, // 回春
		2: NewAuraScript_2, // 狂燥
		3: NewAuraScript_3, // 合欢
		4: NewAuraScript_4, // 吸血
		5: NewAuraScript_5, // 掉防
	}
}

func create_script_object(proto *config.Aura) AuraScript {
	if f, ok := _creators[proto.ScriptId]; ok {
		return f()
	} else {
		log.Error("Not Found AuraScript: AuraProto.Id, AuraProto.ScriptId = %v", proto.Id, proto.ScriptId)
		return nil
	}
}

// ============================================================================
/*
	回血
*/
type AuraScript_1 struct {
	// 在这里存储每个光环自身的数据
}

func NewAuraScript_1() AuraScript {
	return &AuraScript_1{}
}

func (self *AuraScript_1) OnStart(aura *BattleAura) {
	fmt.Println("AuraScript_1 OnStart")
}

func (self *AuraScript_1) OnUpdate(aura *BattleAura) {
	fmt.Println("AuraScript_1 OnUpdate")
	aura.owner.Hp += int(aura.proto.Param1)
	aura.owner.GetBattle().BattlePlayEvent_Effect(
		aura.owner, aura.caster, AET_PropChanged, int32(PropType_HP), int32(aura.proto.Param1), 0, 0)
}

func (self *AuraScript_1) OnEvent(aura *BattleAura, evt BattleCalcEvent, ctx *SkillContext) {
	fmt.Println("AuraScript_1 Event:", evt)
}

func (self *AuraScript_1) OnFinish(aura *BattleAura) {
	fmt.Println("AuraScript_1 OnFinish")
}

// ============================================================================
/*
	加攻
*/
type AuraScript_2 struct {
	// 在这里存储每个光环自身的数据
}

func NewAuraScript_2() AuraScript {
	return &AuraScript_2{}
}

func (self *AuraScript_2) OnStart(aura *BattleAura) {
	fmt.Println("AuraScript_2 OnStart")
	aura.owner.Prop_addi.Atk += float64(aura.proto.Param1)
	aura.owner.GetBattle().BattlePlayEvent_Effect(
		aura.owner, aura.caster, AET_PropChanged, int32(PropType_Atk), int32(aura.proto.Param1), 0, 0)
}

func (self *AuraScript_2) OnUpdate(aura *BattleAura) {
	fmt.Println("AuraScript_2 OnUpdate")
}

func (self *AuraScript_2) OnEvent(aura *BattleAura, evt BattleCalcEvent, ctx *SkillContext) {
	fmt.Println("AuraScript_2 Event:", evt)
}

func (self *AuraScript_2) OnFinish(aura *BattleAura) {
	fmt.Println("AuraScript_2 OnFinish")
	aura.owner.Prop_addi.Atk -= float64(aura.proto.Param1)
	aura.owner.GetBattle().BattlePlayEvent_Effect(
		aura.owner, aura.caster, AET_PropChanged, int32(PropType_Atk), -int32(aura.proto.Param1), 0, 0)
}

// ============================================================================
/*
	减伤
*/
type AuraScript_3 struct {
	// 在这里存储每个光环自身的数据
	times int32
}

func NewAuraScript_3() AuraScript {
	return &AuraScript_3{}
}

func (self *AuraScript_3) OnStart(aura *BattleAura) {
	fmt.Println("AuraScript_3 OnStart")
}

func (self *AuraScript_3) OnUpdate(aura *BattleAura) {
	fmt.Println("AuraScript_3 OnUpdate")
}

func (self *AuraScript_3) OnEvent(aura *BattleAura, evt BattleCalcEvent, ctx *SkillContext) {
	fmt.Println("AuraScript_3 Event:", evt)
	if evt == BCE_AftDef {
		if self.times >= aura.proto.Param1 {
			return
		}
		self.times++
		if ctx.damage.hurt > float64(aura.proto.Param2) {
			ctx.damage.hurt -= float64(aura.proto.Param2)
		} else {
			ctx.damage.hurt = 0
		}
		aura.owner.GetBattle().BattlePlayEvent_Effect(
			aura.owner, aura.caster, AET_PropChanged, int32(PropType_Hurt), -int32(aura.proto.Param1), 0, 0)
		if self.times >= aura.proto.Param1 {
			aura.finish = true
			return
		}
	}
}

func (self *AuraScript_3) OnFinish(aura *BattleAura) {
	fmt.Println("AuraScript_3 OnFinish")
}

// ============================================================================
/*
   吸血光环
*/
type AuraScript_4 struct {
	// 在这里存储每个光环自身的数据
}

func NewAuraScript_4() AuraScript {
	return &AuraScript_4{}
}

func (self *AuraScript_4) OnStart(aura *BattleAura) {
	fmt.Println("AuraScript_4 OnStart")
}

func (self *AuraScript_4) OnUpdate(aura *BattleAura) {
	fmt.Println("AuraScript_4 OnUpdate")
}

func (self *AuraScript_4) OnEvent(aura *BattleAura, evt BattleCalcEvent, ctx *SkillContext) {
	fmt.Println("AuraScript_4 Event:", aura.proto.Id, aura.proto.Level, evt)
	if evt == BCE_Damage {
		aura.owner.Hp += int(aura.proto.Param1)
		if aura.owner.Hp > int(aura.owner.Prop.Hp) {
			aura.owner.Hp = int(aura.owner.Prop.Hp)
		}
		aura.owner.GetBattle().BattlePlayEvent_Effect(
			aura.owner, aura.caster, AET_PropChanged, int32(PropType_HP), int32(aura.proto.Param1), 0, 0)
	}
}

func (self *AuraScript_4) OnFinish(aura *BattleAura) {
	fmt.Println("AuraScript_4 OnFinish")
}

// ============================================================================
/*
   掉防
*/
type AuraScript_5 struct {
	// 在这里存储每个光环自身的数据
}

func NewAuraScript_5() AuraScript {
	return &AuraScript_5{}
}

func (self *AuraScript_5) OnStart(aura *BattleAura) {
	fmt.Println("AuraScript_5 OnStart")
	aura.owner.Prop_addi.Def -= float64(aura.proto.Param1)

	aura.owner.GetBattle().BattlePlayEvent_Effect(
		aura.owner, aura.caster, AET_PropChanged, int32(PropType_Def), -int32(aura.proto.Param1), 0, 0)
}

func (self *AuraScript_5) OnUpdate(aura *BattleAura) {
	fmt.Println("AuraScript_5 OnUpdate")
}

func (self *AuraScript_5) OnEvent(aura *BattleAura, evt BattleCalcEvent, ctx *SkillContext) {
	fmt.Println("AuraScript_5 Event:", aura.proto.Id, aura.proto.Level, evt)
}

func (self *AuraScript_5) OnFinish(aura *BattleAura) {
	fmt.Println("AuraScript_5 OnFinish")
	aura.owner.Prop_addi.Def += float64(aura.proto.Param1)

	aura.owner.GetBattle().BattlePlayEvent_Effect(
		aura.owner, aura.caster, AET_PropChanged, int32(PropType_Def), int32(aura.proto.Param1), 0, 0)
}
