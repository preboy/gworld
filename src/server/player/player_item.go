package player

import (
	_ "core/log"
	"server/game/config"
)

// 道具是否存在

type ItemProxy struct {
	add map[uint32]uint64
	sub map[uint32]uint64
}

func NewItemProxy() *ItemProxy {
	ib := &ItemProxy{}
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
	Items := plr.data.Items
	for id, cnt := range self.add {
		Items[id] += cnt
	}
	for id, cnt := range self.sub {
		if Items[id] >= cnt {
			Items[id] -= cnt
		} else {
			Items[id] = 0
		}
	}
}

func (self *Player) GetItemCnt(id uint32) uint64 {
	return self.data.Items[id]
}

// 玩家使用道具(常规道具)
func (self *Player) DoItem(id uint32, cnt uint64) bool {
	ip := config.GetItemProtoConf().ItemProto(id)
	if ip == nil || ip.Sid == 0 {
		return false
	}

	goods := NewItemProxy()
	goods.Sub(id, cnt)
	if !goods.Enough(self) {
		return false
	}
	goods.Apply(self)

	if script, ok := _item_scripts[ip.Sid]; ok {
		script(self, ip.Param1, ip.Param2, ip.Param3, ip.Param4)
	}

	return true
}
