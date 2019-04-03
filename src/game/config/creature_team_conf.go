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
	Id    uint32             `json:"id"`
	Desc  string             `json:"name"`
	Row11 []*PosCreatureInfo `json:"row11"`
	Row12 []*PosCreatureInfo `json:"row12"`
	Row13 []*PosCreatureInfo `json:"row13"`
	Row21 []*PosCreatureInfo `json:"row21"`
	Row22 []*PosCreatureInfo `json:"row22"`
	Row23 []*PosCreatureInfo `json:"row23"`
}

type CreatureTeamTable struct {
	items map[uint32]*CreatureTeam
}

// ============================================================================

var (
	CreatureTeamConf = &CreatureTeamTable{}
)

// ============================================================================

func (self *CreatureTeamTable) Load() bool {
	file := "CreatureTeam.json"
	var arr []*CreatureTeam

	if !load_json_as_arr(C_Config_Path+file, &arr) {
		return false
	}

	self.items = make(map[uint32]*CreatureTeam)
	for _, v := range arr {
		self.items[v.Id] = v
	}

	log.Info("load [ %s ] OK", file)
	return true
}

func (self *CreatureTeamTable) Query(id uint32) *CreatureTeam {
	return self.items[id]
}

func (self *CreatureTeamTable) Items() map[uint32]*CreatureTeam {
	return self.items
}
