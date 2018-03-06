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

type UnitBase interface {
	Name() string
}

// ==================================================

type BattleUnit struct {
	Base        UnitBase       // 父类
	Prop        *Property      // 战斗属性
	Troop       *BattleTroop   // 队伍
	UnitType    uint32         // 生物类型
	Auras       []*AuraBattle  // 光环(技能ID)
	Skill_extra []*SkillBattle // 额外技能(比较猛的)
	Skill_comm  *SkillBattle   // 普通技能
	Skill_curr  *SkillBattle   // 当前正在释放技能
	Dead        bool           // 是否死亡
	Rival       *BattleUnit    // 战场对手
}

func (self *BattleUnit) Name() string {
	if self.Troop.is_attacker {
		return fmt.Sprintf("(%s[%s][%p])", self.Base.Name(), "攻", self)
	} else {
		return fmt.Sprintf("(%s[%s][%p])", self.Base.Name(), "防", self)
	}
}

func (self *BattleUnit) init_campaign(r *BattleUnit) {
	self.Rival = r
	self.Skill_curr = nil
}

func (self *BattleUnit) clear_campaign() {
	self.Rival = nil
}

func (self *BattleUnit) Update(time uint32) {
	if self.Dead {
		return
	}

	if self.Skill_curr == nil {
		for _, v := range self.Skill_extra {
			if !v.IsFree(time) {
				continue
			}
			self.Skill_curr = v
			break
		}
		if self.Skill_curr == nil {
			self.Skill_curr = self.Skill_comm
		}
		self.Skill_curr.Cast(self, time)
	} else {
		self.Skill_curr.Update(time)
		if self.Skill_curr.IsFinish() {
			self.Skill_curr = nil
		}
	}

	for k, aura := range self.Auras {
		if aura != nil {
			aura.Update(time)
			if aura.IsFinish() {
				self.Auras[k] = nil
			}
		}
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
		if aura.proto.Id == id && aura.proto.Lv == lv {
			self.Auras[k] = nil
			break
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

	battle_troop := &BattleTroop{
		commander: commander,
		l_pioneer: l_pioneer,
		r_pioneer: r_pioneer,
		l_guarder: l_guarder,
		r_guarder: r_guarder,
	}

	commander.Troop = battle_troop
	l_pioneer.Troop = battle_troop
	r_pioneer.Troop = battle_troop
	l_guarder.Troop = battle_troop
	r_guarder.Troop = battle_troop

	return battle_troop
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

	var time uint32
	var bout uint32

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
			r.Dead = true
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
