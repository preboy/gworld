package app

import (
	"gopkg.in/mgo.v2/bson"
)

// ----------------------------------------------------------------------------

type data_t struct {
	K, V int32
}

// ----------------------------------------------------------------------------

type KV_map_t map[int32]int32

func (self KV_map_t) GetBSON() (interface{}, error) {
	var arr []*data_t
	for k, v := range self {
		arr = append(arr, &data_t{K: k, V: v})
	}

	return arr, nil
}

func (self *KV_map_t) SetBSON(raw bson.Raw) error {
	var arr []*data_t
	err := raw.Unmarshal(&arr)
	if err != nil {
		return err
	}

	*self = make(KV_map_t)
	for _, v := range arr {
		(*self)[v.K] = v.V
	}

	return nil
}

// ----------------------------------------------------------------------------
