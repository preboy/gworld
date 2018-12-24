package config

import (
	"core/log"
)

// ============================================================================

type Scene struct {
	Id   uint32 `json:"id"`
	Name string `json:"name"`
}

type SceneTable struct {
	items map[uint32]*Scene
}

// ============================================================================

var (
	SceneConf = &SceneTable{}
)

// ============================================================================

func (self *SceneTable) Load() bool {
	file := "Scene.json"
	var arr []*Scene

	if !load_json_as_arr(C_Config_Path+file, &arr) {
		return false
	}

	self.items = make(map[uint32]*Scene)
	for _, v := range arr {
		self.items[v.Id] = v
	}

	log.Info("load [ %s ] OK", file)
	return true
}

func (self *SceneTable) Query(id uint32) *Scene {
	return self.items[id]
}

func (self *SceneTable) Items() map[uint32]*Scene {
	return self.items
}
