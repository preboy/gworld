package battle

import (
	// "core/log"
	"core/rand"
	"fmt"
	"game/config"
)

const (
	_           BattleCalcEvent = 0 + iota
	BCE_PreSend                 // 发送最终伤害前(攻方, damage_send) 如:百分百暴击
	BCE_PreDef                  // 防御之前(防方, damage_recv) 如:克制暴击、攻击类型克制
	BCE_PreHurt                 // 最终伤害之前(防方, damage_calc)如:抵挡伤害
	BCE_PreBack                 // 生成伤害之后(攻方, kill)如:吸血，死亡触发事件
)

// ============================================================================

type SkillContext struct {
	caster      *BattleUnit
	target      *BattleUnit
	caster_prop *PropertyGroup
	target_prop *PropertyGroup

	kill        bool    // 是否最后一击
	crit        bool    // 是否暴击
	hurt        float64 // 未计算暴击前的伤害
	damage_send float64 // 攻击者造成实际伤害
	damage_recv float64 // 防御者计算格挡之后、计算防御之前的伤害
	damage_calc float64 // 最终造成的实际伤害
}

type BattleSkill struct {
	proto        *config.Skill // 技能原型
	caster       *BattleUnit   // 技能拥有者
	time         uint32        // 当前时间
	cd_time      uint32        // 用于计算CD(技能结束之后开始计算)
	start_time   uint32        // 技能开始释放时间
	update_time  uint32        // 对于有update的技能，记录上次时间
	finish       bool          // 是否完成
	target_major []*BattleUnit // 第一目标
	target_minor []*BattleUnit // 第二目标
}

func NewBattleSkill(id, lv uint32) *BattleSkill {
	proto := config.SkillProtoConf.Query(id, lv)
	if proto != nil {
		return &BattleSkill{
			proto:  proto,
			finish: true,
		}
	}
	return nil
}

func (self *BattleSkill) Cast(caster *BattleUnit, time uint32) {
	self.finish = false
	self.caster = caster
	self.start_time = time
	self.update_time = time + self.proto.Prepare_t
	self.find_target()
	self.onStart()

	targets := []uint32{}
	for _, t := range self.target_major {
		targets = append(targets, t.Id)
	}

	self.caster.GetBattle().BattlePlayEvent_Cast(self.caster, self.proto.Id, self.proto.Lv, targets)

	fmt.Printf("[%d] %s<%d> 释放了 %s\n", time, self.caster.Base.Name(), self.caster.Pos, self.proto.Name)
}

func (self *BattleSkill) Update(time uint32) {
	if self.finish {
		return
	}

	if self.proto.Prepare_t != 0 && time-self.start_time < self.proto.Prepare_t {
		return
	}

	self.time = time
	if self.proto.Itv_t != 0 {
		if time-self.update_time >= self.proto.Itv_t {
			self.update_time = time
			self.onUpdate()
		}
	}
	if time-self.start_time >= self.proto.Prepare_t+self.proto.Effect_t {
		self.onFinish()
		self.caster = nil
		self.finish = true
		self.cd_time = time
	}
}

// CD时间从技能释放结束开始计算
func (self *BattleSkill) IsFree(time uint32) bool {
	return time-self.cd_time >= self.proto.Cd_t
}

func (self *BattleSkill) IsFinish() bool {
	return self.finish
}

func (self *BattleSkill) onStart() {
	// nothing to do
}

func (self *BattleSkill) onUpdate() {
	// 对敌人造成伤害
	if self.proto.Type == 1 {
		for _, target := range self.target_major {
			if target != nil && !target.Dead() {
				self.do_attack(target, true)
			}
		}
		for _, target := range self.target_minor {
			if target != nil && !target.Dead() {
				self.do_attack(target, false)
			}
		}
	} else if self.proto.Type == 2 {
		// TODO
	}
}

func (self *BattleSkill) onFinish() {
	if self.proto.Itv_t == 0 {
		self.onUpdate()
	}

	// take effect for aura
	for _, target := range self.target_major {
		if target != nil && !target.Dead() {
			for _, aura := range self.proto.Aura_major {
				if rand.RandomHitn(int(aura.Prob), 100) {
					target.AddAura(self.caster, aura.Id, aura.Lv)
				}
			}
		}
	}
	for _, target := range self.target_minor {
		if target != nil && !target.Dead() {
			for _, aura := range self.proto.Aura_minor {
				if rand.RandomHitn(int(aura.Prob), 100) {
					target.AddAura(self.caster, aura.Id, aura.Lv)
				}
			}
		}
	}

	fmt.Printf("[%d %s<%d>的技能 <%s> 结束了\n", self.time, self.caster.Base.Name(), self.caster.Pos, self.proto.Name)
}

func (self *BattleSkill) do_attack(target *BattleUnit, major bool) {
	ctx := &SkillContext{}

	ctx.caster = self.caster
	ctx.target = target

	ctx.caster_prop = ctx.caster.Prop
	ctx.target_prop = ctx.target.Prop

	// ------------------------------------------------------------------------
	// 攻击方

	// 计算伤害值
	if major {
		ctx.hurt = self.get_attack_for_target_major()
	} else {
		ctx.hurt = self.get_attack_for_target_minor()
	}

	// 计算暴击
	if rand.RandomHitn(int(ctx.caster_prop.Value(PropType_Crit)), 100) {
		ctx.crit = true
		ctx.damage_send = ctx.hurt * (1 + ctx.caster_prop.Value(PropType_Hurt)/100)
	} else {
		ctx.damage_send = ctx.hurt
	}

	// 计算光环(对damage_send做随后的调整，比如必定暴击)
	for _, aura := range ctx.caster.Auras_battle {
		if aura != nil {
			aura.OnEvent(BCE_PreSend, ctx)
		}
	}

	// ------------------------------------------------------------------------
	// 防御方

	ctx.damage_recv = ctx.damage_send

	// 计算光环(对攻击先行调整，比如格挡暴击、处理攻防类型克制关系)
	for _, aura := range target.Auras_battle {
		if aura != nil {
			aura.OnEvent(BCE_PreDef, ctx)
		}
	}

	ctx.damage_calc = ctx.damage_recv - ctx.target_prop.Value(PropType_Def)
	if ctx.damage_calc < 1 {
		ctx.damage_calc = 1
	}

	// 计算光环(对抵挡伤害类的光环在此工作)
	for _, aura := range target.Auras_battle {
		if aura != nil {
			aura.OnEvent(BCE_PreHurt, ctx)
		}
	}

	origin_hp := ctx.target.Hp

	// 实际伤害
	ctx.target.SubHp(ctx.damage_calc)

	text := " <伤害了>"
	if target.Hp <= 0 {
		text = " <击杀了>"
		ctx.kill = true
	}

	is_crit := uint32(0)
	if ctx.crit {
		is_crit = 1
		text += "[+暴击]"
	}

	// 计算光环
	for _, aura := range ctx.caster.Auras_battle {
		if aura != nil {
			aura.OnEvent(BCE_PreBack, ctx)
		}
	}

	fmt.Printf("[%d] %s %s %s, 伤害=%f/%f, %f/%f\n", self.time, ctx.caster.Name(), text, ctx.target.Name(), ctx.damage_calc, origin_hp, ctx.target.Hp, ctx.target_prop.Value(PropType_HP))

	self.caster.GetBattle().BattlePlayEvent_Hurt(self.caster, target, uint32(ctx.damage_calc), is_crit, 0)
}
