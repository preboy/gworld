package battle

import (
	"fmt"
	"public/protocol/msg"
)

// ==================================================
type BattleEvent uint32
type CampaignEvent uint32

const (
	_                  BattleEvent = 1 + iota
	BattleEvent_PreAtk             // 计算攻击之前 (累积光环的附加攻击)
	BattleEvent_AftDef             // 计算防御之后 (抵挡伤害)
)

const (
	_                        CampaignEvent = 1 + iota // 战斗中的事件
	CampaignEvent_Cast                                // 释放技能
	CampaignEvent_Hurt                                // 受到伤害
	CampaignEvent_AuraGet                             // 得到光环
	CampaignEvent_AuraLose                            // 失去光环
	CampaignEvent_AuraEffect                          // 光环效果
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
	Troop    *BattleTroop // 队伍
	UnitType uint32       // 生物类型
	Pos      uint32       // 位置
	Rival    *BattleUnit  // 战场对手

	// 角色战斗属性
	Prop_base *Property // 战斗属性(可见部分[等级、品质、装备、被动技能])	不可变
	Prop_addi *Property // 战斗属性(附加部分[全局光环加攻防属性]，不可见)	-- 进入战斗之前计算一次
	Prop      *Property // 战斗属性之和
	Hp        uint32    // 当前HP

	// 战斗技能、光环
	Skill_curr      *BattleSkill   // 当前正在释放技能
	Skill_comm      *BattleSkill   // 普攻
	Skill_exclusive []*BattleSkill // 专有技能(比较猛的)
	Auras_battle    []*BattleAura  // 战斗光环，可被清除(战斗加成部分[技能释放之后产生的，祝福、诅咒]、辅助[辅将、主帅为战斗单位加成的])

	// 职业技能、光环
	career_general_skill *SkillCfg // 主帅技能(自用，二选一)
	career_general_aura  *AuraCfg  // 主帅光环(二选一)
	career_guarder_aura  *AuraCfg  // 辅将光环
}

func (self *BattleUnit) Name() string {
	return self.Troop.battle.get_unit_name(self)
}

func (self *BattleUnit) Dead() bool {
	return self.Hp == 0
}

func (self *BattleUnit) AddCampaignDetail(flag CampaignEvent, arg1, arg2, arg3, arg4 uint32) {
	self.Troop.battle.AddCampaignDetail(self, flag, arg1, arg2, arg3, arg4)
}

func (self *BattleUnit) CalcProp() {
	self.Prop = &Property{}
	self.Prop.Hp = self.Prop_base.Hp + self.Prop_addi.Hp
	self.Prop.Atk = self.Prop_base.Atk + self.Prop_addi.Atk
	self.Prop.Def = self.Prop_base.Def + self.Prop_addi.Def
	self.Prop.Crit = self.Prop_base.Crit + self.Prop_addi.Crit
	self.Prop.CritHurt = self.Prop_base.CritHurt + self.Prop_addi.CritHurt
	self.Hp = self.Prop.Hp
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

	// TODO 加辅将光环
	troop := self.Troop
	switch self {
	case troop.l_pioneer:
		fallthrough
	case troop.r_pioneer:
		{
			// 右先锋接受主帅的祝福
			u := troop.m_general
			if u != nil && !u.Dead() {
				a := u.career_general_aura
				if a != nil && a.Id != 0 {
					self.Auras_battle = append(self.Auras_battle, NewAuraBattle(a.Id, a.Lv, true))
				}
			}
		}
	case troop.m_general:
		{
			// 主帅接受两辅将以及自己的祝福
			u := troop.l_guarder
			if u != nil && !u.Dead() {
				a := u.career_guarder_aura
				if a != nil && a.Id != 0 {
					self.Auras_battle = append(self.Auras_battle, NewAuraBattle(a.Id, a.Lv, true))
				}
			}
			u = troop.r_guarder
			if u != nil && !u.Dead() {
				a := u.career_guarder_aura
				if a != nil && a.Id != 0 {
					self.Auras_battle = append(self.Auras_battle, NewAuraBattle(a.Id, a.Lv, true))
				}
			}
			a := self.career_general_aura
			if a != nil && a.Id != 0 {
				self.Auras_battle = append(self.Auras_battle, NewAuraBattle(a.Id, a.Lv, true))
			}
		}
	case troop.l_guarder:
		fallthrough
	case troop.r_guarder:
		{
			// 右辅将接受主将的祝福
			u := troop.m_general
			if u != nil && !u.Dead() {
				a := u.career_general_aura
				if a != nil && a.Id != 0 {
					self.Auras_battle = append(self.Auras_battle, NewAuraBattle(a.Id, a.Lv, true))
				}
			}
		}
	default:
	}

	for _, a := range self.Auras_battle {
		if a != nil {
			a.Reset()
			self.AddCampaignDetail(CampaignEvent_AuraGet, a.proto.Id, a.proto.Level, 0, 0)
		}
	}
}

