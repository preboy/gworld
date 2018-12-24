package config

import (
	"core/log"
)

// ============================================================================

type Object struct {
	Id      uint32 `json:"id"`
	Name    string `json:"name"`
	Title   string `json:"title"`
	SceneId uint32 `json:"sceneId"`
	X       uint32 `json:"x"`
	Y       uint32 `json:"y"`
}

type ObjectTable struct {
	items map[uint32]*Object
}

// ============================================================================

var (
	ObjectConf = &ObjectTable{}
)

// ============================================================================

func (self *ObjectTable) Load() bool {
	file := "Object.json"
	var arr []*Object

	if !load_json_as_arr(C_Config_Path+file, &arr) {
		return false
	}

	self.items = make(map[uint32]*Object)
	for _, v := range arr {
		self.items[v.Id] = v
	}

	log.Info("load [ %s ] OK", file)
	return true
}

func (self *ObjectTable) Query(id uint32) *Object {
	return self.items[id]
}

func (self *ObjectTable) Items() map[uint32]*Object {
	return self.items
}
