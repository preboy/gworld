package config

import (
	"core/log"
	"encoding/json"
	"io/ioutil"
)

type HeroProto struct {
	Id           uint32       `json:"id"`
	Name         string       `json:"name"`
	Level        uint32       `json:"level"`
	Atk          uint32       `json:"atk"`
	Def          uint32       `json:"def"`
	Hp           uint32       `json:"hp"`
	Apm          uint32       `json:"apm"`
	Crit         uint32       `json:"crit"`
	Crit_hurt    uint32       `json:"crit_hurt"`
	Skill_common []*SkillConf `json:"skill_common"`
}

type HeroProtoConf struct {
	items map[uint64]*HeroProto
}

var _HeroProtoConf HeroProtoConf

func GetHeroProtoConf() *HeroProtoConf {
	return &_HeroProtoConf
}

func load_hero() {
	path := "config/HeroProto.json"

	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Error("[HeroProtoConf] loading failed: %s: %s", path, err)
		return
	}

	var arr []*HeroProto
	err = json.Unmarshal(content, &arr)
	if err != nil {
		log.Error("[HeroProtoConf] Unmarshal failed: %s: %s", path, err)
		return
	}

	_HeroProtoConf.items = make(map[uint64]*HeroProto, 0x100)

	for _, v := range arr {
		key := MakeUint64(v.Id, v.Level)
		_HeroProtoConf.items[key] = v
	}

	log.Info("[HeroProtoConf] load OK")
}

func (self *HeroProtoConf) GetHeroProto(id, lv uint32) *HeroProto {
	if self.items == nil {
		return nil
	}

	key := MakeUint64(id, lv)
	return self.items[key]
}
