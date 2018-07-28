package battle

import (
	// "core/log"
	"core/math"
	"fmt"
	"server/game/config"
)

// ==================================================

type SkillDamage struct {
	hurt float64
	crit bool
}

type SkillContext struct {
	caster      *BattleUnit
	target      *BattleUnit
	caster_prop *Property   // 攻击者的属性(此处只读不写)
	target_prop *Property   // 防御者的属性(此处只读不写)
	prop_add    Property    // 攻击者光环加成
	damage_send SkillDamage // 攻击者造成实际伤害
	damage_recv SkillDamage // 防御者计算防御之后的伤害
	damage_sub  SkillDamage // 防御者计算防御之后光环减免部分
	damage      SkillDamage //最终造成的实际伤害
}

type BattleSkill struct {
	proto        *config.SkillProto // 技能原型
	caster       *BattleUnit        // 技能拥有者
	time         uint32             // 当前时间
	cd_time      uint32             // 用于计算CD(技能结束之后开始计算)
	start_time   uint32             // 技能开始释放时间
	update_time  uint32             // 对于有update的技能，记录上次时间
	finish       bool               // 是否完成
	target_major []*BattleUnit      // 第一目标
	target_minor []*BattleUnit      // 第二目标
}

func NewBattleSkill(id, lv uint32) *BattleSkill {
	proto := config.GetSkillProto(id, lv)
	if proto == nil {
		return nil
	}
	sb := &BattleSkill{
		proto:  proto,
		finish: true,
	}
	return sb
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

	self.caster.GetBattle().BattlePlayEvent_Cast(self.caster, self.proto.Id, self.proto.Level, targets)
	// fmt.Println("[技能", time, "]", "释放了技能:", self.proto.Name)
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
				if math.RandomHitn(int(aura.Prob), 100) {
					target.AddAura(self.caster, aura.Id, aura.Lv)
				}
			}
		}
	}
	for _, target := range self.target_minor {
		if target != nil && !target.Dead() {
			for _, aura := range self.proto.Aura_minor {
				if math.RandomHitn(int(aura.Prob), 100) {
					target.AddAura(self.caster, aura.Id, aura.Lv)
				}
			}
		}
	}

	// fmt.Println("[技能", self.time, "]", self.caster.Name(), "释放的技能结束了", self.proto.Name)
}

// 处理技能的附加属性
func attachment_skill_attr(p *Property, props []*config.PropConf) {
	for _, prop := range props {
		switch PropType(prop.Id) {
		case PropType_HP:
			{
				p.Hp += prop.Val
			}
		case PropType_Apm:
			{
				p.Apm += prop.Val
			}
		case PropType_Atk:
			{
				p.Atk += prop.Val
			}
		case PropType_Def:
			{
				p.Def += prop.Val
			}
		case PropType_Crit:
			{
				p.Crit += prop.Val
			}
		case PropType_Hurt:
			{
				p.Hurt += prop.Val
			}
		default:
			{
				fmt.Println("Unsupported Proptype:", prop.Id)
			}
		}
	}
}

func (self *BattleSkill) do_attack(target *BattleUnit, major bool) {
	ctx := &SkillContext{}

	ctx.caster = self.caster
	ctx.target = target

	ctx.caster_prop = ctx.caster.Prop
	ctx.target_prop = ctx.target.Prop

	// ----------- 攻击方 -----------
	props := self.proto.Prop_major
	if !major {
		props = self.proto.Prop_minor
	}

	// step 1: 技能的附加属性
	attachment_skill_attr(&ctx.prop_add, props)

	// step 2: 计算光环
	for _, aura := range ctx.caster.Auras_battle {
		if aura != nil {
			aura.OnEvent(BCE_PreAtk, ctx)
		}
	}

	// step 3: 计算输出伤害
	hurt := ctx.caster_prop.Atk + ctx.prop_add.Atk
	crit := ctx.caster_prop.Crit + ctx.prop_add.Crit

	ctx.damage_send.hurt = hurt
	ctx.damage_send.crit = false

	if math.RandomHitn(int(crit), 100) {
		ctx.damage_send.crit = true
		ctx.damage_send.hurt = hurt * (1 + (ctx.caster_prop.Hurt+ctx.prop_add.Hurt)/100.0)
	}

	// step 4: 处理攻击方光环事件
	for _, aura := range ctx.caster.Auras_battle {
		if aura != nil {
			aura.OnEvent(BCE_Damage, ctx)
		}
	}

	// ----------- 防御方 -----------

	// step 5: 计算防御
	if ctx.damage_send.hurt >= ctx.target_prop.Def {
		ctx.damage_recv.hurt = ctx.damage_send.hurt - ctx.target_prop.Def
	} else {
		ctx.damage_recv.hurt = 0
	}
	ctx.damage_recv.crit = ctx.damage_send.crit

	// step 6: 计算光环
	for _, aura := range target.Auras_battle {
		if aura != nil {
			aura.OnEvent(BCE_AftDef, ctx)
		}
	}

	// step 7: 计算最终伤害
	if ctx.damage_recv.hurt > ctx.damage_sub.hurt {
		ctx.damage.hurt = ctx.damage_recv.hurt - ctx.damage_sub.hurt
	} else {
		ctx.damage.hurt = 0
	}
	ctx.damage.crit = ctx.damage_recv.crit

	text := " <伤害了> "
	if target.Hp > int(ctx.damage.hurt) {
		target.Hp -= int(ctx.damage.hurt)
	} else {
		target.Hp = 0
		text = " <击杀了> "
	}

	var is_crit uint32
	if ctx.damage.crit {
		is_crit = 1
		text += "[+暴击] "
	}

	str := ctx.caster.Name() + "[" + ctx.caster.Skill_curr.proto.Name + "]" + text + target.Name()

	fmt.Println(self.time, str, ctx.damage.hurt, target.Hp, "/", target.Prop.Hp)

	self.caster.GetBattle().BattlePlayEvent_Hurt(self.caster, target, uint32(ctx.damage.hurt), is_crit, 0)
}
