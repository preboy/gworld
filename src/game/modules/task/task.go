package task

import (
	"core/event"
	"game/app"
	"game/config"
	"game/constant"

	"time"
)

// ============================================================================
// event

func init() {
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

// ============================================================================

// 个人活动
type task_item_t struct {
	Id      uint32 // 原型ID
	StartTs uint32 // 开始时间
	Finish  bool
}

type iPlayer interface {
	app.IPlayer

	GetTask() *Task
}

type Task struct {
	plr iPlayer

	Tasks map[uint32]*task_item_t
}

// ============================================================================

func NewTask() *Task {
	return &Task{}
}

func (self *Task) Init(plr iPlayer) {
	self.plr = plr

	self.Tasks = make(map[uint32]*task_item_t)
}

func (self *Task) Add(id uint32) bool {

	conf := config.TaskConf.Query(id)
	if conf == nil {
		return false
	}

	if self.Tasks[id] != nil {
		return false
	}

	self.Tasks[id] = &task_item_t{
		Id:      id,
		StartTs: uint32(time.Now().Unix()),
		Finish:  false,
	}

	return true
}

func (self *Task) Del(id uint32) {
	delete(self.Tasks, id)
}

func (self *Task) Commit(id uint32) {
	//	todo 发奖励

	conf := config.TaskConf.Query(id)
	if conf == nil {
		return
	}

	if self.Tasks[id] != nil {
		return
	}

	// add .

	self.Del(id)
}

// ============================================================================
// Event

func (self *Task) OnKill(mid int32) {

}
