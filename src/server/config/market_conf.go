package config

import (
	"core/log"
	"encoding/json"
	"io/ioutil"
)

type MarketItem struct {
	Index uint32      `json:"index"`
	Class uint32      `json:"class"`
	Src   []*ItemConf `json:"src"`
	Dst   []*ItemConf `json:"dst"`
}

type MarketConf struct {
	items map[uint32]*MarketItem
}

var _MarketConf MarketConf

func GetMarketConf() *MarketConf {
	return &_MarketConf
}

func load_market_conf() {
	path := "config/MarketConf.json"

	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Error("[MarketConf] loading failed: %s: %s", path, err)
		return
	}

	var arr []*MarketItem
	err = json.Unmarshal(content, &arr)
	if err != nil {
		log.Error("[MarketConf] Unmarshal failed: %s: %s", path, err)
		return
	}

	_MarketConf.items = make(map[uint32]*MarketItem)

	for _, v := range arr {
		_MarketConf.items[v.Index] = v
	}

	log.Info("[MarketConf] load OK")
}

func GetMarketItem(index uint32) *MarketItem {
	if _MarketConf.items == nil {
		return nil
	}
	return _MarketConf.items[index]
}
