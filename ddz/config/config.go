package config

import (
	"encoding/json"
	"os"
)

var (
	_conf config
)

// Conf Conf
type config struct {
	Addr4Gambler string `json:"addr_gambler"`
	Addr4Referee string `json:"addr_referee"`
}

// ----------------------------------------------------------------------------

func Load() error {
	data, err := os.ReadFile("./ddz_config.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &_conf)
	if err != nil {
		return err
	}

	return nil
}

func Get() *config {
	return &_conf
}
