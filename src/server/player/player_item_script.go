package player

type ItemScript = func(plr *Player, p1, p2, p3, p4 uint32)

var _item_scripts map[uint32]ItemScript

func init() {
	_item_scripts = make(map[uint32]ItemScript, 0x100)

	// 脚本1: 为玩家增加道具
	_item_scripts[1] = func(plr *Player, p1, p2, p3, p4 uint32) {
		p := NewItemProxy()
		if p1 != 0 && p2 != 0 {
			p.Add(p1, uint64(p2))
		}
		if p3 != 0 && p4 != 0 {
			p.Add(p3, uint64(p4))
		}
		p.Apply(plr)
	}

	// Wait for ...
}
