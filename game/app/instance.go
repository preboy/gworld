package app

type Instance struct {
	Id    uint32            // 原型ID
	Iid   uint32            // 实例ID
	Prop  map[uint32]int64  // 特殊属性：数值
	PropS map[uint32]string // 特殊属性：字符串
}

func NewInstance(id uint32) *Instance {
	return &Instance{
		Id:    id,
		Iid:   GetServerData().GetSeqInstance(),
		Prop:  map[uint32]int64{},
		PropS: map[uint32]string{},
	}
}

func (self *Instance) SetProp(id uint32, val int64) {
	self.Prop[id] = val
}

func (self *Instance) GetProp(id uint32) int64 {
	return self.Prop[id]
}

func (self *Instance) DelProp(id uint32) {
	delete(self.Prop, id)
}

func (self *Instance) SetPropS(id uint32, val string) {
	self.PropS[id] = val
}

func (self *Instance) GetPropS(id uint32) string {
	return self.PropS[id]
}

func (self *Instance) DelPropS(id uint32) {
	delete(self.Prop, id)
}
