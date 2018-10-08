package config

import (
	"core/log"
)

// ============================================================================

type Skill struct {
	Id           uint32          `json:"id"`
	Level        uint32          `json:"level"`
	Name         string          `json:"name"`
	Prepare_t    uint32          `json:"prepare_t"`
	Effect_t     uint32          `json:"effect_t"`
	Itv_t        uint32          `json:"itv_t"`
	Cd_t         uint32          `json:"cd_t"`
	Type         int32           `json:"type"`
	Target_major []int32         `json:"target_major"`
	Target_minor []int32         `json:"target_minor"`
	Aura_major   []*ProbAuraConf `json:"aura_major"`
	Aura_minor   []*ProbAuraConf `json:"aura_minor"`
	Prop_major   []*PropConf     `json:"prop_major"`
	Prop_minor   []*PropConf     `json:"prop_minor"`
	Prop_passive []*PropConf     `json:"prop_passive"`
}

type SkillTable struct {
	items map[uint64]*Skill
}

// ============================================================================

var (
	SkillProtoConf = &SkillTable{}
)

// ============================================================================

func (self *SkillTable) Load() bool {
	file := "Skill.json"
	var arr []*Skill

	if !load_json_as_arr(C_Config_Path+file, &arr) {
		return false
	}

	self.items = make(map[uint64]*Skill)
	for _, v := range arr {
		key := MakeUint64(v.Id, v.Level)
		self.items[key] = v
	}

	log.Info("load [ %s ] OK", file)
	return true
}

func (self *SkillTable) Query(id, lv uint32) *Skill {
	key := MakeUint64(id, lv)
	return self.items[key]

}

func (self *SkillTable) Items() map[uint64]*Skill {
	return self.items
}
