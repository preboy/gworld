package player

import (
	"gopkg.in/mgo.v2/bson"

	"game/app"
)

type instance_map_t map[uint32]*app.Instance

// ============================================================================
// marshal

type inst_prop_i_t struct {
	Id  uint32
	Val int64
}

type inst_prop_s_t struct {
	Id  uint32
	Val string
}

type instance_t struct {
	Id    uint32
	Iid   uint32
	Prop  []*inst_prop_i_t
	PropS []*inst_prop_s_t
}

func (self instance_map_t) GetBSON() (interface{}, error) {
	var arr []*instance_t

	for _, v := range self {
		inst := &instance_t{
			Id:  v.Id,
			Iid: v.Iid,
		}

		for x, y := range v.Prop {
			inst.Prop = append(inst.Prop, &inst_prop_i_t{x, y})
		}

		for x, y := range v.PropS {
			inst.PropS = append(inst.PropS, &inst_prop_s_t{x, y})
		}

		arr = append(arr, inst)
	}

	return arr, nil
}

func (self *instance_map_t) SetBSON(raw bson.Raw) error {
	var arr []*instance_t

	err := raw.Unmarshal(&arr)
	if err != nil {
		return err
	}

	*self = make(instance_map_t)
	for _, v := range arr {
		Inst := &app.Instance{
			Id:    v.Id,
			Iid:   v.Iid,
			Prop:  make(map[uint32]int64),
			PropS: make(map[uint32]string),
		}

		for _, vv := range v.Prop {
			Inst.Prop[vv.Id] = vv.Val
		}

		for _, vv := range v.PropS {
			Inst.PropS[vv.Id] = vv.Val
		}

		(*self)[Inst.Iid] = Inst
	}

	return nil
}

// ============================================================================

func (self *Player) GetInstance(iid uint32) *app.Instance {
	return self.data.Instances[iid]
}

func (self *Player) AddInstance(pid uint32) *app.Instance {
	instance := app.NewInstance(pid)
	self.data.Instances[instance.Iid] = instance
	return instance
}

func (self *Player) DelInstance(iid uint32) {
	delete(self.data.Instances, iid)
}
