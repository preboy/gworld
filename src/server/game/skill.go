package game

import (
	"server/game/config"
)

// 主动技能
type Skill struct {
	Id       uint32 `bson:id"`        // ID
	Level    uint32 `bson:level"`     // 等级
	EffectId uint32 `bson:effect_id"` // 技能附加效果ID
}

// ==================================================

type SkillBattle struct {
	sp          *config.SkillProto
	owner       *BattleUnit //技能拥有者
	start_time  uint32      // 技能释放时间(包含释放过程)，用于计算CD
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
	self.finish = false
	self.owner = u
	self.start_time = time
	self.update_time = time
	self.onStart()
}

func (self *SkillBattle) Update(time uint32) {
	if self.sp.Itv_t != 0 {
		if time-self.update_time > self.sp.Itv_t {
			self.onUpdate()
			self.update_time = time
		}
	}
	if time-self.start_time >= self.sp.Last_t {
		self.onFinish()
		self.finish = true
		self.owner = nil
	}
}

func (self *SkillBattle) InCD(time uint32) bool {
	return time-self.start_time < self.sp.Cd_t
}

func (self *SkillBattle) IsFinish() bool {
	return self.finish
}

func (self *SkillBattle) onStart() {
	// nothing to do
}

func (self *SkillBattle) onUpdate() {
	// 攻击、加光环

}

func (self *SkillBattle) onFinish() {
	if self.sp.Itv_t == 0 {
		self.onUpdate()
	}
}
