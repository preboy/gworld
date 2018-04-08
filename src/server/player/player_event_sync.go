package player

import (
	"fmt"
)

func (self *Player) OnPlayerLevelup(lv_old, lv_new uint32) {
	fmt.Println("Player.OnPlayerLevelup", lv_old, lv_new)
}
