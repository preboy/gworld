package battle

import (
	"fmt"
)

// ==================================================
type BattleEvent uint32

const (
	_                  BattleEvent = 1 + iota
	BattleEvent_PreAtk             // 计算攻击之前 (累积光环的附加攻击)
	BattleEvent_AftDef             // 计算防御之后 (抵挡伤害)
)

type SkillCfg struct {
	Id, Lv uint32
}

type AuraCfg struct {
	Id, Lv uint32
}

type UnitBase interface {
	Name() string
}

// ==================================================

type BattleUnit struct {
	Base     UnitBase     // 父类
	Prop     *Property    // 战斗属性
	Troop    *BattleTroop // 队伍
	UnitType uint32       // 生物类型
	Dead     bool         // 是否死亡
	Rival    *BattleUnit  // 战场对手

	Skill_curr      *SkillBattle   // 当前正在释放技能
	Skill_comm      *SkillBattle   // 普通技能
	Skill_exclusive []*SkillBattle // 专有技能(比较猛的)
	Auras_basic     []*AuraBattle  // 英雄技能、角色加成、等等
	Auras_battle    []*AuraBattle  // 战斗中产生的光环(战斗结束之后保留)
	Auras_guarder   []*AuraBattle  // 辅助光环(战斗之前加，战斗之后结束，包括辅助、主帅加的)

	Skill_commander *SkillBattle // 主将技能(自用，二选一)
	Aura_commander  *AuraCfg     // 主将光环(二选一)
	Aura_guarder    *AuraCfg     // 辅将光环
}

func (self *BattleUnit) Name() string {
	name, _ := self.Troop.battle.get_unit_name(self)
	return name
}

func (self *BattleUnit) init_campaign(r *BattleUnit) {
	self.Rival = r
	self.Skill_curr = nil
	for _, s := range self.Skill_exclusive {
		s.Reset(false)
	}
	if self.Skill_comm != nil {
		self.Skill_comm.Reset(true)
	}
	for _, a := range self.Auras_basic {
		if a != nil {
			a.Reset()
		}
	}
	for _, a := range self.Auras_battle {
		if a != nil {
			a.Reset()
		}
	}

	// TODO 加辅将光环
	troop := self.Troop
	switch self {
	case troop.l_pioneer:
		{
			// 左先锋接受左辅将以及主帅的祝福
			g := troop.l_guarder
			a := g.Aura_guarder
			if g != nil && !g.Dead && a != nil && a.Id != 0 {
				self.Auras_guarder = append(self.Auras_guarder, NewAuraBattle(a.Id, a.Lv))
			}
			g = troop.commander
			a = g.Aura_commander
			if g != nil && !g.Dead && a != nil && a.Id != 0 {
				self.Auras_guarder = append(self.Auras_guarder, NewAuraBattle(a.Id, a.Lv))
			}
		}
	case troop.r_pioneer:
		{
			// 右先锋接受右辅将以及主帅的祝福
			g := troop.r_guarder
			a := g.Aura_guarder
			if g != nil && !g.Dead && a != nil && a.Id != 0 {
				self.Auras_guarder = append(self.Auras_guarder, NewAuraBattle(a.Id, a.Lv))
			}
			g = troop.commander
			a = g.Aura_commander
			if g != nil && !g.Dead && a != nil && a.Id != 0 {
				self.Auras_guarder = append(self.Auras_guarder, NewAuraBattle(a.Id, a.Lv))
			}
		}
	case troop.commander:
		{
			// 主帅接受两辅将以及自己的祝福
			g := troop.l_guarder
			a := g.Aura_guarder
			if g != nil && !g.Dead && a != nil && a.Id != 0 {
				self.Auras_guarder = append(self.Auras_guarder, NewAuraBattle(a.Id, a.Lv))
			}
			g = troop.r_guarder
			a = g.Aura_guarder
			if g != nil && !g.Dead && a != nil && a.Id != 0 {
				self.Auras_guarder = append(self.Auras_guarder, NewAuraBattle(a.Id, a.Lv))
			}
			a = self.Aura_commander
			if a != nil && a.Id != 0 {
				self.Auras_guarder = append(self.Auras_guarder, NewAuraBattle(a.Id, a.Lv))
			}
		}
	case troop.l_guarder:
		{
			// 辅将接受主将以及自己的祝福
			g := troop.commander
			a := g.Aura_commander
			if g != nil && !g.Dead && a != nil && a.Id != 0 {
				self.Auras_guarder = append(self.Auras_guarder, NewAuraBattle(a.Id, a.Lv))
			}
			a = self.Aura_guarder
			if a != nil && a.Id != 0 {
				self.Auras_guarder = append(self.Auras_guarder, NewAuraBattle(a.Id, a.Lv))
			}
		}
	case troop.r_guarder:
		{
			// 辅将接受主将以及自己的祝福
			g := troop.commander
			a := g.Aura_commander
			if g != nil && !g.Dead && a != nil && a.Id != 0 {
				self.Auras_guarder = append(self.Auras_guarder, NewAuraBattle(a.Id, a.Lv))
			}
			a = self.Aura_guarder
			if a != nil && a.Id != 0 {
				self.Auras_guarder = append(self.Auras_guarder, NewAuraBattle(a.Id, a.Lv))
			}
		}
	default:
	}

	for _, a := range self.Auras_guarder {
		if a != nil {
			a.Reset()
		}
	}
}

