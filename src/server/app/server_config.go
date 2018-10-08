package app

import (
	"core/log"
	"core/utils"
)

var (
	_sc *ServerConfig
)

type ServerConfig struct {
	Server_id   uint32 `json:"server_id"`
	Server_name string `json:"server_name"`
	Listen_addr string `json:"listen_addr"`
	PlatID      string `json:"plat_id"`
	DBAddr      string `json:"db_addr"`
	DebugMode   bool   `json:"debug_mode"`
}

func LoadServerConfig(file string) bool {
	var obj *ServerConfig

	if !utils.LoadJsonAsObj(file, &obj) {
		return false
	}

	if obj == nil {
		return false
	}

	_sc = obj
	utils.PrintPretty(_sc, "server cnfig ")

	log.Info("load [ %s ] OK", file)
	return true
}

func GetServerConfig() *ServerConfig {
	return _sc
}

func InDebugMode() bool {
	if _sc != nil {
		return _sc.DebugMode
	}
	return false
}
