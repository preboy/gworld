package player

import (
	"server/modules/achv"
)

func (self *Player) GetId() string {
	return self.data.Pid
}

func (self *Player) GetGrowth() *achv.Growth {
	return self.data.Growth
}

func (self *Player) GetAchv() *achv.Achv {
	return self.data.Achv
}
