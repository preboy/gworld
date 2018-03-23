package config

import (
	"core/log"
	"encoding/json"
	"io/ioutil"
)

type SkillProto struct {
	Id      uint32      `json:"id"`
	Level   uint32      `json:"level"`
	Name    string      `json:"name"`
	Passive int32       `json:"passive"`
	Target  uint32      `json:""`
	Action  uint32      `json:"acttargetion"`
	Last_t  int32       `json:"last_t"`
	Itv_t   int32       `json:"itv_t"`
	Cd_t    int32       `json:"cd_t"`
	Auras   []*AuraConf `json:"aura"`
	Attrs   []*AttrConf `json:"attr"`
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
		key := MakeUint64(v.Id, v.Level)
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
