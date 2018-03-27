package config

import (
	"core/log"
	"encoding/json"
	"io/ioutil"
)

type ItemProto struct {
	Id       uint32  `json:"id"`
	Name     string  `json:"name"`
	Qulity   uint32  `json:"qulity"`
	Type_p   uint32  `json:"type_p"`
	Type_s   uint32  `json:"type_s"`
	Usable   uint32  `json:"usable"`
	ScriptID uint32  `json:"script_id"`
	Param1   int32   `json:"param1"`
	Param2   int32   `json:"param2"`
	Param3   []int32 `json:"param3"`
	Param4   string  `json:"param4"`
}

type ItemProtoConf struct {
	items map[uint32]*ItemProto
}

var _ItemProtoConf ItemProtoConf

func GetItemProtoConf() *ItemProtoConf {
	return &_ItemProtoConf
}

func load_item() {
	path := "config/ItemProto.json"

	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Error("[ItemProtoConf] loading failed: %s: %s", path, err)
		return
	}

	var arr []*ItemProto
	err = json.Unmarshal(content, &arr)
	if err != nil {
		log.Error("[ItemProtoConf] Unmarshal failed: %s: %s", path, err)
		return
	}

	_ItemProtoConf.items = make(map[uint32]*ItemProto)

	for _, v := range arr {
		_ItemProtoConf.items[v.Id] = v
	}

	log.Info("[ItemProtoConf] load OK")
}

func (self *ItemProtoConf) ItemProto(id uint32) *ItemProto {
	if self.items == nil {
		return nil
	}
	return self.items[id]
}
