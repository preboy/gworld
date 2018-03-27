package player

type ItemScript = func(plr *Player, p1 int32, p2 int32, p3 []int32, p4 string)

var _item_scripts map[uint32]ItemScript

func init() {
	_item_scripts = make(map[uint32]ItemScript, 0x40)

	// 脚本1: 为玩家增加道具
	_item_scripts[1] = func(plr *Player, p1 int32, p2 int32, p3 []int32, p4 string) {
		p := NewItemProxy()
		if p1 != 0 && p2 != 0 {
			p.Add(uint32(p1), uint64(p2))
		}
		p.Apply(plr)
	}

	// Wait for ...
}
