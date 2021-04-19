package config

import (
	"gworld/core/log"
)

// ============================================================================

type Market struct {
	Index uint32      `json:"index"`
	Class uint32      `json:"class"`
	Src   []*ItemConf `json:"src"`
	Dst   []*ItemConf `json:"dst"`
}

type MarketTable struct {
	items map[uint32]*Market
}

// ============================================================================

var (
	MarketConf = &MarketTable{}
)

// ============================================================================

func (m *MarketTable) Load() bool {
	file := "Market.json"
	var arr []*Market

	if !load_json_as_arr(C_Config_Path+file, &arr) {
		return false
	}

	m.items = make(map[uint32]*Market)
	for _, v := range arr {
		m.items[v.Index] = v
	}

	log.Info("load [ %s ] OK", file)
	return true
}

func (m *MarketTable) Query(index uint32) *Market {
	return m.items[index]
}

func (m *MarketTable) Items() map[uint32]*Market {
	return m.items
}
