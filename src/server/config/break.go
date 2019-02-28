package config

import (
	"core/log"
)

// ============================================================================

type Break struct {
	BreakId   uint32 `json:"breakId"`
	Name      string `json:"name"`
	OpenLv    uint32 `json:"openLv"`
	TeamId    uint32 `json:"teamId"`
	DropId    uint32 `json:"dropId"`
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

func (self *BreakTable) Load() bool {
	file := "Break.json"
	var arr []*Break

	if !load_json_as_arr(C_Config_Path+file, &arr) {
		return false
	}

	self.items = make(map[uint32]*Break)
	for _, v := range arr {
		key := v.BreakId
		self.items[key] = v
	}

	log.Info("load [ %s ] OK", file)
	return true
}

func (self *BreakTable) Query(lv uint32) *Break {
	return self.items[lv]
}

func (self *BreakTable) Items() map[uint32]*Break {
	return self.items
}
