package config

import (
	"core/log"
	"encoding/json"
	"io/ioutil"
)

type SkillProto struct {
	Id           uint32          `json:"id"`
	Level        uint32          `json:"level"`
	Name         string          `json:"name"`
	Prepare_t    int32           `json:"prepare_t"`
	Effect_t     int32           `json:"effect_t"`
	Itv_t        int32           `json:"itv_t"`
	Cd_t         int32           `json:"cd_t"`
	Type         int32           `json:"type"`
	Target_major []int32         `json:"target_major"`
	Target_minor []int32         `json:"target_minor"`
	Aura_major   []*ProbAuraConf `json:"aura_major"`
	Aura_minor   []*ProbAuraConf `json:"aura_minor"`
	Prop_major   []*PropConf     `json:"prop_major"`
	Prop_minor   []*PropConf     `json:"prop_minor"`
	Prop_passive []*PropConf     `json:"prop_passive"`
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

func GetSkillProto(id, lv uint32) *SkillProto {
	if _SkillProtoConf.items == nil {
		return nil
	}

	key := MakeUint64(id, lv)
	return _SkillProtoConf.items[key]
}
