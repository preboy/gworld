package config

import (
	"core/log"
	"encoding/json"
	"io/ioutil"
)

type RefineSuper struct {
	Level  uint32 `json:"level"`
	Apm    uint32 `json:"apm"`
	Atk    uint32 `json:"atk"`
	Def    uint32 `json:"def"`
	Hp     uint32 `json:"hp"`
	Crit   uint32 `json:"crit"`
	Hurt   uint32 `json:"crit_hurt"`
	Count  uint32 `json:"count"`
	Prob   uint32 `json:"prob"`
	Insure uint32 `json:"insure"`
}

type RefineNormal struct {
	Level uint32 `json:"level"`
	Apm   uint32 `json:"apm"`
	Atk   uint32 `json:"atk"`
	Def   uint32 `json:"def"`
	Hp    uint32 `json:"hp"`
	Crit  uint32 `json:"crit"`
	Hurt  uint32 `json:"crit_hurt"`
	Count uint32 `json:"count"`
	Prob  uint32 `json:"prob"`
}

type RefineSuperConf struct {
	items []*RefineSuper
}

type RefineNormalConf struct {
	items []*RefineNormal
}

var _RefineSuperConf RefineSuperConf
var _RefineNormalConf RefineNormalConf

func GetRefineSuperConf() *RefineSuperConf {
	return &_RefineSuperConf
}

func GetRefineNormalConf() *RefineNormalConf {
	return &_RefineNormalConf
}

func load_refine_super() {
	path := "config/RefineSuper.json"
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Error("[RefineSuper] loading failed: %s: %s", path, err)
		return
	}

	var arr []*RefineSuper
	err = json.Unmarshal(content, &arr)
	if err != nil {
		log.Error("[RefineSuper] Unmarshal failed: %s: %s", path, err)
		return
	}

	_RefineSuperConf.items = make([]*RefineSuper, len(arr)+1)
	for _, v := range arr {
		_RefineSuperConf.items[v.Level] = v
	}

	log.Info("[RefineSuper] load OK")
}

func load_refine_normal() {
	path := "config/RefineNormal.json"
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Error("[RefineNormal] loading failed: %s: %s", path, err)
		return
	}

	var arr []*RefineNormal
	err = json.Unmarshal(content, &arr)
	if err != nil {
		log.Error("[RefineNormal] Unmarshal failed: %s: %s", path, err)
		return
	}

	_RefineNormalConf.items = make([]*RefineNormal, len(arr)+1)
	for _, v := range arr {
		_RefineNormalConf.items[v.Level] = v
	}

	log.Info("[RefineNormal] load OK")
}

func GetRefineSuper(lv uint32) *RefineSuper {
	if _RefineSuperConf.items == nil {
		return nil
	}

	if lv < 1 || int(lv) >= len(_RefineSuperConf.items) {
		return nil
	}

	return _RefineSuperConf.items[lv]
}

func GetRefineNormal(lv uint32) *RefineNormal {
	if _RefineNormalConf.items == nil {
		return nil
	}

	if lv < 1 || int(lv) >= len(_RefineNormalConf.items) {
		return nil
	}

	return _RefineNormalConf.items[lv]
}
