package config

import (
	"core/log"
)

// ============================================================================

type PosCreatureInfo struct {
	Id uint32
	Lv uint32
}

type CreatureTeam struct {
	Id        uint32             `json:"id"`
	Desc      string             `json:"name"`
	L_Pioneer []*PosCreatureInfo `json:"l_pioneer"`
	R_Pioneer []*PosCreatureInfo `json:"r_pioneer"`
	Commander []*PosCreatureInfo `json:"commander"`
	L_Guarder []*PosCreatureInfo `json:"l_guarder"`
	R_Guarder []*PosCreatureInfo `json:"r_guarder"`
}

type CreatureTeamTable struct {
	items map[uint32]*CreatureTeam
}

// ============================================================================

var (
	CreatureTeamConf = &CreatureTeamTable{}
)

// ============================================================================

func (self *CreatureTeamTable) Load() {
	file := "CreatureTeam.json"
	var arr []*CreatureTeam

	if !load_from_json(file, arr) {
		return
	}

	self.items = make(map[uint32]*CreatureTeam)
	for _, v := range arr {
		self.items[v.Id] = v
	}

	log.Info("[%s] load OK", file)
}

func (self *CreatureTeamTable) Query(id uint32) *CreatureTeam {
	return self.items[id]
}

func (self *CreatureTeamTable) Items() map[uint32]*CreatureTeam {
	return self.items
}
