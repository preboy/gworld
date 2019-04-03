package achv

import (
	"core/event"
	"game/app"
	"game/config"
	"game/constant"
	"gopkg.in/mgo.v2/bson"
)

var (
	care map[int32][]int32
)

// ============================================================================
// regular

type iPlayer interface {
	app.IPlayer

	GetGrowth() *Growth
	GetAchv() *Achv
}

// ============================================================================
// base

type growth_t struct {
	Id  int32
	Val int32
}

type Growth struct {
	plr iPlayer

	GrowthV growth_map_t
}

type growth_map_t map[int32]*growth_t

// ============================================================================

func init() {
	care = make(map[int32][]int32)

	event.On(constant.EVT_SYS_ConfigLoaded, func(args ...interface{}) {
		launch := args[0].(bool)
		if launch {
			for _, conf := range config.AchvConf.Items() {
				care[conf.Gid] = append(care[conf.Gid], conf.Id)
			}
		}
	})
}

// ============================================================================
// marshalling

func (self growth_map_t) GetBSON() (interface{}, error) {
	var arr []*growth_t

	for _, v := range self {
		arr = append(arr, v)
	}

	return arr, nil
}

func (self *growth_map_t) SetBSON(raw bson.Raw) error {
	var arr []*growth_t

	err := raw.Unmarshal(&arr)
	if err != nil {
		return err
	}

	*self = make(growth_map_t)
	for _, v := range arr {
		(*self)[v.Id] = v
	}

	return nil
}

// ============================================================================

func NewGrowth() *Growth {
	return &Growth{}
}

func (self *Growth) Init(plr iPlayer) {
	self.plr = plr

	if self.GrowthV == nil {
		self.GrowthV = make(growth_map_t)
	}
}

func (self *Growth) OnEvent(evt *event.Event) {
	// on_changed(id, val)
}

func (self *Growth) change(id int32, val int32, set bool) {
	if self.GrowthV[id] == nil {
		return
	}

	prev := self.GrowthV[id].Val

	if set {
		self.GrowthV[id].Val = val
	} else {
		self.GrowthV[id].Val += val
	}

	self.on_changed(id, prev, self.GrowthV[id].Val)
}

// 成就变化是通知客户端
func (self *Growth) on_changed(id int32, prev, curr int32) {
	for _, achv_id := range care[id] {
		_ = achv_id
		// TODO ...
	}
}
