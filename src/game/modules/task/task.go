package task

import (
	"core/event"
	"game/app"
	"game/constant"
)

// ============================================================================
// event

func init() {
	event.On(constant.Evt_Plr_KillMonster, func(evt uint32, args ...interface{}) {
		plr := args[0].(iPlayer)
		mid := args[1].(int32)

		plr.GetTask().OnKill(mid)
	})

	// do
}

// ============================================================================

// 个人活动
type task_item_t struct {
	Id      int32  // 原型ID
	StartTs uint32 // 开始时间
	LastSec uint32 // 持续时间

	Finish bool
}

type iPlayer interface {
	app.IPlayer

	GetTask() *Task
}

type Task struct {
	plr iPlayer

	// s
}

// ============================================================================

func NewTask() *Task {
	return &Task{}
}

func (self *Task) Init(plr iPlayer) {
	self.plr = plr
}

func (self *Task) Add(id uint32) {
}

// ============================================================================
// Event

func (self *Task) OnKill(mid int32) {
}
