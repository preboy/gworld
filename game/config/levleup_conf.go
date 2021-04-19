package config

import (
	"gworld/core/log"
)

// ============================================================================

type Levelup struct {
	Lv    uint32      `json:"lv"`
	Exp   uint64      `json:"exp"`
	Needs []*ItemConf `json:"needs"`
}

type LevelupTable struct {
	items []*Levelup
}

// ============================================================================

var (
	LevelupConf = &LevelupTable{}
)

// ============================================================================

func (l *LevelupTable) Load() bool {
	file := "Levelup.json"
	var arr []*Levelup

	if !load_json_as_arr(C_Config_Path+file, &arr) {
		return false
	}

	l.items = make([]*Levelup, len(arr))
	for _, v := range arr {
		l.items[v.Lv-1] = v
	}

	log.Info("load [ %s ] OK", file)
	return true
}

func (l *LevelupTable) Query(lv uint32) *Levelup {
	if lv < 1 || int(lv) > len(l.items) {
		return nil
	}

	return l.items[lv-1]
}

func (l *LevelupTable) Items() []*Levelup {
	return l.items
}
