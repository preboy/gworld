package config

import (
	"core/log"
)

// ============================================================================

type Aura struct {
	Id        uint32      `json:"id"`
	Level     uint32      `json:"level"`
	Name      string      `json:"name"`
	Last_t    uint32      `json:"last_t"`
	Itv_t     uint32      `json:"itv_t"`
	Deletable int32       `json:"deletable"`
	Helpful   int32       `json:"helpful"`
	ScriptId  uint32      `json:"script_id"`
	Props     []*PropConf `json:"props"`
	Params    []int32     `json:"params"`
}

type AuraTable struct {
	items map[uint64]*Aura
}

// ============================================================================

var (
	AuraProtoConf = &AuraTable{}
)

// ============================================================================

func (self *AuraTable) Load() bool {
	file := "Aura.json"
	var arr []*Aura

	if !load_json_as_arr(C_Config_Path+file, &arr) {
		return false
	}

	self.items = make(map[uint64]*Aura)
	for _, v := range arr {
		key := MakeUint64(v.Id, v.Level)
		self.items[key] = v
	}

	log.Info("load [ %s ] OK", file)
	return true
}

func (self *AuraTable) Query(id, lv uint32) *Aura {
	key := MakeUint64(id, lv)
	return self.items[key]
}

func (self *AuraTable) Items() map[uint64]*Aura {
	return self.items
}
