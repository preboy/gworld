package config

import (
	"gworld/core/log"
)

// ============================================================================

type Break struct {
	Id        uint32 `json:"id"`
	Name      string `json:"name"`
	OpenLv    uint32 `json:"openLv"`
	TeamId    uint32 `json:"teamId"`
	DropId    uint32 `json:"dropId"`
	LootId    uint32 `json:"lootId"`
	StartCgid uint32 `json:"startCgid"`
	EndCgid   uint32 `json:"endCgid"`
}

type BreakTable struct {
	items map[uint32]*Break
}

// ============================================================================

var (
	BreakConf = &BreakTable{}
)

// ============================================================================

func (b *BreakTable) Load() bool {
	file := "Break.json"
	var arr []*Break

	if !load_json_as_arr(C_Config_Path+file, &arr) {
		return false
	}

	b.items = make(map[uint32]*Break)
	for _, v := range arr {
		key := v.Id
		b.items[key] = v
	}

	log.Info("load [ %s ] OK", file)
	return true
}

func (b *BreakTable) Query(lv uint32) *Break {
	return b.items[lv]
}

func (b *BreakTable) Items() map[uint32]*Break {
	return b.items
}
