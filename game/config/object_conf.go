package config

import (
	"gworld/core/log"
)

// ============================================================================

type Object struct {
	Id      uint32 `json:"id"`
	Name    string `json:"name"`
	Title   string `json:"title"`
	SceneId uint32 `json:"sceneId"`
	X       int32  `json:"x"`
	Y       int32  `json:"y"`
}

type ObjectTable struct {
	items map[uint32]*Object
}

// ============================================================================

var (
	ObjectConf = &ObjectTable{}
)

// ============================================================================

func (o *ObjectTable) Load() bool {
	file := "Object.json"
	var arr []*Object

	if !load_json_as_arr(C_Config_Path+file, &arr) {
		return false
	}

	o.items = make(map[uint32]*Object)
	for _, v := range arr {
		o.items[v.Id] = v
	}

	log.Info("load [ %s ] OK", file)
	return true
}

func (o *ObjectTable) Query(id uint32) *Object {
	return o.items[id]
}

func (o *ObjectTable) Items() map[uint32]*Object {
	return o.items
}
