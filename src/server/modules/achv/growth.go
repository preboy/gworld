package achv

import (
	"gopkg.in/mgo.v2/bson"
	"server/app"
)

// ============================================================================
// regular

type iPlayer interface {
	app.IPlayer

	GetGrowth() *Growth
}

// ============================================================================
// base

type growth_t struct {
	Id  int32
	Val int32

	care []int32 // associated achv id
}

type growth_map_t map[int32]*growth_t

type Growth struct {
	plr iPlayer

	// saved data
	GrowthV growth_map_t

	// unsaved data
	achv map[int32]bool
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

	if self.achv == nil {
		self.achv = make(map[int32]bool)
	}
}

// ============================================================================
