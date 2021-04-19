package config

import (
	"gworld/core/log"
)

// ============================================================================

type Hero struct {
	Id          uint32       `json:"id"`
	Name        string       `json:"name"`
	Apm         uint32       `json:"apm"`
	Talent      uint32       `json:"talent"`
	Aptitude    uint32       `json:"aptitude"`
	SkillCommon []*SkillConf `json:"skill_common"`
	Skill1      uint32       `json:"skill_1"`
	Skill2      uint32       `json:"skill_2"`
}

type HeroTable struct {
	items map[uint32]*Hero
}

// ============================================================================

var (
	HeroConf = &HeroTable{}
)

// ============================================================================

func (h *HeroTable) Load() bool {
	file := "Hero.json"
	var arr []*Hero

	if !load_json_as_arr(C_Config_Path+file, &arr) {
		return false
	}

	h.items = make(map[uint32]*Hero)
	for _, v := range arr {
		h.items[v.Id] = v
	}

	log.Info("load [ %s ] OK", file)
	return true
}

func (h *HeroTable) Query(id uint32) *Hero {
	return h.items[id]
}

func (h *HeroTable) Items() map[uint32]*Hero {
	return h.items
}
