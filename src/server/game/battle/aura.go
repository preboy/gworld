package battle

import (
	"server/game/config"
)

type AuraBattle struct {
	owner       *BattleUnit
	caster      *BattleUnit
	proto       *config.AuraProto
	eff         Effect
	start_time  uint32
	update_time uint32 // 对于有update的技能，记录上次时间
	start       bool   // 就否初始化完成
	finish      bool   // 是否完成
}

func NewAuraBattle(id, lv uint32) *AuraBattle {
	proto := config.GetAuraProtoConf().GetAuraProto(id, lv)
	if proto == nil {
		return nil
	}
	ab := &AuraBattle{
		proto: proto,
	}
	return ab
}

func (self *AuraBattle) Init(caster, owner *BattleUnit) {
	self.owner = owner
	self.caster = caster
}

func (self *AuraBattle) Update(time uint32) {
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
	if self.eff != nil {
		self.eff.OnStart(self)
	}
}

func (self *AuraBattle) onUpdate() {
	if self.eff != nil {
		self.eff.OnUpdate(self)
	}
}

func (self *AuraBattle) onFinish() {
	if self.eff != nil {
		self.eff.OnFinish(self)
	}
}

func (self *AuraBattle) IsFinish() bool {
	return self.finish
}

func (self *AuraBattle) OnEvent(evt BattleEvent, sc *SkillContext) {
	if self.eff != nil {
		self.eff.OnEvent(evt, self, sc)
	}
}
