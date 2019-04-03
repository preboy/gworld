package world

// ============================================================================
// Scene

type Scene struct {
	id   uint32
	objs map[uint32]*Object
}

func (self *Scene) AddObject(obj *Object) {
	obj.scene = self
	self.objs[obj.GetId()] = obj
}

func (self *Scene) GetObject(id uint32) *Object {
	return self.objs[id]
}

func (self *Scene) DelObject(obj *Object) {
	self.objs[obj.GetId()] = nil
}

func (self *Scene) Objects() map[uint32]*Object {
	return self.objs
}

// ============================================================================
// sceneMgr

type sceneMgr struct {
	scenes map[uint32]*Scene
}

func (self *sceneMgr) GetScene(id uint32) *Scene {
	return self.scenes[id]
}
