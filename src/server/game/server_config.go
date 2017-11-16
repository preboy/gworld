package game

import (
	"encoding/json"
)

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
}

func LoadServerConfig(file string) bool {
	if _sc == nil {
		sc := ServerConfig{}
		data, err := ReadFile(file)
		if err == nil {
			err := json.Unmarshal(data, &sc)
			if err == nil {
				_sc = &sc
			} else {
				log.GetLogger().Debug("zcg_err2 : %s", err.Error())
				return false
			}
		} else {
			log.GetLogger().Debug("zcg_err : %s", err.Error())
			return false
		}
	}

	utils.PrintPretty(_sc, "server cnfig ")
	return true
}

func GetServerConfig() *ServerConfig {
	return _sc
}
