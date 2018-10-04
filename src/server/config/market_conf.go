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

func (self *MarketTable) Load() bool {
	file := "Market.json"
	var arr []*Market

	if !load_json_as_arr(file, &arr) {
		return false
	}

	self.items = make(map[uint32]*Market)
	for _, v := range arr {
		self.items[v.Index] = v
	}

	log.Info("load [ %s ] OK", file)
	return true
}

func (self *MarketTable) Query(index uint32) *Market {
	return self.items[index]
}

func (self *MarketTable) Items() map[uint32]*Market {
	return self.items
}
