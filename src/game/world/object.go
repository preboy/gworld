package world

import (
	"game/config"
)

const (
	C_SOBJ_Npc     int32 = 1
	C_SOBJ_Map     int32 = 2
	C_SOBJ_Monster int32 = 3
)

// ============================================================================

// 场景对象
type Object struct {
	proto *config.Object
	scene *Scene
	stat  int32
	args  map[string]string
}

func (self *Object) GetId() uint32 {
	return self.proto.Id
}

func (self *Object) GetPos() (int32, int32) {
	return self.proto.X, self.proto.Y
}

func (self *Object) GetName() string {
	return self.proto.Name
}

func (self *Object) GetProto() *config.Object {
	return self.proto
}

// ============================================================================

type objectMgr struct {
	objs map[uint32]*Object // [id]obj
}

func (self *objectMgr) AddObject(obj *Object) {
	self.objs[obj.GetId()] = obj
}

func (self *objectMgr) GetObject(id uint32) *Object {
	return self.objs[id]
}
