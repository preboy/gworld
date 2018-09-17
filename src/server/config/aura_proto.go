package config

import (
	"core/log"
	"encoding/json"
	"io/ioutil"
)

type AuraProto struct {
	Id       uint32  `json:"id"`
	Level    uint32  `json:"level"`
	Name     string  `json:"name"`
	Last_t   uint32  `json:"last_t"`
	Itv_t    uint32  `json:"itv_t"`
	ScriptId uint32  `json:"script_id"`
	Param1   int32   `json:"param1"`
	Param2   int32   `json:"param2"`
	Param3   []int32 `json:"param3"`
	Param4   string  `json:"param4"`
}

type AuraProtoConf struct {
	items map[uint64]*AuraProto
}

var _AuraProtoConf AuraProtoConf

func GetAuraProtoConf() *AuraProtoConf {
	return &_AuraProtoConf
}

func load_aura() {
	path := "config/AuraProto.json"

	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Error("[AuraProtoConf] loading failed: %s: %s", path, err)
		return
	}

	var arr []*AuraProto
	err = json.Unmarshal(content, &arr)
	if err != nil {
		log.Error("[AuraProtoConf] Unmarshal failed: %s: %s", path, err)
		return
	}

	_AuraProtoConf.items = make(map[uint64]*AuraProto)

	for _, v := range arr {
		key := MakeUint64(v.Id, v.Level)
		_AuraProtoConf.items[key] = v
	}

	log.Info("[AuraProtoConf] load OK")
}

func GetAuraProto(id, lv uint32) *AuraProto {
	if _AuraProtoConf.items == nil {
		return nil
	}

	key := MakeUint64(id, lv)
	return _AuraProtoConf.items[key]
}
