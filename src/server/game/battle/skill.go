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

type BattleSkill struct {
	proto       *config.SkillProto // 技能原型
	owner       *BattleUnit        //技能拥有者
	time        int32              // 当前时间
	cd_time     int32              // 用于计算CD
	start_time  int32              // 技能开始释放时间
	update_time int32              // 对于有update的技能，记录上次时间
	finish      bool               // 是否完成
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

func (self *BattleSkill) Reset(common bool) {
	self.owner = nil
	self.finish = true
	self.start_time = 0
	self.update_time = 0
	if !common {
		self.cd_time = 0
	} else {
		self.cd_time = -self.proto.Cd_t
	}
}

func (self *BattleSkill) Cast(u *BattleUnit, time int32) {
	self.owner = u
	self.finish = false
	self.start_time = time
	self.update_time = time
	self.onStart()
	self.owner.AddCampaignDetail(CampaignEvent_Cast, int32(self.proto.Id), int32(self.proto.Level), 0, 0)
	// fmt.Println("[技能", time, "]", u.Name(), "释放了技能:", self.proto.Name)
}

func (self *BattleSkill) Update(time int32) {
	self.time = time
	if self.finish {
		return
	}
	if self.proto.Itv_t != 0 {
		if time-self.update_time >= self.proto.Itv_t {
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
func (self *BattleSkill) IsFree(time int32) bool {
	return time-self.cd_time >= self.proto.Cd_t
}

func (self *BattleSkill) IsFinish() bool {
	return self.finish
}

func (self *BattleSkill) onStart() {
	// nothing to do
}

func (self *BattleSkill) onUpdate() {
	target := self.owner.Rival

	switch self.proto.Target {
	case 1: // 自己: 对自己只能加光环
		{
			for _, a := range self.proto.Auras {
				// println("==========================  自己  加光环", config.GetAuraProtoConf().GetAuraProto(a.Id, a.Lv).Name)
				self.owner.AddAura(self.owner, a.Id, a.Lv)
			}
		}
	case 2: // 敌人：攻击、光环
		{
			self.do_attack(target)
			for _, a := range self.proto.Auras {
				// println("========================== 敌人   加光环", config.GetAuraProtoConf().GetAuraProto(a.Id, a.Lv).Name)
				target.AddAura(self.owner, a.Id, a.Lv)
			}
		}
	default: // 未知的对象
		{
			fmt.Println("unknown skill target:", self.proto)
		}
	}
}

func (self *BattleSkill) onFinish() {
	if self.proto.Itv_t == 0 {
		self.onUpdate()
	}
	// fmt.Println("[技能", self.time, "]", self.owner.Name(), "释放的技能结束了", self.proto.Name)
}

// 处理技能的附加属性
func attachment_skill_attr(p *Property, attrs []*config.AttrConf) {
	for _, attr := range attrs {
		switch AttrType(attr.Id) {
		case AttrType_HP:
			{
				p.Hp += attr.Val
			}
		case AttrType_Atk:
			{
				p.Atk += attr.Val
			}
		case AttrType_Def:
			{
				p.Def += attr.Val
			}
		case AttrType_Crit:
			{
				p.Crit += attr.Val
			}
		case AttrType_CritHurt:
			{
				p.CritHurt += attr.Val
			}
		default:
			{
				fmt.Println("Unknown Attrtype", attr.Id)
			}
		}
	}
}

func (self *BattleSkill) do_attack(target *BattleUnit) {
	ctx := &SkillContext{}
	ctx.caster = self.owner
	ctx.target = target

	ctx.caster_prop = ctx.caster.Prop
	ctx.target_prop = ctx.target.Prop

	// ----------- 攻击方 -----------

	// step 1: 技能的附加属性
	attachment_skill_attr(&ctx.prop_add, self.proto.Attrs)

	// step 2: 计算光环
	for _, aura := range ctx.caster.Auras_battle {
		if aura != nil {
			aura.OnEvent(BattleEvent_PreAtk, ctx)
		}
	}

	// step 3: 计算输出伤害
	hurt := ctx.caster_prop.Atk + ctx.prop_add.Atk
	crit := ctx.caster_prop.Crit + ctx.prop_add.Crit

	ctx.damage_send.hurt = hurt
	ctx.damage_send.crit = false

	if math.RandomHitn(int(crit), 100) {
		ctx.damage_send.crit = true
		ctx.damage_send.hurt = hurt * (1 + (ctx.caster_prop.CritHurt+ctx.prop_add.CritHurt)/100.0)
	}

	// step 4: 处理攻击方光环事件
	for _, aura := range ctx.caster.Auras_battle {
		if aura != nil {
			aura.OnEvent(BattleEvent_Damage, ctx)
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
			aura.OnEvent(BattleEvent_AftDef, ctx)
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
	if target.Hp > ctx.damage.hurt {
		target.Hp -= ctx.damage.hurt
	} else {
		target.Hp = 0
		text = " <击杀了> "
	}

	var is_crit int32
	if ctx.damage.crit {
		is_crit = 1
		text += "[+暴击]"
	}

	fmt.Println("[技能", self.time, "]", ctx.caster.Name(), text, ctx.target.Name(), ctx.damage.hurt, ctx.target.Hp, "/", ctx.target.Prop.Hp, "[", ctx.caster.Skill_curr.proto.Name, "]")

	target.AddCampaignDetail(CampaignEvent_Hurt, int32(ctx.damage.hurt), is_crit, 0, 0)
}
