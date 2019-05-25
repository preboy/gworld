package config

import (
	"core/log"
)

// ============================================================================

type HeroProp struct {
	Id    uint32      `json:"id"`
	Lv    uint32      `json:"level"`
	Props []*PropConf `json:"props"`
}

type HeroPropTable struct {
	items map[uint64]*HeroProp
}

// ============================================================================

var (
	HeroPropConf = &HeroPropTable{}
)

// ============================================================================

func (self *HeroPropTable) Load() bool {
	file := "HeroProp.json"
	var arr []*HeroProp

	if !load_json_as_arr(C_Config_Path+file, &arr) {
		return false
	}

	self.items = make(map[uint64]*HeroProp)
	for _, v := range arr {
		key := MakeUint64(v.Id, v.Lv)
		self.items[key] = v
	}

	log.Info("load [ %s ] OK", file)
	return true
}

func (self *HeroPropTable) Query(id, lv uint32) *HeroProp {
	key := MakeUint64(id, lv)
	return self.items[key]
}