func (self *BattleUnit) clear_campaign() {
	self.Rival = nil
	for i, a := range self.Auras_battle {
		if a != nil && a.once {
			self.Auras_battle[i] = nil
		}
	}
}

func (self *BattleUnit) Update(time int32) {
	if self.Dead() {
		return
	}

	// skill update
	if self.Skill_curr == nil {
		for _, v := range self.Skill_exclusive {
			if v.IsFree(time) {
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

	for k, aura := range self.Auras_battle {
		if aura != nil {
			aura.Update(time)
			if aura.IsFinish() {
				self.Auras_battle[k] = nil
			}
		}
	}

}

func (self *BattleUnit) AddAura(caster *BattleUnit, id uint32, lv uint32) {
	aura := NewAuraBattle(id, lv, false)
	if aura == nil {
		return
	}
	aura.Init(caster, self)
	self.Auras_battle = append(self.Auras_battle, aura)
	self.AddCampaignDetail(CampaignEvent_AuraGet, id, lv, 0, 0)
}

func (self *BattleUnit) DelAura(id, lv uint32) {
	for k, aura := range self.Auras_battle {
		if aura.proto.Id == id && aura.proto.Level == lv {
			self.Auras_battle[k] = nil
			self.AddCampaignDetail(CampaignEvent_AuraLose, id, lv, 0, 0)
			return
		}
	}
}

func (self *BattleUnit) ToMsg() *msg.BattleUnit {
	u := &msg.BattleUnit{}

	u.Type = self.UnitType
	u.Id = self.Id
	u.Lv = self.Lv
	u.Pos = self.Pos
	u.Hp = self.Prop.Hp
	u.Atk = self.Prop.Atk
	u.Def = self.Prop.Def
	u.Crit = self.Prop.Crit
	u.CritHurt = self.Prop.CritHurt

	if self.Skill_comm != nil {
		u.Comm = &msg.BattleSkill{
			Id: self.Skill_comm.proto.Id,
			Lv: self.Skill_comm.proto.Level,
		}
	}

	for _, v := range self.Skill_exclusive {
		u.Skill = append(u.Skill, &msg.BattleSkill{
			Id: v.proto.Id,
			Lv: v.proto.Level,
		})
	}

	if self.career_general_skill != nil {
		u.CareerGeneralSkill = &msg.BattleSkill{
			Id: self.career_general_skill.Id,
			Lv: self.career_general_skill.Lv,
		}
	}
	if self.career_general_aura != nil {
		u.CareerGeneralAura = &msg.BattleAura{
			Id: self.career_general_aura.Id,
			Lv: self.career_general_aura.Lv,
		}
	}
	if self.career_guarder_aura != nil {
		u.CareerGuarderAura = &msg.BattleAura{
			Id: self.career_guarder_aura.Id,
			Lv: self.career_guarder_aura.Lv,
		}
	}

	return u
}

// ==================================================

type BattleTroop struct {
	battle      *Battle
	is_attacker bool // 是否是挑起战事的一方
	is_massacre bool // 是否击杀对方guarder

	l_pioneer *BattleUnit // 左先锋
	r_pioneer *BattleUnit // 右先锋
	m_general *BattleUnit // 主帅
	l_guarder *BattleUnit // 右辅助
	r_guarder *BattleUnit // 右辅助
}

func NewBattleTroop(l_pioneer, r_pioneer, l_guarder, m_general, r_guarder *BattleUnit) *BattleTroop {
	if m_general == nil {
		return nil
	}

	troop := &BattleTroop{
		m_general: m_general,
		l_pioneer: l_pioneer,
		r_pioneer: r_pioneer,
		l_guarder: l_guarder,
		r_guarder: r_guarder,
	}

	m_general.Troop = troop
	l_pioneer.Troop = troop
	r_pioneer.Troop = troop
	l_guarder.Troop = troop
	r_guarder.Troop = troop

	// 加主帅技能
	s := troop.m_general.career_general_skill
	if s != nil {
		skill := NewSkillBattle(s.Id, s.Lv)
		troop.m_general.Skill_exclusive = append(troop.m_general.Skill_exclusive, skill)
	}

	return troop
}

func (self *BattleTroop) Lose() bool {
	return self.m_general.Dead()
}

// ==================================================

type Battle struct {
	attacker  *BattleTroop
	defender  *BattleTroop
	campaigns []*msg.BattleCampaign
	campaign  *msg.BattleCampaign
	R         uint32 // 0:attacker负  1:attacker胜
	time      int32
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

	if a.l_pioneer != nil {
		a.l_pioneer.Pos = 1
	}
	if a.r_pioneer != nil {
		a.r_pioneer.Pos = 2
	}
	if a.m_general != nil {
		a.m_general.Pos = 3
	}
	if a.l_guarder != nil {
		a.l_guarder.Pos = 4
	}
	if a.r_guarder != nil {
		a.r_guarder.Pos = 5
	}

	if d.l_pioneer != nil {
		d.l_pioneer.Pos = 6
	}
	if d.r_pioneer != nil {
		d.r_pioneer.Pos = 7
	}
	if d.m_general != nil {
		d.m_general.Pos = 8
	}
	if d.l_guarder != nil {
		d.l_guarder.Pos = 9
	}
	if d.r_guarder != nil {
		d.r_guarder.Pos = 10
	}

	return b
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
		case troop.l_pioneer:
			return fmt.Sprintf("%s[攻-%s]", u.Base.Name(), "左先锋")
		case troop.r_pioneer:
			return fmt.Sprintf("%s[攻-%s]", u.Base.Name(), "右先锋")
		case troop.m_general:
			return fmt.Sprintf("%s[攻-%s]", u.Base.Name(), "主帅")
		case troop.l_guarder:
			return fmt.Sprintf("%s[攻-%s]", u.Base.Name(), "左辅助")
		case troop.r_guarder:
			return fmt.Sprintf("%s[攻-%s]", u.Base.Name(), "右辅助")
		default:
			return "unknown[攻]"
		}
	}
	if troop == self.defender {
		switch u {
		case troop.l_pioneer:
			return fmt.Sprintf("%s[防-%s]", u.Base.Name(), "左先锋")
		case troop.r_pioneer:
			return fmt.Sprintf("%s[防-%s]", u.Base.Name(), "右先锋")
		case troop.m_general:
			return fmt.Sprintf("%s[防-%s]", u.Base.Name(), "主帅")
		case troop.l_guarder:
			return fmt.Sprintf("%s[防-%s]", u.Base.Name(), "左辅助")
		case troop.r_guarder:
			return fmt.Sprintf("%s[防-%s]", u.Base.Name(), "右辅助")
		default:
			return "unknown[防]"
		}
	}
	return "unknown"
}

// u: 只能为攻击方的单位
func (self *Battle) get_rival(u *BattleUnit) *BattleUnit {
	if u == nil || u.Dead() {
		return nil
	}

	ta := self.attacker
	td := self.defender

	if u == ta.l_pioneer {
		r := td.r_pioneer
		if r != nil && !r.Dead() {
			return r
		}
		r = td.r_guarder
		if ta.is_massacre && r != nil && !r.Dead() {
			return r
		}
		r = td.m_general
		if !r.Dead() {
			return r
		}
	} else if u == ta.r_pioneer {
		r := td.l_pioneer
		if r != nil && !r.Dead() {
			return r
		}
		r = td.l_guarder
		if ta.is_massacre && r != nil && !r.Dead() {
			return r
		}
		r = td.m_general
		if !r.Dead() {
			return r
		}
	} else if u == ta.m_general {
		r := td.l_pioneer
		if r != nil && !r.Dead() {
			return r
		}
		r = td.r_pioneer
		if r != nil && !r.Dead() {
			return r
		}
		r = td.m_general
		if !r.Dead() {
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

	self.campaign = &msg.BattleCampaign{
		APos: u.Pos,
		DPos: r.Pos,
		AHpS: u.Hp,
		DHpS: r.Hp,
	}
	self.campaign.Details = make([]*msg.CampaignDetail, 0, 0x100)
	self.campaigns = append(self.campaigns, self.campaign)
	self.time = 0

	u.init_campaign(r)
	r.init_campaign(u)

	camp_index := len(self.campaigns)
	fmt.Println("============== Campaign Begin ==============", camp_index)
	fmt.Println(u.Name(), " VS ", r.Name())

	var bout int32
	for {
		bout++
		fmt.Println("场次 回合 时间:", camp_index, bout, self.time)

		u.Update(self.time)
		r.Update(self.time)

		if u.Dead() || r.Dead() {
			break
		}

		// 超时(一分钟 600 = 60*1000/100)
		if bout >= 600 {
			fmt.Println("bout timeout !")
			u.Hp = 0
			break
		}
		self.time += 100
	}

	if u.Dead() {
		fmt.Println(u.Name(), " [输给了] ", r.Name())
	} else {
		fmt.Println(u.Name(), " [战胜了] ", r.Name())
	}

	fmt.Println("============== campaign end ==============", camp_index)

	// 记录结果过程
	self.campaign.AHpE = u.Hp
	self.campaign.DHpE = r.Hp

	u.clear_campaign()
	r.clear_campaign()

	self.time = 0
	self.campaign = nil
}

// 计算战斗
func (self *Battle) Calc() {

	l := self.attacker.l_pioneer
	r := self.attacker.r_pioneer
	g := self.attacker.m_general

	for {
		idle := true
		if l != nil && !l.Dead() && self.InBattle() {
			idle = false
			self.do_campaign(l)
			if !self.InBattle() {
				break
			}
		}
		if r != nil && !r.Dead() && self.InBattle() {
			idle = false
			self.do_campaign(r)
			if !self.InBattle() {
				break
			}
		}
		if idle {
			break
		}
	}

	for self.InBattle() {
		self.do_campaign(g)
	}

	if self.GetWinner() == self.attacker {
		self.R = 1
		fmt.Println("攻击方 胜 !!!", self.campaigns)
	} else {
		self.R = 0
		fmt.Println("防御方 胜 !!!", self.campaigns)
	}

}

func (self *Battle) AddCampaignDetail(u *BattleUnit, flag CampaignEvent, arg1, arg2, arg3, arg4 uint32) {
	self.campaign.Details = append(self.campaign.Details, &msg.CampaignDetail{
		Host: u.Pos,
		Time: uint32(self.time),
		Flag: uint32(flag),
		Arg1: 0,
		Arg2: 0,
		Arg3: 0,
		Arg4: 0,
	})
}

func (self *Battle) GetResult() uint32 {
	return self.R
}

func (self *Battle) ToMsg() *msg.BattleResult {
	r := &msg.BattleResult{}

	r.Win = self.GetResult() == 1
	r.Camps = self.campaigns
	r.Units = make([]*msg.BattleUnit, 0, 10)

	u := self.attacker.l_pioneer
	if u != nil {
		r.Units = append(r.Units, u.ToMsg())
	}
	u = self.attacker.r_pioneer
	if u != nil {
		r.Units = append(r.Units, u.ToMsg())
	}
	u = self.attacker.m_general
	if u != nil {
		r.Units = append(r.Units, u.ToMsg())
	}
	u = self.attacker.l_guarder
	if u != nil {
		r.Units = append(r.Units, u.ToMsg())
	}
	u = self.attacker.r_guarder
	if u != nil {
		r.Units = append(r.Units, u.ToMsg())
	}

	u = self.defender.l_pioneer
	if u != nil {
		r.Units = append(r.Units, u.ToMsg())
	}
	u = self.defender.r_pioneer
	if u != nil {
		r.Units = append(r.Units, u.ToMsg())
	}
	u = self.defender.m_general
	if u != nil {
		r.Units = append(r.Units, u.ToMsg())
	}
	u = self.defender.l_guarder
	if u != nil {
		r.Units = append(r.Units, u.ToMsg())
	}
	u = self.defender.r_guarder
	if u != nil {
		r.Units = append(r.Units, u.ToMsg())
	}

	return r
}
