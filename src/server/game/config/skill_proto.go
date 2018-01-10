package config

import (
	"core/log"
	"encoding/json"
	"io/ioutil"
)

type AuraInfo struct {
	Rate uint32 `json:"rate"`
	Id   uint32 `json:"id"`
	Lv   uint32 `json:"lv"`
}

type AddAttr struct {
	Id  uint32 `json:"id"`
	Val uint32 `json:"val"`
}

type SkillProto struct {
	Id     uint32      `json:"id"`
	Lv     uint32      `json:"lv"`
	Name   string      `json:"name"`
	Target uint32      `json:"target"`
	Itv_t  uint32      `json:"itv_t"`
	Last_t uint32      `json:"last_t"`
	Cd_t   uint32      `json:"cd_t"`
	Type   uint32      `json:"type"`
	Range  uint32      `json:"range"`
	Auras  []*AuraInfo `json:"auras"`
	Attrs  []*AddAttr  `json:"add_attr"`
}

type SkillProtoConf struct {
	items map[uint64]*SkillProto
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

	_SkillProtoConf.items = make(map[uint64]*SkillProto)

	for _, v := range arr {
		key := MakeUint64(v.Id, v.Lv)
		_SkillProtoConf.items[key] = v
	}

	log.Info("[SkillProtoConf] load OK")
}

func (self *SkillProtoConf) GetSkillProto(id, lv uint32) *SkillProto {
	if self.items == nil {
		return nil
	}

	key := MakeUint64(id, lv)
	return self.items[key]
}
