package player

import (
	_ "core/log"
	"public/protocol"
	"public/protocol/msg"
)

// 道具是否存在

type ItemProxy struct {
	msg uint16
	add map[uint32]uint64
	sub map[uint32]uint64
}

func NewItemProxy(msg uint16) *ItemProxy {
	ib := &ItemProxy{
		msg: msg,
	}
	ib.add = make(map[uint32]uint64)
	ib.sub = make(map[uint32]uint64)
	return ib
}

func (self *ItemProxy) Add(id uint32, cnt uint64) {
	self.add[id] += cnt
}

func (self *ItemProxy) Sub(id uint32, cnt uint64) {
	self.sub[id] += cnt
}

// 检测包裹里是否有足够的道具
func (self *ItemProxy) Enough(plr *Player) bool {
	Items := plr.data.Items
	for id, cnt := range self.sub {
		if Items[id] < cnt {
			return false
		}
	}
	return true
}

func (self *ItemProxy) Apply(plr *Player) {
	res := msg.ItemCntChangedNotice{}
	Items := plr.data.Items
	for id, cnt := range self.add {
		Items[id] += cnt
		res.Info = append(res.Info, &msg.ItemCntInfo{
			Add: 1,
			Id:  id,
			Cnt: cnt,
		})
	}
	for id, cnt := range self.sub {
		if Items[id] >= cnt {
			Items[id] -= cnt
			res.Info = append(res.Info, &msg.ItemCntInfo{
				Add: 2,
				Id:  id,
				Cnt: cnt,
			})
		} else {
			Items[id] = 0
			res.Info = append(res.Info, &msg.ItemCntInfo{
				Add: 2,
				Id:  id,
				Cnt: 0,
			})
		}
	}

	plr.SendPacket(protocol.MSG_SC_ItemCntChanged, &res)
}

func (self *Player) GetItemCnt(id uint32) uint64 {
	return self.data.Items[id]
}
