package config

import (
	"core/log"
)

// ============================================================================

type RefineSuper struct {
	Level  uint32 `json:"level"`
	Apm    uint32 `json:"apm"`
	Atk    uint32 `json:"atk"`
	Def    uint32 `json:"def"`
	Hp     uint32 `json:"hp"`
	Crit   uint32 `json:"crit"`
	Hurt   uint32 `json:"crit_hurt"`
	Count  uint32 `json:"count"`
	Prob   uint32 `json:"prob"`
	Insure uint32 `json:"insure"`
}

type RefineNormal struct {
	Level uint32 `json:"level"`
	Apm   uint32 `json:"apm"`
	Atk   uint32 `json:"atk"`
	Def   uint32 `json:"def"`
	Hp    uint32 `json:"hp"`
	Crit  uint32 `json:"crit"`
	Hurt  uint32 `json:"crit_hurt"`
	Count uint32 `json:"count"`
	Prob  uint32 `json:"prob"`
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

	if !load_json_as_arr(file, &arr) {
		return false
	}

	self.items = make(map[uint32]*RefineSuper)
	for _, v := range arr {
		self.items[v.Level] = v
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

	if !load_json_as_arr(file, &arr) {
		return false
	}

	self.items = make(map[uint32]*RefineNormal)
	for _, v := range arr {
		self.items[v.Level] = v
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
