package config

import (
	"gworld/core/log"
)

// ============================================================================

type Skill struct {
	Id           uint32          `json:"id"`
	Lv           uint32          `json:"lv"`
	Name         string          `json:"name"`
	Prepare_t    uint32          `json:"prepare_t"`
	Effect_t     uint32          `json:"effect_t"`
	Itv_t        uint32          `json:"itv_t"`
	Cd_t         uint32          `json:"cd_t"`
	Type         int32           `json:"type"`
	Target_major []int32         `json:"target_major"`
	Ratio_major  float64         `json:"ratio_major"`
	Extra_major  []*ExtraAtkConf `json:"extra_major"`
	Aura_major   []*ProbAuraConf `json:"aura_major"`
	Target_minor []int32         `json:"target_minor"`
	Ratio_minor  float64         `json:"ratio_minor"`
	Extra_minor  []*ExtraAtkConf `json:"extra_minor"`
	Aura_minor   []*ProbAuraConf `json:"aura_minor"`
	Prop_Passive []*PropConf     `json:"prop_passive"`
	Aura_Passive []*ProbAuraConf `json:"aura_passive"`
}

type SkillTable struct {
	items map[uint64]*Skill
}

// ============================================================================

var (
	SkillProtoConf = &SkillTable{}
)

// ============================================================================

func (s *SkillTable) Load() bool {
	file := "Skill.json"
	var arr []*Skill

	if !load_json_as_arr(C_Config_Path+file, &arr) {
		return false
	}

	s.items = make(map[uint64]*Skill)
	for _, v := range arr {
		key := MakeUint64(v.Id, v.Lv)
		s.items[key] = v
	}

	log.Info("load [ %s ] OK", file)
	return true
}

func (s *SkillTable) Query(id, lv uint32) *Skill {
	key := MakeUint64(id, lv)
	return s.items[key]

}

func (s *SkillTable) Items() map[uint64]*Skill {
	return s.items
}