func (self *BattleUnit) clear_campaign() {
	self.Rival = nil
	self.Auras_guarder = nil
}

func (self *BattleUnit) Update(time int32) {
	if self.Dead {
		return
	}

	// skill update
	if self.Skill_curr == nil {
		for _, v := range self.Skill_exclusive {
			if !v.IsFree(time) {
				self.Skill_curr = v
				break
			}
		}
		if self.Skill_curr == nil {
			if self.Skill_comm.IsFree(time) {
				self.Skill_curr = self.Skill_comm
			}
		}
		if self.Skill_curr != nil {
			self.Skill_curr.Cast(self, time)
		}
	} else {
		self.Skill_curr.Update(time)
		if self.Skill_curr.IsFinish() {
			self.Skill_curr = nil
		}
	}

	// aura update
	for k, aura := range self.Auras_basic {
		if aura != nil {
			aura.Update(time)
			if aura.IsFinish() {
				self.Auras_basic[k] = nil
			}
		}
	}
	for k, aura := range self.Auras_battle {
		if aura != nil {
			aura.Update(time)
			if aura.IsFinish() {
				self.Auras_battle[k] = nil
			}
		}
	}
	for k, aura := range self.Auras_guarder {
		if aura != nil {
			aura.Update(time)
			if aura.IsFinish() {
				self.Auras_guarder[k] = nil
			}
		}
	}
}

func (self *BattleUnit) AddAuraBasic(caster *BattleUnit, id uint32, lv uint32) {
	aura := NewAuraBattle(id, lv)
	if aura == nil {
		return
	}
	aura.Init(caster, self)
	self.Auras_basic = append(self.Auras_basic, aura)
}

func (self *BattleUnit) AddAuraBattle(caster *BattleUnit, id uint32, lv uint32) {
	aura := NewAuraBattle(id, lv)
	if aura == nil {
		return
	}
	aura.Init(caster, self)
	self.Auras_battle = append(self.Auras_battle, aura)
}

func (self *BattleUnit) AddAuraGuarder(caster *BattleUnit, id uint32, lv uint32) {
	aura := NewAuraBattle(id, lv)
	if aura == nil {
		return
	}
	aura.Init(caster, self)
	self.Auras_guarder = append(self.Auras_guarder, aura)
}

