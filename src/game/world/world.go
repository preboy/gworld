package world

import (
	"core/log"
	"game/config"
)

var (
	SceneMgr = &sceneMgr{
		scenes: make(map[uint32]*Scene),
	}

	ObjectMgr = &objectMgr{
		objs: make(map[uint32]*Object),
	}
)

// ============================================================================

func init_world() {

	// scenes
	for _, v := range config.SceneConf.Items() {
		id := v.Id
		SceneMgr.scenes[id] = &Scene{
			id:   id,
			objs: make(map[uint32]*Object),
		}
	}

	// objects
	for _, v := range config.ObjectConf.Items() {
		obj := &Object{
			proto: v,
			args:  make(map[string]string),
		}

		ObjectMgr.AddObject(obj)

		scene := SceneMgr.GetScene(v.SceneId)
		if scene != nil {
			scene.AddObject(obj)
		} else {
			log.Error("NOT exist scene id = %d", v.SceneId)
		}
	}
}

// ============================================================================

func Start() {
	init_world()
}

func Stop() {
}
