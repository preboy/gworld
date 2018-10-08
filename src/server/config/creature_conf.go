package config

import (
	"core/log"
)

// ============================================================================

type Creature struct {
	Id          uint32       `json:"id"`
	Level       uint32       `json:"level"`
	Name        string       `json:"name"`
	Apm         uint32       `json:"apm"`
	Atk         uint32       `json:"atk"`
	Def         uint32       `json:"def"`
	Hp          uint32       `json:"hp"`
	Crit        uint32       `json:"crit"`
	Hurt        uint32       `json:"crit_hurt"`
	SkillExtra  []*SkillConf `json:"skill_extra"`
	SkillCommon []*SkillConf `json:"skill_common"`
	Auras       []*AuraConf  `json:"aura"`
}

type CreatureTable struct {
	items map[uint64]*Creature
}

// ============================================================================

var (
	CreatureConf = &CreatureTable{}
)

// ============================================================================

func (self *CreatureTable) Load() bool {
	file := "Creature.json"
	var arr []*Creature

	if !load_json_as_arr(C_Config_Path+file, &arr) {
		return false
	}

	self.items = make(map[uint64]*Creature)
	for _, v := range arr {
		key := MakeUint64(v.Id, v.Level)
		self.items[key] = v
	}

	log.Info("load [ %s ] OK", file)
	return true
}

func (self *CreatureTable) Query(id, lv uint32) *Creature {
	key := MakeUint64(id, lv)
	return self.items[key]
}

func (self *CreatureTable) Items() map[uint64]*Creature {
	return self.items
}