func (self *BattleUnit) DelAura(id, lv uint32) {
	for k, aura := range self.Auras_basic {
		if aura.proto.Id == id && aura.proto.Lv == lv {
			self.Auras_basic[k] = nil
			return
		}
	}
	for k, aura := range self.Auras_battle {
		if aura.proto.Id == id && aura.proto.Lv == lv {
			self.Auras_battle[k] = nil
			return
		}
	}
	for k, aura := range self.Auras_guarder {
		if aura.proto.Id == id && aura.proto.Lv == lv {
			self.Auras_guarder[k] = nil
			return
		}
	}
}

// ==================================================

type BattleTroop struct {
	battle      *Battle
	is_attacker bool // 是否是挑起战事的一方
	is_massacre bool // 是否击杀对方guarder

	l_pioneer *BattleUnit // 左先锋
	r_pioneer *BattleUnit // 右先锋
	commander *BattleUnit // 主帅
	l_guarder *BattleUnit // 右辅助
	r_guarder *BattleUnit // 右辅助
}

func NewBattleTroop(l_pioneer, r_pioneer, l_guarder, commander, r_guarder *BattleUnit) *BattleTroop {
	if commander == nil {
		return nil
	}

	troop := &BattleTroop{
		commander: commander,
		l_pioneer: l_pioneer,
		r_pioneer: r_pioneer,
		l_guarder: l_guarder,
		r_guarder: r_guarder,
	}

	commander.Troop = troop
	l_pioneer.Troop = troop
	r_pioneer.Troop = troop
	l_guarder.Troop = troop
	r_guarder.Troop = troop

	return troop
}

func (self *BattleTroop) Lose() bool {
	return self.commander.Dead
}

// ==================================================
type BattleDetail struct {
	a    *BattleUnit // 攻击方出战者
	d    *BattleUnit // 防御方出战者
	a_hp uint32      // 攻击者用下的HP
	d_hp uint32      // 防御者用下的HP
}

type BattleResult struct {
	R       uint32 // 0:attacker负  1:attacker胜
	Details []*BattleDetail
}

func (self *BattleResult) ToMsg() string {
	return "s"
}

// ==================================================

type Battle struct {
	attacker  *BattleTroop
	defender  *BattleTroop
	result    BattleResult // 战斗结果
	campaigns int          // 战斗次数
}

func NewBattle(a *BattleTroop, d *BattleTroop) *Battle {
	if a == nil || d == nil {
		return nil
	}

	b := &Battle{
		attacker: a,
		defender: d,
	}
	a.battle = b
	d.battle = b

	a.is_attacker = true
	d.is_attacker = false

	return b
}

func (self *Battle) GetWinner() *BattleTroop {
	if self.attacker.Lose() {
		return self.defender
	} else if self.defender.Lose() {
		return self.attacker
	}
	return nil
}

func (self *Battle) get_unit_pos(u *BattleUnit) int {
	troop := u.Troop
	if troop == self.attacker {
		switch u {
		case troop.l_pioneer:
			return 1
		case troop.r_pioneer:
			return 2
		case troop.commander:
			return 3
		case troop.l_guarder:
			return 4
		case troop.r_guarder:
			return 5
		default:
			return 0
		}
	}
	if troop == self.defender {
		switch u {
		case troop.l_pioneer:
			return 6
		case troop.r_pioneer:
			return 7
		case troop.commander:
			return 8
		case troop.l_guarder:
			return 9
		case troop.r_guarder:
			return 10
		default:
			return 0
		}
	}
	return 0
}

