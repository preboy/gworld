package config

import (
	"gworld/core/log"
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

func (g *GlobalTable) Load() bool {
	file := "Global.json"
	var arr []*Global

	if !load_json_as_arr(C_Config_Path+file, &arr) {
		return false
	}

	if len(arr) == 0 {
		return false
	}

	g.item = arr[0]

	log.Info("load [ %s ] OK", file)
	return true
}

func (g *GlobalTable) Get() *Global {
	return g.item
}
