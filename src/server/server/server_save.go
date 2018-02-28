package server

import (
	"server/player"
)

// 服务器停止前，请在此存储数据
func (self *Server) on_stop() {
	player.EachPlayer(func(plr *player.Player) {
		plr.Disconnect()
	})
}
