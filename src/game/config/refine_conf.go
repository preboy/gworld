package config

import (
	"core/log"
)

// ============================================================================

type RefineSuper struct {
	Lv     uint32      `json:"lv"`
	Props  []*PropConf `json:"props"`
	Count  uint32      `json:"count"`
	Prob   uint32      `json:"prob"`
	Insure uint32      `json:"insure"`
}

type RefineNormal struct {
	Lv    uint32      `json:"lv"`
	Props []*PropConf `json:"props"`
	Count uint32      `json:"count"`
	Prob  uint32      `json:"prob"`
}

type RefineSuperTable struct {
	items map[uint32]*RefineSuper
}

type RefineNormalTable struct {
	items map[uint32]*RefineNormal
}

// ============================================================================

var (
	RefineSuperConf  = &RefineSuperTable{}
	RefineNormalConf = &RefineNormalTable{}
)

// ============================================================================

func (self *RefineSuperTable) Load() bool {
	file := "RefineSuper.json"
	var arr []*RefineSuper

	if !load_json_as_arr(C_Config_Path+file, &arr) {
		return false
	}

	self.items = make(map[uint32]*RefineSuper)
	for _, v := range arr {
		self.items[v.Lv] = v
	}

	log.Info("load [ %s ] OK", file)
	return true
}

func (self *RefineSuperTable) Query(id uint32) *RefineSuper {
	return self.items[id]
}

func (self *RefineSuperTable) Items() map[uint32]*RefineSuper {
	return self.items
}

// ============================================================================

func (self *RefineNormalTable) Load() bool {
	file := "RefineNormal.json"
	var arr []*RefineNormal

	if !load_json_as_arr(C_Config_Path+file, &arr) {
		return false
	}

	self.items = make(map[uint32]*RefineNormal)
	for _, v := range arr {
		self.items[v.Lv] = v
	}

	log.Info("load [ %s ] OK", file)
	return true
}

func (self *RefineNormalTable) Query(id uint32) *RefineNormal {
	return self.items[id]
}

func (self *RefineNormalTable) Items() map[uint32]*RefineNormal {
	return self.items
}
