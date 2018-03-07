package battle

import (
	"core/log"
	"server/game/config"
)

type AuraBattle struct {
	owner       *BattleUnit
	caster      *BattleUnit
	proto       *config.AuraProto
	script      AuraScript
	start_time  int32
	update_time int32 // 对于有update的技能，记录上次时间
	start       bool  // 就否初始化完成
	finish      bool  // 是否完成
}

func NewAuraBattle(id, lv uint32) *AuraBattle {
	proto := config.GetAuraProtoConf().GetAuraProto(id, lv)
	if proto == nil {
		log.Error("NewAuraBattle Failed:", id, lv)
		return nil
	}
	ab := &AuraBattle{
		proto:  proto,
		script: create_aura_script(proto),
	}
	return ab
}

func (self *AuraBattle) Init(caster, owner *BattleUnit) {
	self.owner = owner
	self.caster = caster
}

// 每场战斗开始时调用
func (self *AuraBattle) Reset() {
	self.start = false
	self.finish = false
	self.start_time = 0
	self.update_time = 0
}

func (self *AuraBattle) Update(time int32) {
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
		if time-self.update_time > self.proto.Itv_t {
			self.onUpdate()
			self.update_time = time
		}
	}
	if time-self.start_time >= self.proto.Last_t {
		self.onFinish()
		self.finish = true
	}
}

func (self *AuraBattle) onStart() {
	if self.script != nil {
		self.script.OnStart(self)
	}
}

func (self *AuraBattle) onUpdate() {
	if self.script != nil {
		self.script.OnUpdate(self)
	}
}

func (self *AuraBattle) onFinish() {
	if self.script != nil {
		self.script.OnFinish(self)
	}
}

func (self *AuraBattle) IsFinish() bool {
	return self.finish
}

func (self *AuraBattle) OnEvent(evt BattleEvent, ctx *SkillContext) {
	if self.script != nil {
		self.script.OnEvent(evt, self, ctx)
	}
}
