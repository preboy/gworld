package battle

import (
	"fmt"
	"public/protocol/msg"
)

// ==================================================
type BattleCalcEvent uint32 // 用于战斗计算

const (
	_          BattleCalcEvent = 0 + iota
	BCE_PreAtk                 // 计算攻击之前 (累积光环的附加攻击)
	BCE_Damage                 // 发出伤害
	BCE_AftDef                 // 计算防御之后 (抵挡伤害)
)

const (
	MAX_TROOP_MEMBER = 6
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
	Id       uint32       // ID
	Lv       uint32       // 等级
	UnitType uint32       // 生物类型
	Pos      uint32       // 位置 		pos start 1 to 12
	Troop    *BattleTroop // 队伍
	dead     bool         // 是否死亡

	// 角色战斗属性
	Prop_base *Property // 战斗属性(等级、品质、装备、其它养成)							-- 进入战斗前计算
	Prop_addi *Property // 战斗属性(附加部分[BUFF、被动技能加成，以Prop_base为基数])	-- 进入战斗前计算
	Prop_aura *Property // 战斗中光环所加的属性										-- 战斗中产生
	Prop      *Property // 战斗属性之和

	Hp   int    // 当前HP
	Rst  uint32 // 分钟 / Apm
	Last uint32 // 上一次时间

	// 战斗技能、光环
	Skill_curr   *BattleSkill   // 当前正在释放技能
	Skill_comm   *BattleSkill   // 普攻
	Skill_battle []*BattleSkill // 战斗技能
	Auras_battle []*BattleAura  // 战斗光环
}

func (self *BattleUnit) GetBattle() *Battle {
	return self.Troop.battle
}

func (self *BattleUnit) Name() string {
	return self.Troop.battle.get_unit_name(self)
}

func (self *BattleUnit) Dead() bool {
	return self.dead
}

func (self *BattleUnit) Init(pos int) {
	self.Pos = uint32(pos)
	self.Prop_aura = &Property{}
	self.Prop = &Property{}
	self.CalcProp()
	self.Hp = int(self.Prop.Hp)
}

func (self *BattleUnit) CalcProp() {
	self.Prop.Clear()
	self.Prop.AddProperty(self.Prop_base)
	self.Prop.AddProperty(self.Prop_addi)
	self.Prop.AddProperty(self.Prop_aura)
	self.Rst = uint32(60000 / self.Prop.Apm)
}

func (self *BattleUnit) Update(time uint32) {
	if self.Skill_curr == nil {
		// apm checking
		if time < self.Last+self.Rst {
			return
		}
		// 技能
		for _, v := range self.Skill_battle {
			if v.IsFree(time) {
				self.Skill_curr = v
				break
			}
		}
		// 普攻
		if self.Skill_curr == nil {
			if self.Skill_comm.IsFree(time) {
				self.Skill_curr = self.Skill_comm
			}
		}
		// 释放
		if self.Skill_curr != nil {
			self.Last = time
			self.Skill_curr.Cast(self, time)
		}
	} else {
		self.Skill_curr.Update(time)
		if self.Skill_curr.IsFinish() {
			self.Skill_curr = nil
		}
	}

	// update aura
	for k, aura := range self.Auras_battle {
		if aura != nil {
			aura.Update(time)
			if aura.IsFinish() {
				self.Auras_battle[k] = nil
			}
		}
	}

}

func (self *BattleUnit) UpdateLife(time uint32) {
	self.dead = self.Hp <= 0
}

func (self *BattleUnit) AddAura(caster *BattleUnit, id uint32, lv uint32) {
	aura := NewAuraBattle(id, lv)
	if aura == nil {
		return
	}
	aura.Init(caster, self)
	self.Auras_battle = append(self.Auras_battle, aura)
	self.GetBattle().BattlePlayEvent_Aura(self, caster, id, lv, true)
}

func (self *BattleUnit) DelAura(id, lv uint32) {
	for k, aura := range self.Auras_battle {
		if aura.proto.Id == id && aura.proto.Level == lv {
			self.Auras_battle[k] = nil
			self.GetBattle().BattlePlayEvent_Aura(self, aura.caster, id, lv, false)
			break
		}
	}
}

func (self *BattleUnit) ToMsg() *msg.BattleUnit {
	u := &msg.BattleUnit{
		Type: self.UnitType,
		Id:   self.Id,
		Lv:   self.Lv,
		Pos:  self.Pos,
		Apm:  uint32(self.Prop.Apm),
		Atk:  uint32(self.Prop.Atk),
		Def:  uint32(self.Prop.Def),
		Hp:   uint32(self.Prop.Hp),
		Crit: uint32(self.Prop.Crit),
		Hurt: uint32(self.Prop.Hurt),
		Comm: &msg.BattleSkill{self.Skill_comm.proto.Id, self.Skill_comm.proto.Level},
	}

	if self.Troop.IsAttacker() {
		u.Attacker = 1
	}

	for _, skill := range self.Skill_battle {
		u.Skill = append(u.Skill, &msg.BattleSkill{
			skill.proto.Id,
			skill.proto.Level,
		})
	}

	return u
}

// ==================================================

type BattleTroop struct {
	battle   *Battle
	attacker bool
	members  [MAX_TROOP_MEMBER]*BattleUnit // 从第一排到第二排行，从左到右
}

