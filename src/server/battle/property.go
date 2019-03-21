package battle

import (
	"core/log"
	"server/config"
)

const (
	PartBase  uint32 = 0
	PartPerc  uint32 = 1
	PartExtra uint32 = 2
)

const (
	PropType_HP   = 0 // HP
	PropType_Apm  = 1 // 速度
	PropType_Atk  = 2 // 攻击
	PropType_Def  = 3 // 防御
	PropType_Crit = 4 // 暴击
	PropType_Hurt = 5 // 暴伤
)

const (
	C_Property_Number = 6 // 属性总数量
)

// ----------------------------------------------------------------------------

type Property struct {
	base  float64 // 基础属性
	perc  float64 // 百分比加成
	extra float64 // 额外追加
	total float64 // 总属性
	daity bool
}

type PropertyGroup [C_Property_Number]Property

func NewPropertyGroup() *PropertyGroup {
	return new(PropertyGroup)
}

// ----------------------------------------------------------------------------

func (self *Property) Clear() {
	self.base = 0
	self.perc = 0
	self.extra = 0
	self.total = 0
	self.daity = true
}

func (self *Property) add(part uint32, val float64, add bool) {
	if val == 0 || part > PartExtra {
		return
	}

	if !add {
		val = -val
	}

	switch part {
	case PartBase:
		self.base += val
	case PartPerc:
		self.perc += val
	case PartExtra:
		self.extra += val
	}

	self.daity = true
}

func (self *Property) Calc() {
	if !self.daity {
		return
	}

	self.daity = false
	self.total = self.base*(1+self.perc/100) + self.extra

	if self.total < 0 {
		log.Error("Property.Calc ERROR: %f, %f, %f, %f", self.base, self.perc, self.extra, self.total)
		self.total = 0
	}
}

func (self *Property) Value() float64 {
	if self.daity {
		self.Calc()
	}

	return self.total
}

// ----------------------------------------------------------------------------

func (self *PropertyGroup) AddProps(props []*config.PropConf) {
	for _, p := range props {
		if p.Id <= PropType_Hurt {
			self[p.Id].add(p.Part, p.Val, true)
		}
	}
}

func (self *PropertyGroup) SubProps(props []*config.PropConf) {
	for _, p := range props {
		if p.Id <= PropType_Hurt {
			self[p.Id].add(p.Part, p.Val, false)
		}
	}
}

func (self *PropertyGroup) Calc() {
	for i := 0; i < C_Property_Number; i++ {
		self[i].Calc()
	}
}

func (self *PropertyGroup) Value(id uint32) float64 {
	if id <= PropType_Hurt {
		return self[id].Value()
	}
	return 0
}
