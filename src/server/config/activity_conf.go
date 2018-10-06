package config

import (
	"core/log"
)

// ============================================================================

type Activity struct {
	Seq   int    `json:"seq"`
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Desc  string `json:"desc"`
	Open  string `json:"open"`
	Close string `json:"close"`
}

type ActivityTable struct {
	items map[int]*Activity
}

// ============================================================================

var (
	ActivityConf = &ActivityTable{}
)

// ============================================================================

func (self *ActivityTable) Load() bool {
	file := "Activity.json"
	var arr []*Activity

	if !load_json_as_arr(file, &arr) {
		return false
	}

	self.items = make(map[int]*Activity)
	for _, v := range arr {
		self.items[v.Id] = v
	}

	log.Info("load [ %s ] OK", file)
	return true
}

func (self *ActivityTable) Query(id int) *Activity {
	return self.items[id]
}

func (self *ActivityTable) Items() map[int]*Activity {
	return self.items
}

// ============================================================================

func (self *ActivityTable) ForEach(f func(*Activity)) {
	for _, item := range self.items {
		f(item)
	}
}
