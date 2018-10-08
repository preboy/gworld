package config

import (
	"core/log"
)

// ============================================================================

type Hero struct {
	Id          uint32       `json:"id"`
	Level       uint32       `json:"level"`
	Name        string       `json:"name"`
	Apm         uint32       `json:"apm"`
	Atk         uint32       `json:"atk"`
	Def         uint32       `json:"def"`
	Hp          uint32       `json:"hp"`
	Crit        uint32       `json:"crit"`
	Hurt        uint32       `json:"crit_hurt"`
	SkillCommon []*SkillConf `json:"skill_common"`
	Needs       []*ItemConf  `json:"needs"`
	Skill1      uint32       `json:"skill_1"`
	Skill2      uint32       `json:"skill_2"`
}

type HeroTable struct {
	items map[uint64]*Hero
}

// ============================================================================

var (
	HeroConf = &HeroTable{}
)

// ============================================================================

func (self *HeroTable) Load() bool {
	file := "Hero.json"
	var arr []*Hero

	if !load_json_as_arr(C_Config_Path+file, &arr) {
		return false
	}

	self.items = make(map[uint64]*Hero)
	for _, v := range arr {
		key := MakeUint64(v.Id, v.Level)
		self.items[key] = v
	}

	log.Info("load [ %s ] OK", file)
	return true
}

func (self *HeroTable) Query(id, lv uint32) *Hero {
	key := MakeUint64(id, lv)
	return self.items[key]
}

func (self *HeroTable) Items() map[uint64]*Hero {
	return self.items
}
