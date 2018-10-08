package config

import (
	"core/event"
	"core/log"
	"core/utils"
	"server/constant"

	"reflect"
	"strings"
)

const (
	C_Config_Path = "./config/"
)

type IConf interface {
	Load() bool
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

var (
	load_json_as_arr = utils.LoadJsonAsArr
)

// ============================================================================

func LoadAll(launch bool) {
	log.Info("Loading Configs Starting ...")

	succ := true
	for _, conf := range confs {
		if !conf.Load() {
			succ = false
		}
	}

	if launch {
		if !succ {
			log.Fatal("load config NOT all is ok ")
		}
	}

	log.Info("Loading Configs COMPLETE !!!")

	event.Fire(constant.EVT_SYS_ConfigLoaded, launch)
}

func LoadOne(name string) bool {
	for _, conf := range confs {
		t := reflect.TypeOf(conf).Elem()
		file_name := strings.TrimSuffix(t.Name(), "Table")

		if file_name == name {
			return conf.Load()
		}
	}

	log.Info("Loading [ %s ] FAILED, Not found the file !!!")

	return false
}
