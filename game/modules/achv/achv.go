package achv

import (
	"gworld/game/config"

	"gopkg.in/mgo.v2/bson"
)

// ============================================================================
// regular

type achv_t struct {
	Id       int32 // achv id
	GrowthId int32
	Taken    bool
}

type achv_map_t map[int32]*achv_t

type Achv struct {
	plr iPlayer

	AchvV achv_map_t
}

// ============================================================================
// marshalling

func (self achv_map_t) GetBSON() (interface{}, error) {
	var arr []*achv_t

	for _, v := range self {
		arr = append(arr, v)
	}

	return arr, nil
}

func (self *achv_map_t) SetBSON(raw bson.Raw) error {
	var arr []*achv_t

	err := raw.Unmarshal(&arr)
	if err != nil {
		return err
	}

	*self = make(achv_map_t)
	for _, v := range arr {
		(*self)[v.Id] = v
	}

	return nil
}

// ============================================================================

func NewAchv() *Achv {
	return &Achv{}
}

func (self *Achv) Init(plr iPlayer) {
	self.plr = plr

	if self.AchvV == nil {
		self.AchvV = make(achv_map_t)
	}

	for _, conf := range config.AchvConf.Items() {
		if self.AchvV[conf.Id] == nil || self.AchvV[conf.Id].GrowthId != conf.Gid {
			self.AchvV[conf.Id] = &achv_t{
				Id:       conf.Id,
				GrowthId: conf.Gid,
			}
		}
	}

}

// ============================================================================

func (self *Achv) IsCompleted(id int32) bool {
	conf := config.AchvConf.Query(id)
	if conf != nil {
		return false
	}

	return self.GetVal(conf.Gid) >= conf.Val
}

func (self *Achv) GetVal(id int32) int32 {
	conf := config.AchvConf.Query(id)
	if conf == nil {
		return 0
	}

	gid := conf.Gid

	growth := self.plr.GetGrowth()
	if growth.GrowthV[gid] == nil {
		return 0
	}

	return growth.GrowthV[gid].Val
}
