package app

import (
	"gworld/core/log"
	"gworld/game/config"
	"gworld/public/protocol"
	"gworld/public/protocol/msg"
)

type ItemProxy struct {
	msg uint32
	arg []uint32
	add map[uint32]uint64
	sub map[uint32]uint64
}

func NewItemProxy(msg uint32) *ItemProxy {
	ib := &ItemProxy{
		msg: msg,
	}
	ib.add = make(map[uint32]uint64)
	ib.sub = make(map[uint32]uint64)
	return ib
}

func (self *ItemProxy) SetArgs(args ...uint32) *ItemProxy {
	self.arg = args
	return self
}

func (self *ItemProxy) Add(id uint32, cnt uint64) {
	if InDebugMode() {
		conf := config.ItemProtoConf.Query(id)
		if conf == nil {
			log.Warning("item NOT EXIST = %v", id)
		}
	}

	self.add[id] += cnt
}

func (self *ItemProxy) Sub(id uint32, cnt uint64) {
	if InDebugMode() {
		conf := config.ItemProtoConf.Query(id)
		if conf == nil {
			log.Warning("item NOT EXIST = %v", id)
		}
	}

	self.sub[id] += cnt
}

// 检测包裹里是否有足够的道具
func (self *ItemProxy) Enough(plr IPlayer) bool {
	for id, cnt := range self.sub {
		if plr.GetItem(id) < cnt {
			return false
		}
	}
	return true
}

func (self *ItemProxy) Apply(plr IPlayer) *ItemProxy {
	res := msg.ItemUpdate{}

	for id, cnt := range self.add {
		plr.AddItem(id, cnt)
		res.Items = append(res.Items, &msg.Item{
			Flag: 1,
			Id:   id,
			Cnt:  int64(cnt),
		})
	}

	for id, cnt := range self.sub {
		if plr.GetItem(id) >= cnt {
			plr.SubItem(id, cnt)
			res.Items = append(res.Items, &msg.Item{
				Flag: 1,
				Id:   id,
				Cnt:  -int64(cnt),
			})
		} else {
			plr.SetItem(id, 0)
			res.Items = append(res.Items, &msg.Item{
				Flag: 0,
				Id:   id,
				Cnt:  0,
			})
		}
	}

	plr.SendPacket(protocol.MSG_SC_ItemUpdate, &res)

	return self
}

func (self *ItemProxy) ToMsg() (ret []*msg.Item) {
	for id, cnt := range self.add {
		ret = append(ret, &msg.Item{1, id, int64(cnt)})
	}
	for id, cnt := range self.sub {
		ret = append(ret, &msg.Item{1, id, -int64(cnt)})
	}
	return
}
