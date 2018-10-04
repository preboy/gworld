package config

import (
	"core/log"
)

// ============================================================================

type Global struct {
	Name  string      `json:"name"`
	GID   uint32      `json:"gID"`
	Rate  float64     `json:"rate"`
	Enc   []*ItemConf `json:"enc"`
	Level []uint32    `json:"level"`
}

type GlobalTable struct {
	item *Global
}

// ============================================================================

var (
	GlobalConf = &GlobalTable{}
)

// ============================================================================

func (self *GlobalTable) Load() bool {
	file := "Global.json"
	var arr []*Global

	if !load_json_as_arr(file, &arr) {
		return false
	}

	if len(arr) == 0 {
		return false
	}

	self.item = arr[0]

	log.Info("load [ %s ] OK", file)
	return true
}

func (self *GlobalTable) Get() *Global {
	return self.item
}
