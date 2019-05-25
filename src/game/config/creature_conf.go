package config

import (
	"core/log"
)

// ============================================================================

type Creature struct {
	Id          uint32       `json:"id"`
	Name        string       `json:"name"`
	Props       []*PropConf  `json:"props"`
	SkillExtra  []*SkillConf `json:"skill_extra"`
	SkillCommon []*SkillConf `json:"skill_common"`
	Auras       []*AuraConf  `json:"aura"`
}

type CreatureTable struct {
	items map[uint32]*Creature
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

	self.items = make(map[uint32]*Creature)
	for _, v := range arr {
		self.items[v.Id] = v
	}

	log.Info("load [ %s ] OK", file)
	return true
}

func (self *CreatureTable) Query(id uint32) *Creature {
	return self.items[id]
}

func (self *CreatureTable) Items() map[uint32]*Creature {
	return self.items
}
