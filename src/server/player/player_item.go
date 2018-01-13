package player

import (
	"core/log"
)

type ItemProxy struct {
	items map[uint32]int64
}

func NewItemProxy() *ItemProxy {
	ib := &ItemProxy{}
	ib.items = make(map[uint32]int64)
	return ib
}

func (self *ItemProxy) Add(id uint32, cnt int64) {
	if cnt <= 0 {
		log.Error("ItemProxy.Add Param Error:", id, cnt)
		return
	}
	self.items[id] += cnt
}

func (self *ItemProxy) Sub(id uint32, cnt int64) {
	if cnt >= 0 {
		log.Error("ItemProxy.Sub Param Error:", id, cnt)
		return
	}
	self.items[id] -= cnt
}

// 检测包裹里是否有足够的道具
func (self *ItemProxy) Enough(plr *Player) bool {
	Items := plr.data.Items
	for id, cnt := range self.items {
		if cnt < 0 {
			if Items[id] < uint64(-cnt) {
				return false
			}
		}
	}
	return true
}

func (self *ItemProxy) Apply(plr *Player) {
	Items := plr.data.Items
	for id, cnt := range self.items {
		if cnt < 0 {
			if Items[id] >= uint64(-cnt) {
				Items[id] -= uint64(-cnt)
			}
		} else {
			Items[id] += uint64(cnt)
		}
	}
}

func (self *Player) GetItemCnt(id uint32) uint64 {
	return self.data.Items[id]
}
