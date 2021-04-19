package config

import (
	"gworld/core/log"
)

// ============================================================================

type Achv struct {
	Id  int32 `json:"id"`
	Gid int32 `json:"gid"`
	Val int32 `json:"val"`
}

type AchvTable struct {
	items map[int32]*Achv
}

type Growth struct {
	Id int32 `json:"id"`
}

type GrowthTable struct {
	items map[int32]*Growth
}

// ============================================================================

var (
	AchvConf   = &AchvTable{}
	GrowthConf = &GrowthTable{}
)

// ============================================================================

func (a *AchvTable) Load() bool {
	file := "Achv.json"
	var arr []*Achv

	if !load_json_as_arr(C_Config_Path+file, &arr) {
		return false
	}

	a.items = make(map[int32]*Achv)
	for _, v := range arr {
		a.items[v.Id] = v
	}

	log.Info("load [ %s ] OK", file)
	return true
}

func (a *AchvTable) Query(id int32) *Achv {
	return a.items[id]
}

func (a *AchvTable) Items() map[int32]*Achv {
	return a.items
}

// ============================================================================

func (a *GrowthTable) Load() bool {
	file := "Growth.json"
	var arr []*Growth

	if !load_json_as_arr(C_Config_Path+file, &arr) {
		return false
	}

	a.items = make(map[int32]*Growth)
	for _, v := range arr {
		a.items[v.Id] = v
	}

	log.Info("load [ %s ] OK", file)
	return true
}

func (a *GrowthTable) Query(id int32) *Growth {
	return a.items[id]
}

func (a *GrowthTable) Items() map[int32]*Growth {
	return a.items
}
