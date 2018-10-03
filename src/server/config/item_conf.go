package config

import (
	"core/log"
)

// ============================================================================

type Item struct {
	Id      uint32  `json:"id"`
	Name    string  `json:"name"`
	Qulity  uint32  `json:"qulity"`
	Type_p  uint32  `json:"type_p"`
	Type_s  uint32  `json:"type_s"`
	UseType uint32  `json:"use_type"`
	Param1  int32   `json:"param1"`
	Param2  int32   `json:"param2"`
	Param3  []int32 `json:"param3"`
	Param4  string  `json:"param4"`
}

type ItemTable struct {
	items map[uint32]*Item
}

// ============================================================================

var (
	ItemProtoConf = &ItemTable{}
)

// ============================================================================

func (self *ItemTable) Load() {
	file := "Item.json"
	var arr []*Item

	if !load_from_json(file, arr) {
		return
	}

	self.items = make(map[uint32]*Item)
	for _, v := range arr {
		self.items[v.Id] = v
	}

	log.Info("[%s] load OK", file)
}

func (self *ItemTable) Query(id uint32) *Item {
	return self.items[id]
}

func (self *ItemTable) Items() map[uint32]*Item {
	return self.items
}
