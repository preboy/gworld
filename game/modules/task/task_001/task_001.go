package task_001

import (
	"gworld/core/event"
	"gworld/game/constant"
	"gworld/game/modules/task"

	"gworld/game/player"
)

// ============================================================================
// Task demo

const (
	// NOTE: 每一个活动ID不可相同
	TaskID = constant.TaskId_Demo
)

var (
	_this_task = &TaskImpl{}
)

type task_data_t struct {
	Cnt int32
	Str string
}

// ============================================================================
// TaskImpl

type TaskImpl struct{}

func (self *TaskImpl) NewTaskData() interface{} {
	return &task_data_t{}
}

func (self *TaskImpl) GetTaskData(plr task.IPlayer) *task_data_t {
	data := plr.GetTask().GetData(TaskID)
	if data != nil {
		return data.(*task_data_t)
	}

	return nil
}

// ============================================================================
// init

func init() {
	task.RegTask(TaskID, _this_task)

	// 注册事件
	event.On(constant.Evt_Plr_Login, func(evt uint32, args ...interface{}) {
		pid := args[0].(string)
		plr := player.GetPlayer(pid)
		data := _this_task.GetTaskData(plr)
		if data == nil {
			return
		}

		// todo
	})
}
