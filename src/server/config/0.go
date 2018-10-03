package config

import (
	"core/log"
	"encoding/json"
	"io/ioutil"
)

const (
	C_Path = "config/"
)

type IConf interface {
	Load()
}

var (
	confs = []IConf{
		AchvConf,
		ActivityConf,
		AuraProtoConf,
		CreatureConf,
		CreatureTeamConf,
		GlobalConf,
		GrowthConf,
		HeroConf,
		MarketConf,
		ItemProtoConf,
		RefineNormalConf,
		RefineSuperConf,
		SkillProtoConf,
	}
)

// ============================================================================

func Load() {
	log.Info("Loading Config Starting ...")

	for _, conf := range confs {
		conf.Load()
	}

	log.Info("Loading Config COMPLETE !!!")
}

func load_from_json(file string, arr interface{}) bool {
	content, err := ioutil.ReadFile(C_Path + file)
	if err != nil {
		log.Error("loading [%s] failed! err = %s", file, err)
		return false
	}

	err = json.Unmarshal(content, &arr)
	if err != nil {
		log.Error("Unmarshal [%s] failed! %s", file, err)
		return false
	}

	return true
}
