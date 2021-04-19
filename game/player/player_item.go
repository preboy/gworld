package player

import (
	"gopkg.in/mgo.v2/bson"
)

type item_map_t map[uint32]uint64

// ============================================================================
// marshal

type item_t struct {
	Id  uint32
	Cnt uint64
}

func (self item_map_t) GetBSON() (interface{}, error) {
	var arr []*item_t

	for k, v := range self {
		arr = append(arr, &item_t{
			Id:  k,
			Cnt: v,
		})
	}

	return arr, nil
}

func (self *item_map_t) SetBSON(raw bson.Raw) error {
	var arr []*item_t

	err := raw.Unmarshal(&arr)
	if err != nil {
		return err
	}

	*self = make(item_map_t)
	for _, v := range arr {
		(*self)[v.Id] = v.Cnt
	}

	return nil
}

// ============================================================================

func (self *Player) GetItem(id uint32) uint64 {
	return self.data.Items[id]
}

func (self *Player) SetItem(id uint32, cnt uint64) {
	self.data.Items[id] = cnt
}

func (self *Player) AddItem(id uint32, cnt uint64) {
	self.data.Items[id] += cnt
}

func (self *Player) SubItem(id uint32, cnt uint64) {
	self.data.Items[id] -= cnt
}
