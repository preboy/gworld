package config

import (
	"core/log"
)

// ============================================================================

type Cond struct {
	Id     uint32  `json:"condId"`
	Type   uint32  `json:"condType"`
	Params []int32 `json:"params"`
}

type CondTable struct {
	items map[uint32]*Cond
}

// ============================================================================

var (
	CondConf = &CondTable{}
)

// ============================================================================

func (self *CondTable) Load() bool {
	file := "Cond.json"
	var arr []*Cond

	if !load_json_as_arr(C_Config_Path+file, &arr) {
		return false
	}

	self.items = make(map[uint32]*Cond)
	for _, v := range arr {
		key := v.Id
		self.items[key] = v
	}

	log.Info("load [ %s ] OK", file)
	return true
}

func (self *CondTable) Query(id uint32) *Cond {
	return self.items[id]
}

func (self *CondTable) Items() map[uint32]*Cond {
	return self.items
}
