package config

import (
	"gworld/core/log"
)

// ============================================================================

type HeroTalent struct {
	Id    uint32      `json:"id"`
	Lv    uint32      `json:"lv"`
	Props []*PropConf `json:"props"`
}

type HeroTalentTable struct {
	items map[uint64]*HeroTalent
}

// ============================================================================

var (
	HeroTalentConf = &HeroTalentTable{}
)

// ============================================================================

func (h *HeroTalentTable) Load() bool {
	file := "HeroTalent.json"
	var arr []*HeroTalent

	if !load_json_as_arr(C_Config_Path+file, &arr) {
		return false
	}

	h.items = make(map[uint64]*HeroTalent)
	for _, v := range arr {
		key := MakeUint64(v.Id, v.Lv)
		h.items[key] = v
	}

	log.Info("load [ %s ] OK", file)
	return true
}

func (h *HeroTalentTable) Query(id, lv uint32) *HeroTalent {
	key := MakeUint64(id, lv)
	return h.items[key]
}
