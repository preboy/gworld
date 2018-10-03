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

func (self *GlobalTable) Load() {
	file := "Global.json"
	var arr []*Global

	if !load_from_json(file, arr) {
		return
	}

	self.item = arr[0]

	log.Info("[%s] load OK", file)
}

func (self *GlobalTable) Get() *Global {
	return self.item
}
