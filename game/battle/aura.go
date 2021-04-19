package battle

import (
	"gworld/game/config"
)

type BattleAura struct {
	owner       *BattleUnit
	caster      *BattleUnit
	proto       *config.Aura
	script      AuraScript
	start_time  uint32
	update_time uint32 // 对于有update的技能，记录上次时间
	start       bool   // 就否初始化完成
	finish      bool   // 是否完成
}

func NewAuraBattle(id, lv uint32) *BattleAura {
	proto := config.AuraProtoConf.Query(id, lv)
	return &BattleAura{
		proto:  proto,
		script: create_script_object(proto),
	}
}

func (self *BattleAura) Init(caster, owner *BattleUnit) {
	self.owner = owner
	self.caster = caster
}

func (self *BattleAura) IsFinish() bool {
	return self.finish
}

func (self *BattleAura) Update(time uint32) {
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

func (self *BattleAura) OnEvent(evt BattleCalcEvent, ctx *SkillContext) {
	if self.script != nil {
		self.script.OnEvent(self, evt, ctx)
	}
}

func (self *BattleAura) onFinish() {
	if self.script != nil {
		self.script.OnFinish(self)
	}
}
