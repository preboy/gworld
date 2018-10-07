package config

import (
	"core/event"
	"core/log"
	"server/constant"

	"encoding/json"
	"io/ioutil"
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

// ============================================================================

func load_json_as_arr(file string, arr interface{}) bool {
	content, err := ioutil.ReadFile(C_Config_Path + file)
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
