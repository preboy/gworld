package config

import (
	"gworld/core/log"
)

// ============================================================================

type Activity struct {
	Seq   int32  `json:"seq"`
	Id    int32  `json:"id"`
	Name  string `json:"name"`
	Desc  string `json:"desc"`
	Open  string `json:"open"`
	Close string `json:"close"`
}

type ActivityTable struct {
	items map[int32]*Activity
}

// ============================================================================

var (
	ActivityConf = &ActivityTable{}
)

// ============================================================================

func (a *ActivityTable) Load() bool {
	file := "Activity.json"
	var arr []*Activity

	if !load_json_as_arr(C_Config_Path+file, &arr) {
		return false
	}

	a.items = make(map[int32]*Activity)
	for _, v := range arr {
		a.items[v.Id] = v
	}

	log.Info("load [ %s ] OK", file)
	return true
}

func (a *ActivityTable) Query(id int32) *Activity {
	return a.items[id]
}

func (a *ActivityTable) Items() map[int32]*Activity {
	return a.items
}

// ============================================================================

func (a *ActivityTable) ForEach(f func(*Activity)) {
	for _, item := range a.items {
		f(item)
	}
}
