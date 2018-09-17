package config

import (
	"core/log"
	"encoding/json"
	"io/ioutil"
)

type ActivityItem struct {
	Seq   int    `json:"seq"`
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Desc  string `json:"desc"`
	Open  string `json:"open"`
	Close string `json:"close"`
}

type ActivityConf struct {
	items map[int]*ActivityItem
}

var _ActivityConf ActivityConf

func GetActivityConf() *ActivityConf {
	return &_ActivityConf
}

func load_activity() {
	path := "config/ActivityConf.json"

	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Error("[ActivityConf] loading failed: %s: %s", path, err)
		return
	}

	var arr []*ActivityItem
	err = json.Unmarshal(content, &arr)
	if err != nil {
		log.Error("[ActivityConf] Unmarshal failed: %s: %s", path, err)
		return
	}

	_ActivityConf.items = make(map[int]*ActivityItem)

	for _, v := range arr {
		key := v.Seq
		_ActivityConf.items[key] = v
	}

	log.Info("[ActivityConf] load OK")
}

// --------------------------------------------------------------------

func (self *ActivityConf) ForEach(f func(*ActivityItem)) {
	for _, item := range self.items {
		f(item)
	}
}
