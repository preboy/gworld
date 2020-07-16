package task

import (
	"core/event"
	"game/config"
	"game/constant"
)

// ============================================================================
// 模块内常量

const (
	V = 1
)

// ============================================================================
// event

func init() {

	// 杀怪
	event.On(constant.Evt_Plr_KillMonster, func(evt uint32, args ...interface{}) {
		plr := args[0].(iPlayer)
		mid := args[1].(int32)

		for _, t := range plr.GetTask().Tasks {
			if t.Finish {
				continue
			}

			// 计数
			conf := config.TaskConf.Query(t.Id)
			if conf == nil {
				continue
			}

			if conf.Type == constant.TaskType_Kill {
				continue
			}

			// todo record
			mid = mid
		}
	})

	// other event
}
