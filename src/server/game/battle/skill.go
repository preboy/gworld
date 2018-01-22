package battle

import (
	"core/log"
	"core/math"
	"fmt"
	"server/game/config"
)

// ==================================================

type SkillBattle struct {
	sp          *config.SkillProto
	owner       *BattleUnit //技能拥有者
	cd_time     uint32      // 用于计算CD
	start_time  uint32      // 技能开始释放时间
	update_time uint32      // 对于有update的技能，记录上次时间
	finish      bool        // 是否完成
}

func NewSkillBattle(id, lv uint32) *SkillBattle {
	sp := config.GetSkillProtoConf().GetSkillProto(id, lv)
	if sp == nil {
		return nil
	}
	sb := &SkillBattle{
		sp: sp,
	}
	return sb
}

func (self *SkillBattle) Cast(u *BattleUnit, time uint32) {
	self.owner = u
	self.finish = false
	self.cd_time = 0
	self.start_time = time
	self.update_time = time
	fmt.Println(u.Name(), "释放了技能", self.sp.Id)
	self.onStart()
}

func (self *SkillBattle) Update(time uint32) {
	if self.sp.Itv_t != 0 {
		if time-self.update_time > self.sp.Itv_t {
			self.onUpdate()
			self.update_time = time
		}
	}

	if !self.finish && time-self.start_time >= self.sp.Last_t {
		self.onFinish()
		self.owner = nil
		self.finish = true
		self.cd_time = time
	}
}

// CD时间从技能释放结束开始计算
// 普通技能的CD时间应配置为0
func (self *SkillBattle) IsFree(time uint32) bool {
	if time-self.cd_time >= self.sp.Cd_t {
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
	//释放一次技能: 攻击、加光环
	targets := self.find_targets()
	for _, target := range targets {
		if self.sp.Type == 1 {
			// 攻击
			// 安全性检测
			if target == self.owner {
				fmt.Println("[WARNING]", self.owner.Name(), "要对自己造成伤害", self.sp.Id)
			}
			self.do_attack(target)
		} else if self.sp.Type == 2 {
			// 加光环
			for _, a := range self.sp.Auras {
				target.AddAura(self.owner, a.Id, a.Lv)
			}
		}
	}
}

func (self *SkillBattle) onFinish() {
	if self.sp.Itv_t == 0 {
		self.onUpdate()
	}
}

// private method
func (self *SkillBattle) find_targets() (targets []*BattleUnit) {
	switch self.sp.Target {
	case 0: // 0：自己
		targets = append(targets, self.owner)
	case 1: // 1: 己方全体
		targets = append(targets, self.owner.GetAllies(true)...)
	case 2: // 2: 敌人
		targets = append(targets, self.owner.GetRivals(false)...)
	case 3: // 3: 敌方全体
		targets = append(targets, self.owner.GetRivals(true)...)
	default:
		log.Warning("Invalid Skill Target", self.sp.Target)
	}
	return
}

func (self *SkillBattle) do_attack(target *BattleUnit) {
	sc := &SkillContext{}
	sc.caster = self.owner
	sc.target = target

	sc.caster_prop = sc.caster.Prop
	sc.target_prop = sc.target.Prop

	// step 1: pre calc attack
	for _, aura := range self.owner.Auras {
		if aura != nil {
			aura.OnEvent(BattleEvent_PreAtk, sc)
		}
	}
	// step 2: calc attack
	hurt := sc.caster_prop.Atk + sc.prop_add.Atk
	crit := sc.caster_prop.Crit + sc.prop_add.Crit
	sc.damage_send.hurt = hurt
	sc.damage_send.crit = false
	if math.RandomHitn(int(crit), 100) {
		sc.damage_send.crit = true
		sc.damage_send.hurt = hurt * (sc.caster_prop.Crit_hurt + sc.prop_add.Crit_hurt)
	}
	// step 3: send damage
	for _, aura := range self.owner.Auras {
		if aura != nil {
			aura.OnEvent(BattleEvent_Damage, sc)
		}
	}
	// step 4 : recv damage
	for _, aura := range target.Auras {
		if aura != nil {
			aura.OnEvent(BattleEvent_Damage, sc)
		}
	}

	// step 5: 计算防御
	hurt = sc.damage_send.hurt - sc.target_prop.Def
	if hurt < 0 {
		hurt = 1
	}
	sc.damage_recv.hurt = hurt
	// step 6: 防御之后光环减免
	for _, aura := range target.Auras {
		if aura != nil {
			aura.OnEvent(BattleEvent_AftDef, sc)
		}
	}
	// step 7 : 计算实际伤害
	sc.damage.hurt = sc.damage_recv.hurt - sc.damage_sub.hurt
	fmt.Println(sc.caster.Name(), " 对 ", sc.target.Name(), "造成了伤害:", sc.damage.hurt)
	if sc.damage.hurt < target.Prop.Hp_cur {
		target.Prop.Hp_cur -= sc.damage.hurt
	} else {
		target.Prop.Hp_cur = 0
		target.Dead = true
		fmt.Println(sc.target.Name(), "战死")
	}

}
