package app

import (
	"public/protocol"
	"public/protocol/msg"
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
	self.add[id] += cnt
}

func (self *ItemProxy) Sub(id uint32, cnt uint64) {
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
	res := msg.ItemCntChangedNotice{}

	for id, cnt := range self.add {
		plr.AddItem(id, cnt)
		res.Info = append(res.Info, &msg.ItemCntInfo{
			Add: 1,
			Id:  id,
			Cnt: cnt,
		})
	}

	for id, cnt := range self.sub {
		if plr.GetItem(id) >= cnt {
			plr.SubItem(id, cnt)
			res.Info = append(res.Info, &msg.ItemCntInfo{
				Add: 2,
				Id:  id,
				Cnt: cnt,
			})
		} else {
			plr.SetItem(id, 0)
			res.Info = append(res.Info, &msg.ItemCntInfo{
				Add: 2,
				Id:  id,
				Cnt: 0,
			})
		}
	}

	plr.SendPacket(protocol.MSG_SC_ItemCntChanged, &res)

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
