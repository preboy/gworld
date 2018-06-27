package config

import (
	"core/log"
	"encoding/json"
	"io/ioutil"
)

type CreatureProto struct {
	Id          uint32       `json:"id"`
	Level       uint32       `json:"level"`
	Name        string       `json:"name"`
	Apm         uint32       `json:"apm"`
	Atk         uint32       `json:"atk"`
	Def         uint32       `json:"def"`
	Hp          uint32       `json:"hp"`
	Crit        uint32       `json:"crit"`
	Hurt        uint32       `json:"crit_hurt"`
	SkillExtra  []*SkillConf `json:"skill_extra"`
	SkillCommon []*SkillConf `json:"skill_common"`
	Auras       []*AuraConf  `json:"aura"`
}

type CreatureProtoConf struct {
	items map[uint64]*CreatureProto
}

var _CreatureProtoConf CreatureProtoConf

func GetCreatureProtoConf() *CreatureProtoConf {
	return &_CreatureProtoConf
}

func load_creature() {
	path := "config/CreatureProto.json"

	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Error("[CreatureProtoConf] loading failed: %s: %s", path, err)
		return
	}

	var arr []*CreatureProto
	err = json.Unmarshal(content, &arr)
	if err != nil {
		log.Error("[CreatureProtoConf] Unmarshal failed: %s: %s", path, err)
		return
	}

	_CreatureProtoConf.items = make(map[uint64]*CreatureProto, 0x400)

	for _, v := range arr {
		key := MakeUint64(v.Id, v.Level)
		_CreatureProtoConf.items[key] = v
	}

	log.Info("[CreatureProtoConf] load OK")
}

func GetCreatureProto(id, lv uint32) *CreatureProto {
	if _CreatureProtoConf.items == nil {
		return nil
	}

	key := MakeUint64(id, lv)
	return _CreatureProtoConf.items[key]
}