func (self *Battle) get_unit_name(u *BattleUnit) (string, int) {
	troop := u.Troop
	if troop == self.attacker {
		switch u {
		case troop.l_pioneer:
			return fmt.Sprintf("%s[攻-%s]", u.Base.Name(), "左先锋"), 1
		case troop.r_pioneer:
			return fmt.Sprintf("%s[攻-%s]", u.Base.Name(), "右先锋"), 2
		case troop.commander:
			return fmt.Sprintf("%s[攻-%s]", u.Base.Name(), "主帅"), 3
		case troop.l_guarder:
			return fmt.Sprintf("%s[攻-%s]", u.Base.Name(), "左辅助"), 4
		case troop.r_guarder:
			return fmt.Sprintf("%s[攻-%s]", u.Base.Name(), "右辅助"), 5
		default:
			return "unknown[攻]", 0
		}
	}
	if troop == self.defender {
		switch u {
		case troop.l_pioneer:
			return fmt.Sprintf("%s[防-%s]", u.Base.Name(), "左先锋"), 6
		case troop.r_pioneer:
			return fmt.Sprintf("%s[防-%s]", u.Base.Name(), "右先锋"), 7
		case troop.commander:
			return fmt.Sprintf("%s[防-%s]", u.Base.Name(), "主帅"), 8
		case troop.l_guarder:
			return fmt.Sprintf("%s[防-%s]", u.Base.Name(), "左辅助"), 9
		case troop.r_guarder:
			return fmt.Sprintf("%s[防-%s]", u.Base.Name(), "右辅助"), 10
		default:
			return "unknown[防]", 0
		}
	}
	return "unknown", 0
}

// u: 只能为攻击方的单位
func (self *Battle) get_rival(u *BattleUnit) *BattleUnit {
	if u == nil || u.Dead {
		return nil
	}

	ta := self.attacker
	td := self.defender

	if u == ta.l_pioneer {
		r := td.r_pioneer
		if r != nil && !r.Dead {
			return r
		}
		r = td.r_guarder
		if ta.is_massacre && r != nil && !r.Dead {
			return r
		}
		r = td.commander
		if !r.Dead {
			return r
		}
	} else if u == ta.r_pioneer {
		r := td.l_pioneer
		if r != nil && r.Dead {
			return r
		}
		r = td.l_guarder
		if ta.is_massacre && r != nil && !r.Dead {
			return r
		}
		r = td.commander
		if !r.Dead {
			return r
		}
	} else if u == ta.commander {
		r := td.l_pioneer
		if r != nil && !r.Dead {
			return r
		}
		r = td.r_pioneer
		if r != nil && !r.Dead {
			return r
		}
		r = td.commander
		if !r.Dead {
			return r
		}
	}

	return nil
}

func (self *Battle) do_campaign(u *BattleUnit) {
	r := self.get_rival(u)
	if r == nil {
		fmt.Println("self.get_rival: nil", u)
		return
	}

	u.init_campaign(r)
	r.init_campaign(u)

	self.campaigns++

	var time int32
	var bout int32

	for {
		// 打一轮
		bout++
		fmt.Println("场次 回合 时间:", self.campaigns, bout, time)

		u.Update(time)
		r.Update(time)

		if u.Dead || r.Dead {
			break
		}

		// 超时(一分钟 600 = 60*1000/100)
		if bout >= 600 {
			fmt.Println("bout timeout !")
			u.Dead = true
			u.Prop.Hp_cur = 0
			break
		}
		time += 100
	}

	u.clear_campaign()
	r.clear_campaign()

	// 记录结果过程
	self.result.Details = append(self.result.Details, &BattleDetail{
		a:    u,
		d:    r,
		a_hp: u.Prop.Hp_cur,
		d_hp: r.Prop.Hp_cur,
	})

}

// 计算战斗
func (self *Battle) Calc() {

	l := self.attacker.l_pioneer
	r := self.attacker.r_pioneer
	c := self.attacker.commander

	for {
		if l != nil && !l.Dead && self.GetWinner() == nil {
			self.do_campaign(l)
			if self.GetWinner() != nil {
				return
			}
		}
		if r != nil && !r.Dead && self.GetWinner() == nil {
			self.do_campaign(r)
			if self.GetWinner() != nil {
				return
			}
		}
	}

	for !c.Dead && self.GetWinner() == nil {
		self.do_campaign(c)
		if self.GetWinner() != nil {
			return
		}
	}
}

func (self *Battle) GetResult() *BattleResult {
	if self.GetWinner() == self.attacker {
		self.result.R = 1
		fmt.Println("攻击者 胜 !!!")
	} else {
		self.result.R = 0
		fmt.Println("防御者 胜 !!!")
	}
	return &self.result
}
