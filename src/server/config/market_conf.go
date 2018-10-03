package config

import (
	"core/log"
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

func (self *MarketTable) Load() {
	file := "Market.json"
	var arr []*Market

	if !load_from_json(file, arr) {
		return
	}

	self.items = make(map[uint32]*Market)
	for _, v := range arr {
		self.items[v.Index] = v
	}

	log.Info("[%s] load OK", file)
}

func (self *MarketTable) Query(index uint32) *Market {
	return self.items[index]
}

func (self *MarketTable) Items() map[uint32]*Market {
	return self.items
}
