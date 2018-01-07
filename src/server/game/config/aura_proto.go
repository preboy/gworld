package config

import (
	"core/log"
	"encoding/json"
	"io/ioutil"
)

type AuraProto struct {
	Id       uint32 `json:"id"`
	Level    uint32 `json:"level"`
	Name     string `json:"name"`
	Duration uint32 `json:"duration"`
	Uptime   uint32 `json:"uptime"`
	Sid      uint32 `json:"sid"`
	Param1   uint32 `json:"param1"`
	Param2   uint32 `json:"param2"`
	Param3   uint32 `json:"param3"`
	Param4   uint32 `json:"param4"`
}

type AuraProtoConf struct {
	items map[uint32]*AuraProto
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

	_AuraProtoConf.items = make(map[uint32]*AuraProto)

	for _, v := range arr {
		_AuraProtoConf.items[v.Id] = v
	}

	log.Info("[AuraProtoConf] load OK")
}

func (self *AuraProtoConf) GetAuraProto(id uint32) *AuraProto {
	if self.items == nil {
		return nil
	}

	return self.items[id]
}
