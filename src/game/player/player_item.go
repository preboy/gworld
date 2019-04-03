package player

import (
	_ "core/log"
)

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
