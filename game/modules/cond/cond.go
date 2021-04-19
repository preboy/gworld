package cond

import (
	"gworld/game/app"
	"gworld/game/config"
)

const (
	CondType_Plr_Lv        uint32 = iota + 1 // 玩家等级
	CondType_Svr_Open_Days                   // 开服天数
	CondType_Duration                        // 时间段
)

// ============================================================================

var checkers = map[uint32]func(app.IPlayer, []int32) bool{

	// 玩家等级
	CondType_Plr_Lv: func(plr app.IPlayer, params []int32) bool {
		return true
	},

	// 开服天数
	CondType_Svr_Open_Days: func(plr app.IPlayer, params []int32) bool {
		return true
	},
}

// ============================================================================

func Check(plr app.IPlayer, id uint32) bool {
	conf := config.CondConf.Query(id)
	if conf == nil {
		return false
	}

	if fn, ok := checkers[conf.Type]; ok {
		if fn(plr, conf.Params) {
			return true
		}
	}

	return false
}
