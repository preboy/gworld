package config

import (
	"core/log"
	"encoding/json"
	"io/ioutil"
)

type LevelupExp struct {
	Level uint32 `json:"level"`
	Exp   uint32 `json:"exp"`
}

type LevelupExpConf struct {
	items []*LevelupExp
}

var _LevelupExpConf LevelupExpConf

func GetLevelupExpConf() *LevelupExpConf {
	return &_LevelupExpConf
}

func load_levelup_exp() {
	path := "config/LevelupExp.json"
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Error("[LevelupExpConf] loading failed: %s: %s", path, err)
		return
	}

	var arr []*LevelupExp
	err = json.Unmarshal(content, &arr)
	if err != nil {
		log.Error("[LevelupExpConf] Unmarshal failed: %s: %s", path, err)
		return
	}

	_LevelupExpConf.items = make([]*LevelupExp, len(arr)+1)

	for _, v := range arr {
		_LevelupExpConf.items[v.Level] = v
	}

	log.Info("[LevelupExpConf] load OK")
}

func GetLevelupExp(lv uint32) *LevelupExp {
	if _LevelupExpConf.items == nil {
		return nil
	}
	if lv < 1 || int(lv) >= len(_LevelupExpConf.items) {
		return nil
	}
	return _LevelupExpConf.items[lv]
}
