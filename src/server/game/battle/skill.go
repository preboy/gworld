package battle

import (
	// "core/log"
	"core/math"
	"fmt"
	"server/game/config"
)

// ==================================================

type SkillDamage struct {
	hurt uint32
	crit bool
}

type SkillContext struct {
	caster      *BattleUnit
	target      *BattleUnit
	caster_prop *Property   // 攻击者的基本属性(只读)
	target_prop *Property   // 防御者的基本属性(只读)
	prop_add    Property    // 攻击者光环加成
	damage_send SkillDamage // 攻击者造成实际伤害
	damage_recv SkillDamage // 防御者计算防御之后的伤害
	damage_sub  SkillDamage // 防御者计算防御之后光环减免部分
	damage      SkillDamage //最终造成的实际伤害
}

type SkillBattle struct {
	proto       *config.SkillProto // 技能原型
	owner       *BattleUnit        //技能拥有者
	cd_time     int32              // 用于计算CD
	start_time  int32              // 技能开始释放时间
	update_time int32              // 对于有update的技能，记录上次时间
	finish      bool               // 是否完成
}

func NewSkillBattle(id, lv uint32) *SkillBattle {
	proto := config.GetSkillProtoConf().GetSkillProto(id, lv)
	if proto == nil {
		return nil
	}
	sb := &SkillBattle{
		proto:  proto,
		finish: true,
	}
	return sb
}

func (self *SkillBattle) Reset(common bool) {
	self.owner = nil
	self.finish = true
	self.start_time = 0
	self.update_time = 0
	if !common {
		self.cd_time = 0
	} else {
		self.cd_time = -(self.proto.Cd_t)
	}
}

func (self *SkillBattle) Cast(u *BattleUnit, time int32) {
	self.owner = u
	self.finish = false
	self.start_time = time
	self.update_time = time
	self.onStart()
	fmt.Println(u.Name(), "释放了技能", self.proto.Id)
}

func (self *SkillBattle) Update(time int32) {
	if self.finish {
		return
	}
	if self.proto.Itv_t != 0 {
		if time-self.update_time > self.proto.Itv_t {
			self.onUpdate()
			self.update_time = time
		}
	}
	if time-self.start_time >= self.proto.Last_t {
		self.onFinish()
		self.owner = nil
		self.finish = true
		self.cd_time = time
	}
}

// CD时间从技能释放结束开始计算
// 普通技能的CD时间应配置为0
func (self *SkillBattle) IsFree(time int32) bool {
	if time-self.cd_time >= self.proto.Cd_t {
		return true
	}
	return false
}

func (self *SkillBattle) IsFinish() bool {
	return self.finish
}

func (self *SkillBattle) onStart() {
	// nothing to do
}

func (self *SkillBattle) onUpdate() {
	target := self.owner.Rival
	switch self.proto.Type {
	case 1: // 攻击目标
		{
			if target == self.owner {
				fmt.Println("[WARNING]", self.owner.Name(), "要对自己造成伤害", self.proto.Id)
			}
			self.do_attack(target)
		}
	case 2: // 加光环
		{
			for _, a := range self.proto.Auras {
				target.AddAuraBattle(self.owner, a.Id, a.Lv)
			}
		}
	default:
		{
			fmt.Println("unknown skill type", self.proto.Type)
		}
	}
}

func (self *SkillBattle) onFinish() {
	if self.proto.Itv_t == 0 {
		self.onUpdate()
	}
}

func (self *SkillBattle) do_attack(target *BattleUnit) {
	ctx := &SkillContext{}
	ctx.caster = self.owner
	ctx.target = target

	ctx.caster_prop = ctx.caster.Prop
	ctx.target_prop = ctx.target.Prop

	// step 1: 计算光环
	for _, aura := range ctx.caster.Auras_basic {
		if aura != nil {
			aura.OnEvent(BattleEvent_PreAtk, ctx)
		}
	}
	for _, aura := range ctx.caster.Auras_battle {
		if aura != nil {
			aura.OnEvent(BattleEvent_PreAtk, ctx)
		}
	}
	for _, aura := range ctx.caster.Auras_guarder {
		if aura != nil {
			aura.OnEvent(BattleEvent_PreAtk, ctx)
		}
	}

	// step 2: 计算输出伤害
	hurt := ctx.caster_prop.Atk + ctx.prop_add.Atk
	crit := ctx.caster_prop.Crit + ctx.prop_add.Crit
	ctx.damage_send.hurt = hurt
	ctx.damage_send.crit = false
	if math.RandomHitn(int(crit), 100) {
		ctx.damage_send.crit = true
		ctx.damage_send.hurt = hurt * (ctx.caster_prop.Crit_hurt + ctx.prop_add.Crit_hurt)
	}

	// step 3: 计算防御
	hurt = ctx.damage_send.hurt - ctx.target_prop.Def
	if hurt < 0 {
		hurt = 1
	}
	ctx.damage_recv.hurt = hurt

	// step 4: 计算光环
	for _, aura := range target.Auras_basic {
		if aura != nil {
			aura.OnEvent(BattleEvent_AftDef, ctx)
		}
	}
	for _, aura := range target.Auras_battle {
		if aura != nil {
			aura.OnEvent(BattleEvent_AftDef, ctx)
		}
	}
	for _, aura := range target.Auras_guarder {
		if aura != nil {
			aura.OnEvent(BattleEvent_AftDef, ctx)
		}
	}

	// step 5: 计算最终伤害
	ctx.damage.hurt = ctx.damage_recv.hurt - ctx.damage_sub.hurt
	if ctx.damage.hurt < target.Prop.Hp_cur {
		target.Prop.Hp_cur -= ctx.damage.hurt
		fmt.Println(ctx.caster.Name(), " 伤害了 ", ctx.target.Name(), ctx.damage.hurt)
	} else {
		target.Prop.Hp_cur = 0
		target.Dead = true
		fmt.Println(ctx.caster.Name(), " <击杀了> ", ctx.target.Name(), ctx.damage.hurt)
	}

}
