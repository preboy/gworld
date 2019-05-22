package app

import (
	"core/log"
	"core/utils"
)

var (
	conf *Config
	game *game_t // curr game config
	this string  // curr game id
)

type common_t struct {
	Plat      string `json:"plat"`
	DebugMode bool   `json:"debug_mode"`
	DBCenter  string `json:"db_c"`
}

type sdk_t struct {
	Port int `json:"port"`
}

type auth_t struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type bill_t struct {
	Port int `json:"port"`
}

type pull_t struct {
	Port int `json:"port"`
}
type push_t struct {
	Port int `json:"port"`
}

type admin_t struct {
	Port int `json:"port"`
}

type router_t struct {
	Port int `json:"port"`
}

type game_t struct {
	Host   string `json:"host"`
	Port   uint32 `json:"port"`
	DBGame string `json:"db_game"`
	DBStat string `json:"db_stat"`
}

type Config struct {
	Common *common_t          `json:"common"`
	Sdk    *sdk_t             `json:"sdk"`
	Auth   *auth_t            `json:"auth"`
	Bill   *bill_t            `json:"bill"`
	Pull   *pull_t            `json:"pull"`
	Push   *push_t            `json:"push"`
	Admin  *admin_t           `json:"admin"`
	Router *router_t          `json:"router"`
	Games  map[string]*game_t `json:"games"`
}

func LoadConfig(file string, svr string) bool {
	var obj *Config

	if !utils.LoadJsonAsObj(file, &obj) {
		return false
	}

	if obj == nil {
		return false
	}

	this = svr
	conf = obj

	for k, v := range conf.Games {
		if k == svr {
			game = v
		}
	}

	if game == nil {
		log.Fatal("Not Found game id: %s", svr)
	}

	log.Info("load [ %s ] OK", file)
	return true
}

func GetPlat() string {
	return conf.Common.Plat
}

func GetGameId() string {
	return this
}

func GetConfig() *Config {
	return conf
}

func GetGameConfig() *game_t {
	return game
}

func InDebugMode() bool {
	if conf != nil {
		return conf.Common.DebugMode
	}
	return false
}
