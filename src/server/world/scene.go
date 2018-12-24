package world

// ============================================================================
// Scene

type Scene struct {
	id   int32
	objs map[int32]*Object
}

func (self *Scene) AddObject(x, y, obj *Object) bool {
	id := obj.GetId()
	if self.GetObject(id) != nil {
		return false
	}

	_, _, scene := obj.GetPos()
	if scene != nil {
		scene.DelObject(obj)
	}

	obj.SetPos(x, y, self)
	self.objs[id] = obj

	return true
}

func (self *Scene) GetObject(id int32) *Object {
	return self.objs[id]
}

func (self *Scene) DelObject(obj *Object) {
	self.objs[obj.GetId()] = nil
}

// ============================================================================
// SceneMgr

type SceneMgr struct {
	scenes map[int32]*Scene
}

func (self *Scene) GetScene(id int32) *Scene {
	return self.scenes[id]
}

// ============================================================================
// export

func Start() {

}

func Close() {

}
