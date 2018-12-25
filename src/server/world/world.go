package world

import (
	"server/config"
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
		SceneMgr.GetScene(v.SceneId).AddObject(obj)
	}
}

// ============================================================================

func Start() {
	init_world()
}

func Stop() {

}
