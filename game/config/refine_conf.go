package config

import (
	"gworld/core/log"
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

func (r *RefineSuperTable) Load() bool {
	file := "RefineSuper.json"
	var arr []*RefineSuper

	if !load_json_as_arr(C_Config_Path+file, &arr) {
		return false
	}

	r.items = make(map[uint32]*RefineSuper)
	for _, v := range arr {
		r.items[v.Lv] = v
	}

	log.Info("load [ %s ] OK", file)
	return true
}

func (r *RefineSuperTable) Query(id uint32) *RefineSuper {
	return r.items[id]
}

func (r *RefineSuperTable) Items() map[uint32]*RefineSuper {
	return r.items
}

// ============================================================================

func (r *RefineNormalTable) Load() bool {
	file := "RefineNormal.json"
	var arr []*RefineNormal

	if !load_json_as_arr(C_Config_Path+file, &arr) {
		return false
	}

	r.items = make(map[uint32]*RefineNormal)
	for _, v := range arr {
		r.items[v.Lv] = v
	}

	log.Info("load [ %s ] OK", file)
	return true
}

func (r *RefineNormalTable) Query(id uint32) *RefineNormal {
	return r.items[id]
}

func (r *RefineNormalTable) Items() map[uint32]*RefineNormal {
	return r.items
}