func NewBattleTroop(members ...*BattleUnit) *BattleTroop {
	troop := &BattleTroop{}

	l := len(members)
	if l > MAX_TROOP_MEMBER {
		l = MAX_TROOP_MEMBER
	}

	for i := 0; i < l; i++ {
		members[i].Troop = troop
		troop.members[i] = members[i]
	}

	return troop
}

func (self *BattleTroop) Init(attacker bool) {
	self.attacker = attacker
	for i := 0; i < MAX_TROOP_MEMBER; i++ {
		if self.members[i] != nil {
			if attacker {
				self.members[i].Init(i + 1)
			} else {
				self.members[i].Init(i + 1 + MAX_TROOP_MEMBER)
			}
		}
	}
}

func (self *BattleTroop) Lose() (ret bool) {
	ret = true
	for i := 0; i < MAX_TROOP_MEMBER; i++ {
		member := self.members[i]
		if member != nil && !member.Dead() {
			ret = false
			return
		}
	}
	return
}

func (self *BattleTroop) IsAttacker() bool {
	return self.attacker
}

func (self *BattleTroop) IsDefender() bool {
	return !self.attacker
}

// 敌方队伍
func (self *BattleTroop) GetRival() *BattleTroop {
	if self.attacker {
		return self.battle.GetDefender()
	} else {
		return self.battle.GetAttacher()
	}
}

// ==================================================

type Battle struct {
	attacker *BattleTroop
	defender *BattleTroop
	R        uint32 // 0:attacker负  1:attacker胜
	time     uint32
	Result   msg.BattleResult
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

	b.Init()

	return b
}

func (self *Battle) Init() {
	self.attacker.Init(true)
	self.defender.Init(false)

	// 装玩家
	for _, u := range self.attacker.members {
		self.Result.Units = append(self.Result.Units, u.ToMsg())
	}
	for _, u := range self.defender.members {
		self.Result.Units = append(self.Result.Units, u.ToMsg())
	}
}

func (self *Battle) GetAttacher() *BattleTroop {
	return self.attacker
}

func (self *Battle) GetDefender() *BattleTroop {
	return self.defender
}

func (self *Battle) InBattle() bool {
	return self.GetWinner() == nil
}

func (self *Battle) GetWinner() *BattleTroop {
	if self.attacker.Lose() {
		return self.defender
	} else if self.defender.Lose() {
		return self.attacker
	}
	return nil
}

func (self *Battle) get_unit_name(u *BattleUnit) string {
	troop := u.Troop
	if troop == self.attacker {
		switch u {
		case troop.members[0]:
			return fmt.Sprintf("%s[攻-%s]", u.Base.Name(), "一左")
		case troop.members[1]:
			return fmt.Sprintf("%s[攻-%s]", u.Base.Name(), "一中")
		case troop.members[2]:
			return fmt.Sprintf("%s[攻-%s]", u.Base.Name(), "一右")
		case troop.members[3]:
			return fmt.Sprintf("%s[攻-%s]", u.Base.Name(), "二左")
		case troop.members[4]:
			return fmt.Sprintf("%s[攻-%s]", u.Base.Name(), "二中")
		case troop.members[5]:
			return fmt.Sprintf("%s[攻-%s]", u.Base.Name(), "二右")
		}
	}

	if troop == self.defender {
		switch u {
		case troop.members[0]:
			return fmt.Sprintf("%s[防-%s]", u.Base.Name(), "一左")
		case troop.members[1]:
			return fmt.Sprintf("%s[防-%s]", u.Base.Name(), "一中")
		case troop.members[2]:
			return fmt.Sprintf("%s[防-%s]", u.Base.Name(), "一右")
		case troop.members[3]:
			return fmt.Sprintf("%s[防-%s]", u.Base.Name(), "二左")
		case troop.members[4]:
			return fmt.Sprintf("%s[防-%s]", u.Base.Name(), "二中")
		case troop.members[5]:
			return fmt.Sprintf("%s[防-%s]", u.Base.Name(), "二右")
		}
	}

	return "unknown"
}

// 计算战斗
func (self *Battle) Calc() {

	var bout int32

	for {
		if self.attacker.Lose() {
			self.Result.Win = false
			break
		}
		if self.defender.Lose() {
			self.Result.Win = true
			break
		}

		bout++

		// 战斗
		for _, u := range self.attacker.members {
			if u != nil && !u.Dead() {
				u.Update(self.time)
			}
		}
		for _, u := range self.attacker.members {
			if u != nil && !u.Dead() {
				u.Update(self.time)
			}
		}

		// 判生死
		for _, u := range self.attacker.members {
			if u != nil && !u.Dead() {
				u.UpdateLife(self.time)
			}
		}
		for _, u := range self.attacker.members {
			if u != nil && !u.Dead() {
				u.UpdateLife(self.time)
			}
		}

		// 超时检测(5分钟 3000 = 5*60*1000/100)
		if bout >= 3000 {
			fmt.Println("bout timeout !")
			self.Result.Win = false
			break
		}

		fmt.Println("fuckyou ", self.defender, self.attacker)

		self.time += 100
	}

}

func (self *Battle) GetResult() bool {
	return self.Result.Win
}

func (self *Battle) ToMsg() *msg.BattleResult {
	return &self.Result
}
