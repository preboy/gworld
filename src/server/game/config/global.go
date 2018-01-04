package config

import (
	"core/log"
	"encoding/json"
	"io/ioutil"
)

type GlobalConf struct {
	Name  string   `json:"name"`
	GID   uint32   `json:"gID"`
	Rate  float64  `json:"rate"`
	Enc   []*ENC   `json:"enc"`
	Level []uint32 `json:"level"`
}

var _GlobalConf *GlobalConf

func GetGlobalConf() *GlobalConf {
	return _GlobalConf
}

func load_global() {
	path := "config/global.json"

	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Error("[GlobalConf] loading failed: %s: %s", path, err)
		return
	}

	var arr []*GlobalConf
	err = json.Unmarshal(content, &arr)
	if err != nil {
		log.Error("[GlobalConf] Unmarshal failed: %s: %s", path, err)
		return
	}

	_GlobalConf = arr[0]

	log.Info("[GlobalConf] load OK")
}
