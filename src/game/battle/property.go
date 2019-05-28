package battle

import (
	"fmt"

	"core/log"
	"game/config"
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

var type_name = map[int]string{
	PropType_HP:   "hp",
	PropType_Apm:  "apm",
	PropType_Atk:  "atk",
	PropType_Def:  "def",
	PropType_Crit: "crit",
	PropType_Hurt: "hurt",
}

// ----------------------------------------------------------------------------

type Property struct {
	base  float64 // 基础属性
	perc  float64 // 百分比加成
	extra float64 // 额外追加
	total float64 // 总属性
	dirty bool
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
	self.dirty = true
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

	self.dirty = true
}

func (self *Property) Calc() {
	if !self.dirty {
		return
	}

	self.dirty = false
	self.total = self.base*(1+self.perc/100) + self.extra

	if self.total < 0 {
		log.Error("Property.Calc ERROR: %f, %f, %f, %f", self.base, self.perc, self.extra, self.total)
		self.total = 0
	}
}

func (self *Property) Value() float64 {
	if self.dirty {
		self.Calc()
	}

	return self.total
}

// ----------------------------------------------------------------------------

func (self *PropertyGroup) AddProp(id uint32, part uint32, val float64) {
	if id <= PropType_Hurt {
		self[id].add(part, val, true)
	}
}

func (self *PropertyGroup) SubProp(id uint32, part uint32, val float64) {
	if id <= PropType_Hurt {
		self[id].add(part, val, false)
	}
}

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

func (self *PropertyGroup) Dump() (ret string) {
	ret = "{\n"
	for i := 0; i < C_Property_Number; i++ {
		ret += fmt.Sprintf("\t{ %s\t= %-20.2f[%-20.2f %-20.2f %-20.2f]},\n", type_name[i], self[i].Value(), self[i].base, self[i].perc, self[i].extra)
	}
	ret += "},\n"
	return
}
