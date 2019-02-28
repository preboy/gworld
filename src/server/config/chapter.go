package config

import (
	"core/log"
)

// ============================================================================

type Chapter struct {
	Id         uint32 `json:"id"`
	Name       string `json:"name"`
	BreakStart uint32 `json:"breakStart"`
	BreakEnd   uint32 `json:"breakEnd"`
	DropId     uint32 `json:"dropId"`
}

type ChapterTable struct {
	items map[uint32]*Chapter
}

// ============================================================================

var (
	ChapterConf = &ChapterTable{}
)

// ============================================================================

func (self *ChapterTable) Load() bool {
	file := "Chapter.json"
	var arr []*Chapter

	if !load_json_as_arr(C_Config_Path+file, &arr) {
		return false
	}

	self.items = make(map[uint32]*Chapter)
	for _, v := range arr {
		key := v.Id
		self.items[key] = v
	}

	log.Info("load [ %s ] OK", file)
	return true
}

func (self *ChapterTable) Query(lv uint32) *Chapter {
	return self.items[lv]
}

func (self *ChapterTable) Items() map[uint32]*Chapter {
	return self.items
}
