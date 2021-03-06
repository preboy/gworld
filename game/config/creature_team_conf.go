package config

import (
	"gworld/core/log"
)

// ============================================================================

type CreatureTeam struct {
	Id    uint32 `json:"id"`
	Desc  string `json:"name"`
	Row11 uint32 `json:"row11"`
	Row12 uint32 `json:"row12"`
	Row13 uint32 `json:"row13"`
	Row21 uint32 `json:"row21"`
	Row22 uint32 `json:"row22"`
	Row23 uint32 `json:"row23"`
}

type CreatureTeamTable struct {
	items map[uint32]*CreatureTeam
}

// ============================================================================

var (
	CreatureTeamConf = &CreatureTeamTable{}
)

// ============================================================================

func (c *CreatureTeamTable) Load() bool {
	file := "CreatureTeam.json"
	var arr []*CreatureTeam

	if !load_json_as_arr(C_Config_Path+file, &arr) {
		return false
	}

	c.items = make(map[uint32]*CreatureTeam)
	for _, v := range arr {
		c.items[v.Id] = v
	}

	log.Info("load [ %s ] OK", file)
	return true
}

func (c *CreatureTeamTable) Query(id uint32) *CreatureTeam {
	return c.items[id]
}
