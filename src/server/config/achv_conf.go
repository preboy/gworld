package config

import (
	"core/log"
)

// ============================================================================

type Achv struct {
	Id  uint32 `json:"id"`
	Gid uint32 `json:"gid"`
	Val uint32 `json:"val"`
}

type AchvTable struct {
	items map[uint32]*Achv
}

type Growth struct {
	Id uint32 `json:"id"`
}

type GrowthTable struct {
	items map[uint32]*Growth
}

// ============================================================================

var (
	AchvConf   = &AchvTable{}
	GrowthConf = &GrowthTable{}
)

// ============================================================================

func (self *AchvTable) Load() {
	file := "Achv.json"
	var arr []*Achv

	if !load_from_json(file, arr) {
		return
	}

	self.items = make(map[uint32]*Achv)
	for _, v := range arr {
		self.items[v.Id] = v
	}

	log.Info("[%s] load OK", file)
}

func (self *AchvTable) Query(id uint32) *Achv {
	return self.items[id]
}

func (self *AchvTable) Items() map[uint32]*Achv {
	return self.items
}

// ============================================================================

func (self *GrowthTable) Load() {
	file := "Growth.json"
	var arr []*Growth

	if !load_from_json(file, arr) {
		return
	}

	self.items = make(map[uint32]*Growth)
	for _, v := range arr {
		self.items[v.Id] = v
	}

	log.Info("[%s] load OK", file)
}

func (self *GrowthTable) Query(id uint32) *Growth {
	return self.items[id]
}

func (self *GrowthTable) Items() map[uint32]*Growth {
	return self.items
}
