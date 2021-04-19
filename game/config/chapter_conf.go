package config

import (
	"gworld/core/log"
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

func (c *ChapterTable) Load() bool {
	file := "Chapter.json"
	var arr []*Chapter

	if !load_json_as_arr(C_Config_Path+file, &arr) {
		return false
	}

	c.items = make(map[uint32]*Chapter)
	for _, v := range arr {
		key := v.Id
		c.items[key] = v
	}

	log.Info("load [ %s ] OK", file)
	return true
}

func (c *ChapterTable) Query(lv uint32) *Chapter {
	return c.items[lv]
}

func (c *ChapterTable) Items() map[uint32]*Chapter {
	return c.items
}
