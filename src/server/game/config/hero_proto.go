package config

import (
	"core/log"
	"encoding/json"
	"io/ioutil"
)

type HeroProto struct {
	Id        uint32 `json:"id"`
	Name      string `json:"name"`
	Level     uint32 `json:"level"`
	Atk       uint32 `json:"atk"`
	Def       uint32 `json:"def"`
	Hp        uint32 `json:"hp"`
	Apm       uint32 `json:"apm"`
	Crit      uint32 `json:"crit"`
	Crit_hurt uint32 `json:"crit_hurt"`
}

type HeroProtoConf struct {
	items map[uint32]*HeroProto
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

	_HeroProtoConf.items = make(map[uint32]*HeroProto)

	for _, v := range arr {
		_HeroProtoConf.items[v.Id] = v
	}

	log.Info("[HeroProtoConf] load OK")
}

func (self *HeroProtoConf) GetHeroProto(id, lv uint32) *HeroProto {
	if self.items == nil {
		return nil
	}

	return self.items[id]
}
