package config

import (
	"core/log"
	"encoding/json"
	"io/ioutil"
)

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

type CreatureTeamConf struct {
	items map[uint32]*CreatureTeam
}

var _CreatureTeamConf CreatureTeamConf

func GetCreatureTeamConf() *CreatureTeamConf {
	return &_CreatureTeamConf
}

func load_creature_team() {
	path := "config/CreatureTeam.json"

	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Error("[CreatureTeamConf] loading failed: %s: %s", path, err)
		return
	}

	var arr []*CreatureTeam
	err = json.Unmarshal(content, &arr)
	if err != nil {
		log.Error("[CreatureTeamConf] Unmarshal failed: %s: %s", path, err)
		return
	}

	_CreatureTeamConf.items = make(map[uint32]*CreatureTeam, 0x400)

	for _, v := range arr {
		_CreatureTeamConf.items[v.Id] = v
	}

	log.Info("[CreatureTeamConf] load OK")
}

func GetCreatureTeam(id uint32) *CreatureTeam {
	if _CreatureTeamConf.items == nil {
		return nil
	}

	return _CreatureTeamConf.items[id]
}
