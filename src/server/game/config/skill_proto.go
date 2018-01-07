package config

import (
	"core/log"
	"encoding/json"
	"io/ioutil"
)

type AuraInfo struct {
	Rate uint32 `json:"rate"`
	Id   uint32 `json:"id"`
}

type AddAttr struct {
	Id  uint32 `json:"id"`
	Val uint32 `json:"val"`
}

type SkillProto struct {
	Id        uint32      `json:"id"`
	Level     uint32      `json:"level"`
	Name      string      `json:"name"`
	Object    uint32      `json:"object"`
	Cast_time uint32      `json:"cast_time"`
	Cd_time   uint32      `json:"cd_time"`
	Type      uint32      `json:"type"`
	Range     uint32      `json:"range"`
	Auras     []*AuraInfo `json:"auras"`
	Attrs     []*AddAttr  `json:"add_attr"`
}

type SkillProtoConf struct {
	items map[uint32]*SkillProto
}

var _SkillProtoConf SkillProtoConf

func GetSkillProtoConf() *SkillProtoConf {
	return &_SkillProtoConf
}

func load_skill() {
	path := "config/SkillProto.json"

	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Error("[SkillProtoConf] loading failed: %s: %s", path, err)
		return
	}

	var arr []*SkillProto
	err = json.Unmarshal(content, &arr)
	if err != nil {
		log.Error("[SkillProtoConf] Unmarshal failed: %s: %s", path, err)
		return
	}

	_SkillProtoConf.items = make(map[uint32]*SkillProto)

	for _, v := range arr {
		_SkillProtoConf.items[v.Id] = v
	}

	log.Info("[SkillProtoConf] load OK")
}

func (self *SkillProtoConf) GetSkillProto(id uint32) *SkillProto {
	if self.items == nil {
		return nil
	}

	return self.items[id]
}
