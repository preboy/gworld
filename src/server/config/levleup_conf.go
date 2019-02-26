package config

import (
	"core/log"
)

// ============================================================================

type Levelup struct {
	Lv  uint32 `json:"lv"`
	Exp uint64 `json:"exp"`
}

type LevelupTable struct {
	items map[uint32]*Levelup
}

// ============================================================================

var (
	LevelupConf = &LevelupTable{}
)

// ============================================================================

func (self *LevelupTable) Load() bool {
	file := "Levelup.json"
	var arr []*Levelup

	if !load_json_as_arr(C_Config_Path+file, &arr) {
		return false
	}

	self.items = make(map[uint32]*Levelup)
	for _, v := range arr {
		key := v.Lv
		self.items[key] = v
	}

	log.Info("load [ %s ] OK", file)
	return true
}

func (self *LevelupTable) Query(lv uint32) *Levelup {
	return self.items[lv]
}

func (self *LevelupTable) Items() map[uint32]*Levelup {
	return self.items
}
