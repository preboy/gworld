package loop

import (
	"game/player"
)

// 服务器停止前，请在此存储数据
func (self *Loop) onstop() {
	player.EachPlayer(func(plr *player.Player) {
		plr.Disconnect()
	})
}
