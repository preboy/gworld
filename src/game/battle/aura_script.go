package battle

import (
	"core/log"
	"fmt"
	"game/config"
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
		1: NewAuraScript_1, // 每次update恢复HP
		2: NewAuraScript_2, // 开始、结束时分别增加、减少属性
		3: NewAuraScript_3, // 光环时间段时，抵挡x次伤害
		4: NewAuraScript_4, // 给敌人造成伤害时，吸血
		5: NewAuraScript_4, // 加光环
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
	1 每次update恢复HP
*/

type AuraScript_1 struct {
}

func NewAuraScript_1() AuraScript {
	return &AuraScript_1{}
}

func (self *AuraScript_1) OnStart(aura *BattleAura) {
}

func (self *AuraScript_1) OnUpdate(aura *BattleAura) {
	val := aura.owner.AddHp(float64(aura.proto.Params[1]))

	aura.owner.GetBattle().BattlePlayEvent_Effect(
		aura.owner, aura.caster, AET_PropChanged, int32(PropType_HP), int32(val), 0, 0)
}

func (self *AuraScript_1) OnEvent(aura *BattleAura, evt BattleCalcEvent, ctx *SkillContext) {
}

func (self *AuraScript_1) OnFinish(aura *BattleAura) {
}

// ============================================================================
/*
	2 开始、结束时分别增加、减少属性
*/

type AuraScript_2 struct {
}

func NewAuraScript_2() AuraScript {
	return &AuraScript_2{}
}

func (self *AuraScript_2) OnStart(aura *BattleAura) {
	aura.owner.Prop.AddProps(aura.proto.Props)

	for _, v := range aura.proto.Props {
		aura.owner.GetBattle().BattlePlayEvent_Effect(
			aura.owner, aura.caster, AET_PropChanged, int32(v.Id), int32(v.Val), 0, 0)
	}
}

func (self *AuraScript_2) OnUpdate(aura *BattleAura) {
}

func (self *AuraScript_2) OnEvent(aura *BattleAura, evt BattleCalcEvent, ctx *SkillContext) {
}

func (self *AuraScript_2) OnFinish(aura *BattleAura) {
	aura.owner.Prop.SubProps(aura.proto.Props)
	for _, v := range aura.proto.Props {
		aura.owner.GetBattle().BattlePlayEvent_Effect(
			aura.owner, aura.caster, AET_PropChanged, int32(v.Id), -int32(v.Val), 0, 0)
	}
}

// ============================================================================
/*
	3 光环时间段时，抵挡x次伤害
*/

type AuraScript_3 struct {
	curr_times int32
	times      int32
	hurt       float64
}

func NewAuraScript_3() AuraScript {
	return &AuraScript_3{}
}

func (self *AuraScript_3) OnStart(aura *BattleAura) {
	self.times = aura.proto.Params[0]
	self.hurt = float64(aura.proto.Params[1])
}

func (self *AuraScript_3) OnUpdate(aura *BattleAura) {
}

func (self *AuraScript_3) OnEvent(aura *BattleAura, evt BattleCalcEvent, ctx *SkillContext) {
	if evt == BCE_PreHurt {
		if self.curr_times >= self.times {
			return
		}
		self.curr_times++
		if ctx.damage_calc > self.hurt {
			ctx.damage_calc -= self.hurt
		} else {
			ctx.damage_calc = 1
		}

		aura.owner.GetBattle().BattlePlayEvent_Effect(
			aura.owner, aura.caster, AET_PropChanged, int32(PropType_Hurt), -int32(self.hurt), 0, 0)

		if self.curr_times >= self.times {
			aura.finish = true
			return
		}
	}
}

func (self *AuraScript_3) OnFinish(aura *BattleAura) {
}

// ============================================================================
/*
   4 给敌人造成伤害时，吸收造成伤害百分比的血
*/

type AuraScript_4 struct {
}

func NewAuraScript_4() AuraScript {
	return &AuraScript_4{}
}

func (self *AuraScript_4) OnStart(aura *BattleAura) {
}

func (self *AuraScript_4) OnUpdate(aura *BattleAura) {
}

func (self *AuraScript_4) OnEvent(aura *BattleAura, evt BattleCalcEvent, ctx *SkillContext) {
	fmt.Println("AuraScript_4 Event:", aura.proto.Id, aura.proto.Lv, evt)

	if evt == BCE_PreBack {
		val := aura.owner.AddHp(float64(aura.proto.Params[0]) * ctx.damage_calc)

		aura.owner.GetBattle().BattlePlayEvent_Effect(
			aura.owner, aura.caster, AET_PropChanged, int32(PropType_HP), int32(val), 0, 0)
	}
}

func (self *AuraScript_4) OnFinish(aura *BattleAura) {
	fmt.Println("AuraScript_4 OnFinish")
}

// ============================================================================
/*
   未完成
   5 加光环，给某些战斗位置的单位加光环，参数还没有想好
   参数：[光环ID，光环LV，位置1，位置2...]
*/

type AuraScript_5 struct {
}

func NewAuraScript_5() AuraScript {
	return &AuraScript_5{}
}

func (self *AuraScript_5) OnStart(aura *BattleAura) {
	// todo

	// aura_id := aura.proto.Params[0]
	// aura_lv := aura.proto.Params[1]
}

func (self *AuraScript_5) OnUpdate(aura *BattleAura) {
}

func (self *AuraScript_5) OnEvent(aura *BattleAura, evt BattleCalcEvent, ctx *SkillContext) {
}

func (self *AuraScript_5) OnFinish(aura *BattleAura) {
	fmt.Println("AuraScript_5 OnFinish")
}
