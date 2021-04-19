package config

import (
	"gworld/core/log"
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

func (i *ItemTable) Load() bool {
	file := "Item.json"
	var arr []*Item

	if !load_json_as_arr(C_Config_Path+file, &arr) {
		return false
	}

	i.items = make(map[uint32]*Item)
	for _, v := range arr {
		i.items[v.Id] = v
	}

	log.Info("load [ %s ] OK", file)
	return true
}

func (i *ItemTable) Query(id uint32) *Item {
	return i.items[id]
}

func (i *ItemTable) Items() map[uint32]*Item {
	return i.items
}
