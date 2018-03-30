package battle

import (
	"core/log"
	"server/game/config"
)

type BattleAura struct {
	owner       *BattleUnit
	caster      *BattleUnit
	proto       *config.AuraProto
	script      AuraScript
	start_time  int32
	update_time int32 // 对于有update的技能，记录上次时间
	start       bool  // 就否初始化完成
	finish      bool  // 是否完成
	once        bool  // 是否一次性光环(战斗结束就删除,辅助光环使用此功能)
}

func NewAuraBattle(id, lv uint32, once bool) *BattleAura {
	proto := config.GetAuraProto(id, lv)
	if proto == nil {
		log.Error("NewAuraBattle Failed:", id, lv)
		return nil
	}

	ab := &BattleAura{
		proto:  proto,
		script: create_aura_script(proto),
	}

	ab.once = once

	return ab
}

func (self *BattleAura) Init(caster, owner *BattleUnit) {
	self.owner = owner
	self.caster = caster
}

// 每场战斗开始时调用
func (self *BattleAura) Reset() {
	self.start = false
	self.finish = false
	self.start_time = 0
	self.update_time = 0
}

func (self *BattleAura) Update(time int32) {
	if self.finish {
		return
	}

	if !self.start {
		self.start = true
		self.start_time = time
		self.update_time = time
		self.onStart()
	}
	if self.proto.Itv_t != 0 {
		if time-self.update_time >= self.proto.Itv_t {
			self.onUpdate()
			self.update_time = time
		}
	}
	if time-self.start_time >= self.proto.Last_t {
		self.onFinish()
		self.finish = true
	}
}

func (self *BattleAura) onStart() {
	if self.script != nil {
		self.script.OnStart(self)
	}
}

func (self *BattleAura) onUpdate() {
	if self.script != nil {
		self.script.OnUpdate(self)
	}
}

func (self *BattleAura) onFinish() {
	if self.script != nil {
		self.script.OnFinish(self)
	}
}

func (self *BattleAura) IsFinish() bool {
	return self.finish
}

func (self *BattleAura) OnEvent(evt BattleEventType, ctx *SkillContext) {
	if self.script != nil {
		self.script.OnEvent(evt, self, ctx)
	}
}
