package battle

// ==================================================
type BattleEvent uint32

const (
	_                  BattleEvent = 1 + iota
	BattleEvent_PreAtk             // 计算攻击之前 (累积光环的附加攻击)
	BattleEvent_Damage             // 计算伤害 (双方暂不做任何计算)
	BattleEvent_AftDef             // 计算防御之后 (抵挡伤害)
)

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

// ==================================================

type BattleUnit struct {
	UnitType   uint32         // 生物类型
	troop      *BattleTroop   // 队伍
	Prop       Property       // 战斗属性
	Skills     []*SkillBattle // 主动释放技能
	Auras      []*AuraBattle  // 光环(技能ID)
	Dead       bool           // 是否死亡
	Skill_Curr *SkillBattle   // 是否正在释放技能
}

func (self *BattleUnit) Update(time uint32) {
	if self.Dead {
		return
	}
	// 释放技能
	if self.Skill_Curr == nil {
		// 释放
		for _, v := range self.Skills {
			if v.InCD(time) {
				continue
			}
			self.Skill_Curr = v
			break
		}
		if self.Skill_Curr != nil {
			self.Skill_Curr.Cast(self, time)
		}
	} else {
		self.Skill_Curr.Update(time)
		if self.Skill_Curr.IsFinish() {
			self.Skill_Curr = nil
		}
	}
	// 光环
	for k, aura := range self.Auras {
		if aura != nil {
			aura.Update(time)
			if aura.IsFinish() {
				self.Auras[k] = nil
			}
		}
	}
}

// 寻找对手 all:敌方所有单位
func (self *BattleUnit) GetRivals(all bool) (ret []*BattleUnit) {
	if all {
		return self.troop.GetRivals()
	} else {
		r := self.troop.GetRival(self)
		if r != nil {
			ret = append(ret, r)
		}
		return
	}
}

// 寻找所有的队友 include_myself:是否包括自己
func (self *BattleUnit) GetAllies(include_myself bool) []*BattleUnit {
	if include_myself {
		return self.troop.GetMembers(nil)
	} else {
		return self.troop.GetMembers(self)
	}
}

func (self *BattleUnit) AddAura(caster *BattleUnit, id uint32, lv uint32) {
	aura := NewAuraBattle(id, lv)
	if aura == nil {
		return
	}
	aura.Init(caster, self)
	self.Auras = append(self.Auras, aura)
}

func (self *BattleUnit) DelAura(id, lv uint32) {
	for k, aura := range self.Auras {
		if aura.sp.Id == id && aura.sp.Lv == lv {
			self.Auras[k] = nil
			break
		}
	}
}

// ==================================================

type BattleTroop struct {
	battle *Battle
	top    *BattleUnit
	mid    *BattleUnit
	btm    *BattleUnit
}

func NewBattleTroop(t *BattleUnit, m *BattleUnit, b *BattleUnit) *BattleTroop {
	bt := &BattleTroop{
		top: t,
		mid: m,
		btm: b,
	}
	t.troop = bt
	m.troop = bt
	b.troop = bt
	return bt
}

func (self *BattleTroop) SetTop(u *BattleUnit) {
	self.top = u
}

func (self *BattleTroop) SetMid(u *BattleUnit) {
	self.mid = u
}

func (self *BattleTroop) SetBtm(u *BattleUnit) {
	self.btm = u
}

func (self *BattleTroop) Lose() bool {
	return (self.top == nil || self.top.Dead) && (self.mid == nil || self.mid.Dead) && (self.btm == nil || self.btm.Dead)
}

func (self *BattleTroop) Update(time uint32) {
	if self.top != nil {
		self.top.Update(time)
	}
	if self.mid != nil {
		self.mid.Update(time)
	}
	if self.btm != nil {
		self.btm.Update(time)
	}
}

func (self *BattleTroop) GetRivals() []*BattleUnit {
	troop := self.battle.GetAnotherTroop(self)
	return troop.GetMembers(nil)
}

func (self *BattleTroop) GetRival(u *BattleUnit) *BattleUnit {
	if u == nil {
		return nil
	}
	troop := self.battle.GetAnotherTroop(self)
	if u == self.top {
		return troop.top
	}
	if u == self.mid {
		return troop.mid
	}
	if u == self.btm {
		return troop.btm
	}
	return nil
}

func (self *BattleTroop) GetMembers(exclude *BattleUnit) (ret []*BattleUnit) {
	if exclude == nil || exclude != self.top {
		if self.top != nil {
			ret = append(ret, self.top)
		}
	}
	if exclude == nil || exclude != self.mid {
		if self.mid != nil {
			ret = append(ret, self.mid)
		}
	}
	if exclude == nil || exclude != self.btm {
		if self.btm != nil {
			ret = append(ret, self.btm)
		}
	}
	return
}

// ==================================================

type BattleResult struct {
	Win uint32 // 0:attacker负  1:attacker胜
}

// 生成发给客户端的消息
func (self *BattleResult) ToMsg() string {
	return "{}"
}

// ==================================================

type Battle struct {
	attacker *BattleTroop
	defender *BattleTroop
}

func NewBattle(a *BattleTroop, d *BattleTroop) *Battle {
	b := &Battle{
		attacker: a,
		defender: d,
	}
	a.battle = b
	d.battle = b
	return b
}

// ==================================================

func (self *Battle) GetAnotherTroop(troop *BattleTroop) *BattleTroop {
	if self.attacker == troop {
		return self.defender
	} else {
		return self.attacker
	}
}

// 计算战斗
func (self *Battle) Calc() *BattleResult {

	var time uint32
	br := &BattleResult{}

	for {
		// 打一轮
		self.attacker.Update(time)
		self.defender.Update(time)

		// 战斗是否结束
		if self.attacker.Lose() {
			br.Win = 0
			break
		} else if self.defender.Lose() {
			br.Win = 1
			break
		}

		time += 200
	}

	return br
}
