package world

const (
	C_SOBJ_Npc     int32 = 1
	C_SOBJ_Map     int32 = 2
	C_SOBJ_Monster int32 = 3
)

// ============================================================================

// 场景对象
type Object struct {
	x     int32
	y     int32
	proto *config.Object
	scene *Scene
	stat  int32
	args  map[string]string
}

func (self *Object) GetId() int32 {
	return self.GetProto.Id()
}

func (self *Object) GetProto() *config.Object {
	return self.proto
}

func (self *Object) SetPos(x, y int32, scene *Scene) {
	self.x = x
	self.y = y
	self.scene = scene
}

func (self *Object) GetPos() (int32, int32, *Scene) {
	return self.x, self.y, self.scene
}

// ============================================================================

type ObjectMgr struct {
	objs map[int32]*Object // [id]obj
}

func (self *ObjectMgr) AddObject(obj *Object) {
	self.objs[obj.GetId()] = obj
}

func (self *ObjectMgr) GetObject(id int32) *Object {
	return self.objs[id]
}

// ============================================================================
// export

func CreateObject() *Object {
	return nil
}
