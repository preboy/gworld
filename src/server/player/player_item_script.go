package player

import (
	"server/game/config"
)

type ItemScript = func(plr *Player, ip *config.ItemProto, cnt uint32)

var _item_scripts map[uint32]ItemScript

func init() {
	_item_scripts = make(map[uint32]ItemScript, 0x40)

	// 脚本1: 为玩家增加道具,使用param1,param2参数
	_item_scripts[1] = func(plr *Player, ip *config.ItemProto, cnt uint32) {
		p := NewItemProxy(1)
		if ip.Param1 != 0 && ip.Param2 != 0 {
			p.Add(uint32(ip.Param1), uint64(uint32(ip.Param2)*cnt))
		}
		p.Apply(plr)
	}

	// 脚本2: Wait for ...
}
