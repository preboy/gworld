package config

import (
	"core/log"
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

func (self *LevelupTable) Load() bool {
	file := "Levelup.json"
	var arr []*Levelup

	if !load_json_as_arr(C_Config_Path+file, &arr) {
		return false
	}

	self.items = make([]*Levelup, len(arr))
	for _, v := range arr {
		self.items[v.Lv-1] = v
	}

	log.Info("load [ %s ] OK", file)
	return true
}

func (self *LevelupTable) Query(lv uint32) *Levelup {
	if lv < 1 || int(lv) > len(self.items) {
		return nil
	}

	return self.items[lv-1]
}

func (self *LevelupTable) Items() []*Levelup {
	return self.items
}
